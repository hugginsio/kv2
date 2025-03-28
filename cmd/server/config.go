// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"os"
	"path"

	"git.huggins.io/kv2/internal/kms"
	"github.com/rs/zerolog/log"
)

type Configuration struct {
	CloudStorage     string
	ConfigurationDir string
	DevMode          bool
	PrivateKey       string
	PublicKey        string
	TsAuthKey        string
}

// Retrieves the configuration from the environment.
func RetrieveConfiguration() Configuration {
	configuration := Configuration{
		CloudStorage: os.Getenv("KV2_CLOUD_STORAGE"),
		DevMode:      os.Getenv("KV2_DEV_MODE") == "true",
		PrivateKey:   os.Getenv("KV2_PRIVATE_KEY"),
		PublicKey:    os.Getenv("KV2_PUBLIC_KEY"),
		TsAuthKey:    os.Getenv("KV2_TS_AUTHKEY"),
	}

	// Go ahead and run this since we don't fetch the config anywhere else.
	return preflight(configuration)
}

// Check for misconfigurations, print warnings, etc.
func preflight(configuration Configuration) Configuration {
	if configuration.DevMode {
		log.Warn().Msg("RUNNING IN DEVELOPMENT MODE")
	}

	if !configuration.DevMode && configuration.TsAuthKey == "" {
		log.Fatal().Msg("KV2_TS_AUTHKEY is required outside of development mode")
	} else {
		configuration.TsAuthKey = kms.KmsMiddleware(configuration.TsAuthKey)
	}

	if !configuration.DevMode && configuration.PrivateKey == "" {
		log.Fatal().Msg("KV2_PRIVATE_KEY is required")
	} else {
		configuration.PrivateKey = kms.KmsMiddleware(configuration.PrivateKey)
	}

	if !configuration.DevMode && configuration.PublicKey == "" {
		log.Fatal().Msg("KV2_PUBLIC_KEY is required")
	} else {
		configuration.PublicKey = kms.KmsMiddleware(configuration.PublicKey)
	}

	if value, err := os.UserConfigDir(); err != nil {
		log.Fatal().Err(err).Msg("failed to get user config dir")
	} else {
		configuration.ConfigurationDir = path.Join(value, "kv2")
	}

	return configuration
}
