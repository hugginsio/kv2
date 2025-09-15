// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"context"
	"fmt"
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

func NewConnectHandler(backend data.Backend, mux *http.ServeMux) *ServerHandler {
	server := &ServerHandler{
		backend: backend,
	}

	path, handler := nexusv2connect.NewKv2NexusServiceHandler(server, connect.WithCompressMinBytes(860))
	mux.Handle(path, handler)

	return server
}

func (m *ServerHandler) ListSecret(ctx context.Context, req *connect.Request[nexusv2.ListSecretRequest]) (*connect.Response[nexusv2.ListSecretResponse], error) {
	return connect.NewResponse(&nexusv2.ListSecretResponse{}), nil
}

func (m *ServerHandler) CreateSecret(ctx context.Context, req *connect.Request[nexusv2.CreateSecretRequest]) (*connect.Response[nexusv2.CreateSecretResponse], error) {
	// TODO: validate request

	fmt.Println(req.Msg.PublicKey)
	encryptionKey, err := m.backend.GetOrCreateEncryptionKey(req.Msg.PublicKey)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	signature := data.CreatedSignature{
		CreatedBy: "unknown",
	}

	secret := &data.Secret{
		Title: req.Msg.Title,
		SecretVersion: []data.SecretVersion{
			{
				Version:          1,
				Content:          req.Msg.Content,
				CreatedSignature: signature,
				Key:              *encryptionKey,
			},
		},
		CreatedSignature: signature,
	}

	_, err = m.backend.CreateSecret(secret)

	// TODO: redo error handling here
	return connect.NewResponse(&nexusv2.CreateSecretResponse{}), nil
}

func (m *ServerHandler) UpdateSecret(context.Context, *connect.Request[nexusv2.UpdateSecretRequest]) (*connect.Response[nexusv2.UpdateSecretResponse], error) {
	return connect.NewResponse(&nexusv2.UpdateSecretResponse{}), nil
}
