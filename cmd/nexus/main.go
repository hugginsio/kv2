// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"path/filepath"

	goversion "github.com/caarlos0/go-version"
	"github.com/hugginsio/kv2/internal/data"
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
	backend, err := data.NewSqliteBackend(&data.SqliteConfiguration{Path: configuration.Persistence.Endpoint})
	if err != nil {
		// TODO: better error handling
		fmt.Println(err)
		panic(err)
	}

	// TODO: configure backup cronjob if enabled

	port := "8080"

	// TODO: setup tsnet connection if enabled
	if configuration.Tailnet.Enabled && configuration.Tailnet.TLS {
		port = "443"
	}

	slog.Info("API listening on", "port", port)
	mux := http.NewServeMux()
	handler := NewConnectHandler(backend, mux)
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		// TODO: better error handling
		panic(err)
	}

	ServeHealthEndpoint()

	if err := http.Serve(ln, handler); err != nil {
		// TODO: better error handling
		panic(err)
	}
}
