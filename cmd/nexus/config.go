// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	_ "embed"
	"errors"
)

//go:embed default.yaml
var DefaultConfigYAML []byte

// Nexus configuration. For defaults, see `default.yaml`.
type Config struct {
	Persistence PersistenceConfig `yaml:"persistence"`
	Tailnet     TailnetConfig     `yaml:"tailnet"`
	Kms         KmsConfig         `yaml:"kms"`
}

type PersistenceConfig struct {
	Enabled     bool          `yaml:"enabled"`      // Toggle data persistence. If false (for development purposes), uses in-memory SQLite.
	Type        string        `yaml:"type"`         // Specify backend type, e.g. `sqlite`.
	Endpoint    string        `yaml:"endpoint"`     // Specify data endpoint, e.g. `/data/kv2.db`.
	Backup      BackupConfig  `yaml:"backup"`       // Configure automatic backups.
	Restore     RestoreConfig `yaml:"restore"`      // Configure automatic restore.
	MaxVersions int           `yaml:"max-versions"` // Maximum number of versions to keep for each secret.
}

type TailnetConfig struct {
	Enabled  bool   `yaml:"enabled"`  // Toggle tailnet connectivity.
	KeyVar   string `yaml:"key-var"`  // Specify the environment variable name holding the authorization key.
	KeyKms   string `yaml:"key-kms"`  // Specify the KMS identifier for the authorization key, e.g. `gsm://projects/1234/secrets/ts-authkey`.
	Hostname string `yaml:"hostname"` // Specify the hostname for tailnet connectivity.
	TLS      bool   `yaml:"tls"`      // Toggle TLS requirement for server.
}

type BackupConfig struct {
	Enabled  bool   `yaml:"enabled"`  // Toggle automatic backups.
	Schedule string `yaml:"schedule"` // Specify backup schedule via a Cron expression.
	Endpoint string `yaml:"endpoint"` // Specify backup endpoint, e.g. `gcs://kv2-backup`.
}

type RestoreConfig struct {
	Enabled bool `yaml:"enabled"` // Toggle automatic restore.
}

type KmsConfig struct {
	Enabled bool `yaml:"enabled"` // Toggle KMS integration.
}

func Validate(config *Config) []error {
	var results []error
	if config == nil {
		results = append(results, errors.New("config cannot be nil"))
		return results
	}

	if config.Persistence.Enabled {
		if config.Persistence.Type != "sqlite" {
			results = append(results, errors.New("persistence.type must be sqlite"))
		}

		if config.Persistence.Endpoint == "" {
			results = append(results, errors.New("persistence.endpoint cannot be empty"))
		}

		if config.Persistence.Backup.Enabled {
			if config.Persistence.Backup.Schedule == "" {
				results = append(results, errors.New("persistence.backup.schedule cannot be empty while backup is enabled"))
			}
			if config.Persistence.Backup.Endpoint == "" {
				results = append(results, errors.New("persistence.backup.endpoint cannot be empty while backup is enabled"))
			}
		}
	}

	if config.Tailnet.Enabled {
		if config.Tailnet.KeyVar == "" && config.Tailnet.KeyKms == "" {
			results = append(results, errors.New("tailnet.authkey-var or tailnet.authkey-kms must be specified"))
		}

		if config.Tailnet.Hostname == "" {
			results = append(results, errors.New("tailnet.hostname cannot be empty"))
		}
	}

	// TODO: kms block

	return results
}
