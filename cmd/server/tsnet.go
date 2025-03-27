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

	ln, err := server.Listen("tcp", ":80")
	if err != nil {
		log.Fatal().Err(err).Str("addr", ln.Addr().String()).Msg("failed to listen")
	}

	if _, err := server.Up(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("failed to start tsnet")
	}

	return ln
}
