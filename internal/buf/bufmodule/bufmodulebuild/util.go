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

package bufmodulebuild

import (
	"errors"
	"fmt"
	"sort"

	"github.com/bufbuild/buf/internal/buf/bufmodule"
	"github.com/bufbuild/buf/internal/pkg/normalpath"
	"github.com/bufbuild/buf/internal/pkg/stringutil"
)

func applyModulePaths(
	module bufmodule.Module,
	roots []string,
	fileOrDirPaths *[]string,
	fileOrDirPathsAllowNotExist bool,
	pathType normalpath.PathType,
) (bufmodule.Module, error) {
	if fileOrDirPaths == nil {
		return module, nil
	}
	targetPaths, err := pathsToTargetPaths(roots, *fileOrDirPaths, pathType)
	if err != nil {
		return nil, err
	}
	if fileOrDirPathsAllowNotExist {
		return bufmodule.ModuleWithTargetPathsAllowNotExist(module, targetPaths)
	}
	return bufmodule.ModuleWithTargetPaths(module, targetPaths)
}

func pathsToTargetPaths(roots []string, paths []string, pathType normalpath.PathType) ([]string, error) {
	if len(roots) == 0 {
		// this should never happen
		return nil, errors.New("no roots on config")
	}

	targetPaths := make([]string, len(paths))
	for i, path := range paths {
		targetPath, err := pathToTargetPath(roots, path, pathType)
		if err != nil {
			return nil, err
		}
		targetPaths[i] = targetPath
	}
	return targetPaths, nil
}

func pathToTargetPath(roots []string, path string, pathType normalpath.PathType) (string, error) {
	var matchingRoots []string
	for _, root := range roots {
		if normalpath.ContainsPath(root, path, pathType) {
			matchingRoots = append(matchingRoots, root)
		}
	}
	switch len(matchingRoots) {
	case 0:
		// this is a user error and will likely happen often
		return "", fmt.Errorf(
			"path %q is not contained within any of roots %s - note that specified paths "+
				"cannot be roots, but must be contained within roots",
			path,
			stringutil.SliceToHumanStringQuoted(roots),
		)
	case 1:
		targetPath, err := normalpath.Rel(matchingRoots[0], path)
		if err != nil {
			return "", err
		}
		// just in case
		return normalpath.NormalizeAndValidate(targetPath)
	default:
		// this should never happen
		return "", fmt.Errorf("%q is contained in multiple roots %s", path, stringutil.SliceToHumanStringQuoted(roots))
	}
}

// normalizeAndCheckPaths verifies that:
//
//   - No paths are empty.
//   - All paths are normalized and validated if pathType is Relative.
//   - All paths are normalized if pathType is Absolute.
//
// If sortAndCheckDuplicates is true:

//   - All paths are unique.
//   - No path contains another path.
//
// Normalizes the paths.
// Sorts the paths if sortAndCheckDuplicates is true.
// Makes the paths absolute if pathType is Absolute.
func normalizeAndCheckPaths(
	paths []string,
	name string,
	pathType normalpath.PathType,
	sortAndCheckDuplicates bool,
) ([]string, error) {
	if len(paths) == 0 {
		return paths, nil
	}
	outputs := make([]string, len(paths))
	for i, path := range paths {
		if path == "" {
			return nil, fmt.Errorf("%s contained an empty path", name)
		}
		output, err := normalpath.NormalizeAndTransformForPathType(path, pathType)
		if err != nil {
			// user error
			return nil, err
		}
		outputs[i] = output
	}
	if sortAndCheckDuplicates {
		return sortAndCheckDuplicatePaths(outputs, name, pathType)
	}
	return outputs, nil
}

// TODO: refactor this
func sortAndCheckDuplicatePaths(outputs []string, name string, pathType normalpath.PathType) ([]string, error) {
	sort.Strings(outputs)
	for i := 0; i < len(outputs); i++ {
		for j := i + 1; j < len(outputs); j++ {
			output1 := outputs[i]
			output2 := outputs[j]

			if output1 == output2 {
				return nil, fmt.Errorf("duplicate %s %q", name, output1)
			}
			if normalpath.EqualsOrContainsPath(output2, output1, pathType) {
				return nil, fmt.Errorf("%s %q is within %s %q which is not allowed", name, output1, name, output2)
			}
			if normalpath.EqualsOrContainsPath(output1, output2, pathType) {
				return nil, fmt.Errorf("%s %q is within %s %q which is not allowed", name, output2, name, output1)
			}
		}
	}

	return outputs, nil
}

type buildOptions struct {
	moduleIdentity bufmodule.ModuleIdentity
	// If nil, all files are considered targets.
	// If empty (but non-nil), the module will have no target paths.
	paths              *[]string
	pathsAllowNotExist bool
}

type buildModuleFileSetOptions struct {
	workspace bufmodule.Workspace
}
