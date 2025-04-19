// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"git.huggins.io/kv2/internal/backup"
	"git.huggins.io/kv2/internal/backup/nobackup"
	"git.huggins.io/kv2/internal/crypto"
	"git.huggins.io/kv2/internal/crypto/age"
	"git.huggins.io/kv2/internal/crypto/nocrypto"
	"git.huggins.io/kv2/internal/database"
	"git.huggins.io/kv2/internal/database/sqlite"
	"git.huggins.io/kv2/internal/server"
	"git.huggins.io/kv2/internal/version"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	log.Info().Str("version", version.VersionInfo().GitVersion).Str("platform", version.VersionInfo().Platform).Msg("kv2 is starting")

	appConfig := RetrieveConfiguration()

	defaultDatabasePath := fmt.Sprintf("%s/kv2.db", appConfig.ConfigurationDir)
	databaseConfiguration := sqlite.Configuration{
		Dsn: fmt.Sprintf("file:%s", defaultDatabasePath),
	}

	if appConfig.DevMode {
		databaseConfiguration.Dsn = ":memory:"
	}

	var cloudStorage backup.CloudBackup = nobackup.Initialize()
	if !appConfig.DevMode && appConfig.CloudStorage != "" {
		if provider, err := backup.DetermineStorageProvider(appConfig.CloudStorage, defaultDatabasePath); err != nil {
			log.Fatal().Err(err).Msg("failed to configure cloud storage provider")
		} else {
			cloudStorage = *provider
		}

		if _, err := os.Stat(defaultDatabasePath); !os.IsNotExist(err) {
			log.Info().Msg("database found, skipping restore")
		} else {
			if err := cloudStorage.Restore(); err != nil {
				log.Fatal().Err(err).Msg("database restore failed")
			}
		}
	}

	var database database.Database
	database, err := sqlite.Initialize(databaseConfiguration)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load database")
	}

	var crypto crypto.Crypto
	if appConfig.DevMode {
		crypto = nocrypto.Initialize()
	} else {
		crypto, err = age.Initialize(age.Configuration{
			PrivateKey: appConfig.PrivateKey,
			PublicKey:  appConfig.PublicKey,
		})

		if err != nil {
			log.Fatal().Err(err).Msg("failed to initialize crypto")
		}
	}

	mux := http.NewServeMux()
	serverConfig := server.Configuration{
		CloudBackup: &cloudStorage,
		Crypto:      &crypto,
		Database:    &database,
		Mux:         mux,
	}

	_ = server.Initialize(serverConfig)
	var listener net.Listener

	if appConfig.DevMode {
		ln, err := net.Listen("tcp", ":8081")
		if err != nil {
			log.Fatal().Err(err).Str("addr", ln.Addr().String()).Msg("failed to listen")
		}

		listener = ln
	} else {
		listener = Tsnet(appConfig)
	}

	go ServeHealthEndpoint()

	log.Info().Str("addr", listener.Addr().String()).Msg("serving API")
	if err := http.Serve(listener, h2c.NewHandler(mux, &http2.Server{})); err != nil {
		log.Error().Err(err).Msg("failed to start server")
	}
}
