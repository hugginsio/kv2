package main

import (
	"context"
	"net"

	"github.com/rs/zerolog/log"
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

	_, err := server.Up(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("tailscale failed to start")
	}

	ln, err := server.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal().Err(err).Str("addr", ln.Addr().String()).Msg("failed to listen")
	}

	if _, err := server.Up(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("failed to start tsnet")
	}

	return ln
}
