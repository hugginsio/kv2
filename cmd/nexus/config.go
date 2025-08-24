// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

type Config struct {
	Persistence PersistenceConfig `yaml:"persistence"`
	Tailnet     TailnetConfig     `yaml:"tailnet"`
}

type PersistenceConfig struct {
	Enabled  bool   `yaml:"enabled"`  // Toggle data persistence. If false (for development purposes), uses in-memory SQLite.
	Type     string `yaml:"type"`     // Specify backend type (e.g. `sqlite`).
	Endpoint string `yaml:"endpoint"` // Specify data endpoint. For `sqlite`, this is a filepath and defaults to `/data/kv2-nexus.db`.
}

type TailnetConfig struct {
	Enabled    bool   `yaml:"enabled"`     // Toggle tailnet connectivity.
	AuthkeyVar string `yaml:"authkey-var"` // Specify the environment variable name holding the authorization key.
}

// TODO: config validation method
// TODO: if backend is sqlite and endpoint is not absolute path, join with pwd
