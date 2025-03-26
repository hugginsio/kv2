// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"fmt"
	"log"
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
)

func main() {
	log.Println(version.Line())
	log.Println(version.Seymour())

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
			log.Fatal("Failed to configure cloud storage provider: ", err)
		} else {
			cloudStorage = *provider
		}

		if _, err := os.Stat(defaultDatabasePath); !os.IsNotExist(err) {
			log.Println("Existing database found, skipping restore")
		} else {
			if err := cloudStorage.Restore(); err != nil {
				log.Println("Restore failed:", err)
			}
		}
	}

	var database database.Database
	database, err := sqlite.Initialize(databaseConfiguration)
	if err != nil {
		log.Fatalln("failed to load database:", err)
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
			log.Fatalln("failed to initialize crypto:", err)
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
		ln, err := net.Listen("tcp", ":80")
		if err != nil {
			log.Fatalln(err)
		}

		listener = ln
	} else {
		listener = Tsnet(appConfig)
	}

	go ServeHealthEndpoint()

	log.Println("Serving API & healthcheck")
	if err := http.Serve(listener, mux); err != nil {
		log.Fatalln("failed to start server:", err)
	}
}
