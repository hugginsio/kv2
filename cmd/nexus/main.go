// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"log/slog"
	"os"
	"path/filepath"

	goversion "github.com/caarlos0/go-version"
	"github.com/hugginsio/kv2/internal/discovery"
	"github.com/hugginsio/kv2/internal/yaml"
)

func main() {
	if discovery.IsDevel() {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	slog.Info("nexus is starting", "version", goversion.GetVersionInfo().GitVersion)

	configurationParent := "/app/"
	if !discovery.IsContainerized() {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		configurationParent = wd
	}

	var configuration Config
	configurationFile := "config.yaml"
	if discovery.IsDevel() {
		configurationFile = "devel.yaml"
	}

	configurationPath := filepath.Join(configurationParent, configurationFile)

	slog.Debug("loading configuration", "path", configurationPath)

	if err := yaml.LoadYamlBytes(DefaultConfigYAML, &configuration); err != nil {
		// TODO: better error handling
		panic(err)
	}

	if err := yaml.LoadYaml(configurationPath, &configuration); err != nil {
		// TODO: better error handling
		panic(err)
	}

	if errs := Validate(&configuration); len(errs) > 0 {
		for _, err := range errs {
			slog.Error("configuration validation error", "error", err)
		}

		// TODO: better error handling
		panic(errs)
	}

	if !filepath.IsAbs(configuration.Persistence.Endpoint) {
		configuration.Persistence.Endpoint = filepath.Join(configurationParent, configuration.Persistence.Endpoint)
	}

	// TODO: restore from backup if enabled and present

	slog.Debug("loading database", "endpoint", configuration.Persistence.Endpoint)

	// TODO: database init

	// TODO: setup tsnet connection if enabled

	// TODO: configure backup cronjob if enabled

	// TODO: setup web service
}
