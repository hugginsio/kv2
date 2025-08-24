// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"log/slog"
	"os"

	goversion "github.com/caarlos0/go-version"
	"github.com/hugginsio/kv2/internal/discovery"
)

func main() {
	slog.Info("nexus is starting", "version", goversion.GetVersionInfo().GitVersion)

	// TODO: check for KV2_DEVEL environment variable and enable debug logs if so

	configurationParent := "/app/"
	if !discovery.IsContainerized() {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		configurationParent = wd
	}

	slog.Debug("loading configuration", "path", configurationParent)

	// TODO: environment discovery
	// TODO: load config
	// TODO: validate config
	// TODO: database init
	// TODO: setup tsnet connection if needed
	// TODO: setup web service
}
