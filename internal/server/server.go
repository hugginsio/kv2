// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"connectrpc.com/connect"
	secretsv1 "git.huggins.io/kv2/api/secrets/v1"
	"git.huggins.io/kv2/api/secrets/v1/secretsv1connect"
	"git.huggins.io/kv2/internal/backup"
	"git.huggins.io/kv2/internal/crypto"
	"git.huggins.io/kv2/internal/database"
	"git.huggins.io/kv2/internal/version"
	"github.com/rs/zerolog/log"
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

func auditRequest[T any](req *connect.Request[T], message string) {
	peer := strings.Split(req.Peer().Addr, ":")[0]
	log.Debug().Str("peer", peer).Str(USER_AGENT, req.Header().Get(USER_AGENT)).Msg(message)
}

func (h *HttpServer) CreateSecret(ctx context.Context, req *connect.Request[secretsv1.CreateSecretRequest]) (*connect.Response[secretsv1.CreateSecretResponse], error) {
	auditRequest(req, fmt.Sprintf("creating %s", req.Msg.Key))
	if err := h.database.Create(req.Msg); err != nil {
		// TODO: error handling?
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := &secretsv1.CreateSecretResponse{
		Secret: &secretsv1.SecretMetadata{Key: req.Msg.Key, Versions: []uint32{1}},
	}

	return &connect.Response[secretsv1.CreateSecretResponse]{Msg: res}, nil
}

func (h *HttpServer) GetSecret(ctx context.Context, req *connect.Request[secretsv1.GetSecretRequest]) (*connect.Response[secretsv1.GetSecretResponse], error) {
	auditRequest(req, fmt.Sprintf("retrieving %s", req.Msg.Key))
	res, err := h.database.Read(req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return &connect.Response[secretsv1.GetSecretResponse]{Msg: res}, nil
}

func (h *HttpServer) UpdateSecret(ctx context.Context, req *connect.Request[secretsv1.UpdateSecretRequest]) (*connect.Response[secretsv1.UpdateSecretResponse], error) {
	auditRequest(req, fmt.Sprintf("updating %s", req.Msg.Key))
	if res, err := h.database.Update(req.Msg); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	} else {
		return &connect.Response[secretsv1.UpdateSecretResponse]{Msg: res}, nil
	}
}

func (h *HttpServer) DeleteSecret(ctx context.Context, req *connect.Request[secretsv1.DeleteSecretRequest]) (*connect.Response[secretsv1.DeleteSecretResponse], error) {
	auditRequest(req, fmt.Sprintf("deleting %s", req.Msg.Key))
	if err := h.database.Delete(req.Msg); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return &connect.Response[secretsv1.DeleteSecretResponse]{Msg: &secretsv1.DeleteSecretResponse{}}, nil
}

func (h *HttpServer) RevertSecret(ctx context.Context, req *connect.Request[secretsv1.RevertSecretRequest]) (*connect.Response[secretsv1.RevertSecretResponse], error) {
	auditRequest(req, fmt.Sprintf("reverting %s", req.Msg.Key))
	if res, err := h.database.Revert(req.Msg); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	} else {
		return &connect.Response[secretsv1.RevertSecretResponse]{Msg: res}, nil
	}
}

func (h *HttpServer) ListSecrets(ctx context.Context, req *connect.Request[secretsv1.ListSecretsRequest]) (*connect.Response[secretsv1.ListSecretsResponse], error) {
	auditRequest(req, "ListSecrets")
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

	auditRequest(req, fmt.Sprintf("backing up to %s", *req.Msg.Name))
	if err := h.backup.Backup(backupName); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return &connect.Response[secretsv1.BackupResponse]{Msg: &secretsv1.BackupResponse{}}, nil
}

func (h *HttpServer) ApplicationVersionInfo(ctx context.Context, req *connect.Request[secretsv1.ApplicationVersionInfoRequest]) (*connect.Response[secretsv1.ApplicationVersionInfoResponse], error) {
	auditRequest(req, "version info requested")
	return &connect.Response[secretsv1.ApplicationVersionInfoResponse]{Msg: &secretsv1.ApplicationVersionInfoResponse{Info: version.VersionInfo()}}, nil
}
