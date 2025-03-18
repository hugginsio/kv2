// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"log"
	"net"
	"net/http"

	"git.huggins.io/kv2/internal/crypto"
	"git.huggins.io/kv2/internal/crypto/age"
	"git.huggins.io/kv2/internal/database"
	"git.huggins.io/kv2/internal/database/sqlite"
	"git.huggins.io/kv2/internal/server"
	"git.huggins.io/kv2/internal/version"
)

func main() {
	log.Println(version.Line())
	log.Println(version.Seymour())

	appConfig := RetrieveConfiguration()

	databaseConfiguration := sqlite.Configuration{
		Dsn: "kv2.db",
	}

	if appConfig.DevMode {
		databaseConfiguration.Dsn = ":memory:"
	} else {
		databaseConfiguration.Dsn = "kv2.db"
	}

	var database database.Database
	database, err := sqlite.Initialize(databaseConfiguration)
	if err != nil {
		log.Fatalln("failed to load database:", err)
	}

	var crypto crypto.Crypto
	crypto, err = age.Initialize(age.Configuration{
		PrivateKey: appConfig.PrivateKey,
		PublicKey:  appConfig.PublicKey,
	})

	if err != nil {
		log.Fatalln("failed to initialize crypto:", err)
	}

	mux := http.NewServeMux()
	serverConfig := server.Configuration{
		Crypto:   &crypto,
		Database: &database,
		Mux:      mux,
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
