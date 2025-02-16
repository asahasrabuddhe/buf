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

// Package bufbreaking contains the breaking change detection functionality.
//
// The primary entry point to this package is the Handler.
package bufbreaking

import (
	"context"

	"github.com/bufbuild/buf/internal/buf/bufanalysis"
	"github.com/bufbuild/buf/internal/buf/bufcheck"
	"github.com/bufbuild/buf/internal/buf/bufcheck/bufbreaking/internal/bufbreakingv1"
	"github.com/bufbuild/buf/internal/buf/bufcheck/bufbreaking/internal/bufbreakingv1beta1"
	"github.com/bufbuild/buf/internal/buf/bufcheck/internal"
	"github.com/bufbuild/buf/internal/buf/bufimage"
	"go.uber.org/zap"
)

// Handler handles the main breaking functionality.
type Handler interface {
	// Check runs the breaking checks.
	//
	// The image should have source code info for this to work properly. The previousImage
	// does not need to have source code info.
	//
	// Images should be filtered with regards to imports before passing to this function.
	Check(
		ctx context.Context,
		config *Config,
		previousImage bufimage.Image,
		image bufimage.Image,
	) ([]bufanalysis.FileAnnotation, error)
}

// NewHandler returns a new Handler.
func NewHandler(logger *zap.Logger) Handler {
	return newHandler(logger)
}

// Rule is a rule.
type Rule interface {
	bufcheck.Rule

	internalBreaking() *internal.Rule
}

// Config is the check config.
type Config struct {
	// Rules are the rules to run.
	//
	// Rules will be sorted by first categories, then id when Configs are
	// created from this package, i.e. created wth ConfigBuilder.NewConfig.
	Rules                  []Rule
	IgnoreIDToRootPaths    map[string]map[string]struct{}
	IgnoreRootPaths        map[string]struct{}
	IgnoreUnstablePackages bool
}

// GetRules returns the rules.
//
// Should only be used for printing.
func (c *Config) GetRules() []bufcheck.Rule {
	return rulesToBufcheckRules(c.Rules)
}

// NewConfigV1Beta1 returns a new Config.
func NewConfigV1Beta1(externalConfig ExternalConfigV1Beta1) (*Config, error) {
	internalConfig, err := internal.ConfigBuilder{
		Use:                           externalConfig.Use,
		Except:                        externalConfig.Except,
		IgnoreRootPaths:               externalConfig.Ignore,
		IgnoreIDOrCategoryToRootPaths: externalConfig.IgnoreOnly,
		IgnoreUnstablePackages:        externalConfig.IgnoreUnstablePackages,
	}.NewConfig(
		bufbreakingv1beta1.VersionSpec,
	)
	if err != nil {
		return nil, err
	}
	return internalConfigToConfig(internalConfig), nil
}

// NewConfigV1 returns a new Config.
func NewConfigV1(externalConfig ExternalConfigV1) (*Config, error) {
	internalConfig, err := internal.ConfigBuilder{
		Use:                           externalConfig.Use,
		Except:                        externalConfig.Except,
		IgnoreRootPaths:               externalConfig.Ignore,
		IgnoreIDOrCategoryToRootPaths: externalConfig.IgnoreOnly,
		IgnoreUnstablePackages:        externalConfig.IgnoreUnstablePackages,
	}.NewConfig(
		bufbreakingv1.VersionSpec,
	)
	if err != nil {
		return nil, err
	}
	return internalConfigToConfig(internalConfig), nil
}

// GetAllRulesV1Beta1 gets all known rules.
//
// Should only be used for printing.
func GetAllRulesV1Beta1() ([]bufcheck.Rule, error) {
	config, err := NewConfigV1Beta1(
		ExternalConfigV1Beta1{
			Use: bufbreakingv1beta1.VersionSpec.AllCategories,
		},
	)
	if err != nil {
		return nil, err
	}
	return rulesToBufcheckRules(config.Rules), nil
}

// GetAllRulesV1 gets all known rules.
//
// Should only be used for printing.
func GetAllRulesV1() ([]bufcheck.Rule, error) {
	config, err := NewConfigV1(
		ExternalConfigV1{
			Use: bufbreakingv1.VersionSpec.AllCategories,
		},
	)
	if err != nil {
		return nil, err
	}
	return rulesToBufcheckRules(config.Rules), nil
}

// ExternalConfigV1Beta1 is an external config.
type ExternalConfigV1Beta1 struct {
	Use    []string `json:"use,omitempty" yaml:"use,omitempty"`
	Except []string `json:"except,omitempty" yaml:"except,omitempty"`
	// IgnoreRootPaths
	Ignore []string `json:"ignore,omitempty" yaml:"ignore,omitempty"`
	// IgnoreIDOrCategoryToRootPaths
	IgnoreOnly             map[string][]string `json:"ignore_only,omitempty" yaml:"ignore_only,omitempty"`
	IgnoreUnstablePackages bool                `json:"ignore_unstable_packages,omitempty" yaml:"ignore_unstable_packages,omitempty"`
}

// ExternalConfigV1 is an external config.
type ExternalConfigV1 struct {
	Use    []string `json:"use,omitempty" yaml:"use,omitempty"`
	Except []string `json:"except,omitempty" yaml:"except,omitempty"`
	// IgnoreRootPaths
	Ignore []string `json:"ignore,omitempty" yaml:"ignore,omitempty"`
	// IgnoreIDOrCategoryToRootPaths
	IgnoreOnly             map[string][]string `json:"ignore_only,omitempty" yaml:"ignore_only,omitempty"`
	IgnoreUnstablePackages bool                `json:"ignore_unstable_packages,omitempty" yaml:"ignore_unstable_packages,omitempty"`
}

func internalConfigToConfig(internalConfig *internal.Config) *Config {
	return &Config{
		Rules:                  internalRulesToRules(internalConfig.Rules),
		IgnoreIDToRootPaths:    internalConfig.IgnoreIDToRootPaths,
		IgnoreRootPaths:        internalConfig.IgnoreRootPaths,
		IgnoreUnstablePackages: internalConfig.IgnoreUnstablePackages,
	}
}

func configToInternalConfig(config *Config) *internal.Config {
	return &internal.Config{
		Rules:                  rulesToInternalRules(config.Rules),
		IgnoreIDToRootPaths:    config.IgnoreIDToRootPaths,
		IgnoreRootPaths:        config.IgnoreRootPaths,
		IgnoreUnstablePackages: config.IgnoreUnstablePackages,
	}
}

func rulesToBufcheckRules(rules []Rule) []bufcheck.Rule {
	if rules == nil {
		return nil
	}
	s := make([]bufcheck.Rule, len(rules))
	for i, e := range rules {
		s[i] = e
	}
	return s
}
