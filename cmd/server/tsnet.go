package main

import (
	"context"
	"log"
	"net"

	"tailscale.com/tsnet"
)

type TsnetServer struct{}

func Tsnet(config Configuration) net.Listener {
	server := &tsnet.Server{
		AuthKey:   config.TsAuthKey,
		Ephemeral: true,
		Hostname:  "kv2",
		UserLogf:  log.Printf,
	}

	if config.DevMode {
		server.Logf = log.Printf
	}

	ln, err := server.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := server.Up(context.Background()); err != nil {
		log.Fatalln(err)
	}

	return ln
}
