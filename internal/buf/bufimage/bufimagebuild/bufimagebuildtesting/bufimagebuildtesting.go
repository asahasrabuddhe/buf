// Copyright 2020-2021 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bufimagebuildtesting

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/bufbuild/buf/internal/buf/bufimage/bufimagebuild"
	"github.com/bufbuild/buf/internal/buf/bufmodule"
	"github.com/bufbuild/buf/internal/buf/bufmodule/bufmodulebuild"
	"github.com/bufbuild/buf/internal/pkg/storage/storagemem"
	"go.uber.org/zap"
)

const fileSplitter = "----------------\n"

// Fuzz is the entrypoint for the fuzzer.
// We use https://github.com/dvyukov/go-fuzz for fuzzing.
// Please follow the instructions
// in their README for help with running the fuzz targets.
func Fuzz(data []byte) int {
	i, _ := fuzz(data)
	return i
}

func fuzz(data []byte) (int, error) {
	files := bytes.Split(data, []byte(fileSplitter))
	pathToData := make(map[string][]byte)
	for _, file := range files {
		buf := bytes.NewBuffer(file)
		header, err := buf.ReadBytes('\n')
		if err != nil {
			return 0, err
		}
		content, err := io.ReadAll(buf)
		if err != nil {
			return 0, err
		}
		// Trim comment prefix and newline suffix
		fileName := strings.TrimSpace(strings.TrimPrefix(string(header), "// "))
		pathToData[fileName] = content
	}
	bucket, err := storagemem.NewReadBucket(pathToData)
	if err != nil {
		return 0, err
	}
	config, err := bufmodulebuild.NewConfigV1(bufmodulebuild.ExternalConfigV1{})
	if err != nil {
		return 0, err
	}
	module, err := bufmodulebuild.NewModuleBucketBuilder(zap.NewNop()).BuildForBucket(
		context.Background(),
		bucket,
		config,
	)
	if err != nil {
		return 0, err
	}
	moduleFileSet, err := bufmodulebuild.NewModuleFileSetBuilder(
		zap.NewNop(),
		bufmodule.NewNopModuleReader(),
	).Build(
		context.Background(),
		module,
	)
	if err != nil {
		return 0, err
	}
	builder := bufimagebuild.NewBuilder(zap.NewNop())
	_, fileAnnotations, err := builder.Build(context.Background(), moduleFileSet)
	if err != nil {
		// Panic when we captured a panic in the builder
		if strings.Contains(err.Error(), "panic: ") {
			panic(err)
		}
		return 0, err
	}
	if len(fileAnnotations) > 0 {
		return 0, fmt.Errorf("build failed: %v", fileAnnotations)
	}
	return 1, nil
}
