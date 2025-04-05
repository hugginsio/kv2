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
		log.Fatal().Err(err).Msg("Tailscale failed to start")
	}

	if len(server.CertDomains()) == 0 {
		log.Fatal().Msg("no TLS domains found in Tailscale")
	}

	ln, err := server.ListenTLS("tcp", ":443")
	if err != nil {
		log.Fatal().Err(err).Str("addr", ln.Addr().String()).Msg("failed to listen")
	}

	return ln
}
