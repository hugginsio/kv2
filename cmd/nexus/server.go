// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	nexusv2 "github.com/hugginsio/kv2/api/nexus/v2"
	"github.com/hugginsio/kv2/api/nexus/v2/nexusv2connect"
	"github.com/hugginsio/kv2/internal/data"
)

type ServerHandler struct {
	backend data.Backend
	mux     *http.ServeMux
}

// ServeHTTP implements http.Handler.
func (c *ServerHandler) ServeHTTP(http.ResponseWriter, *http.Request) {
	panic("unimplemented")
}

func NewConnectHandler(backend data.Backend, mux *http.ServeMux) *ServerHandler {
	server := &ServerHandler{
		backend: backend,
	}

	path, handler := nexusv2connect.NewKv2NexusServiceHandler(server, connect.WithCompressMinBytes(860))
	mux.Handle(path, handler)

	return server
}

func (c *ServerHandler) CreateSecret(context.Context, *connect.Request[nexusv2.CreateSecretRequest]) (*connect.Response[nexusv2.CreateSecretResponse], error) {
	panic("unimplemented")
}

func (c *ServerHandler) UpdateSecret(context.Context, *connect.Request[nexusv2.UpdateSecretRequest]) (*connect.Response[nexusv2.UpdateSecretResponse], error) {
	panic("unimplemented")
}
