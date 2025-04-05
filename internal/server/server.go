// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package server

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	secretsv1 "git.huggins.io/kv2/api/secrets/v1"
	"git.huggins.io/kv2/api/secrets/v1/secretsv1connect"
	"git.huggins.io/kv2/internal/backup"
	"git.huggins.io/kv2/internal/crypto"
	"git.huggins.io/kv2/internal/database"
)

type Configuration struct {
	CloudBackup *backup.CloudBackup
	Crypto      *crypto.Crypto
	Database    *database.Database
	Mux         *http.ServeMux
}

type HttpServer struct {
	backup   backup.CloudBackup
	crypto   crypto.Crypto
	database database.Database
}

// Initialize an HTTP server.
func Initialize(config Configuration) *HttpServer {
	server := &HttpServer{
		backup:   *config.CloudBackup,
		crypto:   *config.Crypto,
		database: *config.Database,
	}

	path, handler := secretsv1connect.NewKv2ServiceHandler(server, connect.WithCompressMinBytes(860))
	config.Mux.Handle(path, handler)

	return server
}

func (h *HttpServer) CreateSecret(ctx context.Context, req *connect.Request[secretsv1.CreateSecretRequest]) (*connect.Response[secretsv1.CreateSecretResponse], error) {
	if err := h.database.Create(req.Msg); err != nil {
		// TODO: error handling?
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := &secretsv1.CreateSecretResponse{
		Secret: &secretsv1.SecretMetadata{Key: req.Msg.Key, Version: []uint32{1}},
	}

	return &connect.Response[secretsv1.CreateSecretResponse]{Msg: res}, nil
}

func (h *HttpServer) GetSecret(ctx context.Context, req *connect.Request[secretsv1.GetSecretRequest]) (*connect.Response[secretsv1.GetSecretResponse], error) {
	res, err := h.database.Read(req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return &connect.Response[secretsv1.GetSecretResponse]{Msg: res}, nil
}

func (h *HttpServer) UpdateSecret(ctx context.Context, req *connect.Request[secretsv1.UpdateSecretRequest]) (*connect.Response[secretsv1.UpdateSecretResponse], error) {
	if res, err := h.database.Update(req.Msg); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	} else {
		return &connect.Response[secretsv1.UpdateSecretResponse]{Msg: res}, nil
	}
}

func (h *HttpServer) DeleteSecret(ctx context.Context, req *connect.Request[secretsv1.DeleteSecretRequest]) (*connect.Response[secretsv1.DeleteSecretResponse], error) {
	if err := h.database.Delete(req.Msg); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return &connect.Response[secretsv1.DeleteSecretResponse]{Msg: &secretsv1.DeleteSecretResponse{}}, nil
}

func (h *HttpServer) RevertSecret(ctx context.Context, req *connect.Request[secretsv1.RevertSecretRequest]) (*connect.Response[secretsv1.RevertSecretResponse], error) {
	if res, err := h.database.Revert(req.Msg); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	} else {
		return &connect.Response[secretsv1.RevertSecretResponse]{Msg: res}, nil
	}
}

func (h *HttpServer) ListSecrets(ctx context.Context, req *connect.Request[secretsv1.ListSecretsRequest]) (*connect.Response[secretsv1.ListSecretsResponse], error) {
	res, err := h.database.List()
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return &connect.Response[secretsv1.ListSecretsResponse]{Msg: res}, nil
}

func (h *HttpServer) Backup(ctx context.Context, req *connect.Request[secretsv1.BackupRequest]) (*connect.Response[secretsv1.BackupResponse], error) {
	backupName := req.Msg.GetName()
	if backupName == "" {
		backupName = "kv2.db"
	}

	if err := h.backup.Backup(backupName); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return &connect.Response[secretsv1.BackupResponse]{Msg: &secretsv1.BackupResponse{}}, nil
}
