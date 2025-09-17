// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"context"
	"errors"
	"testing"

	"connectrpc.com/connect"
	nexusv2 "github.com/hugginsio/kv2/api/nexus/v2"
	"github.com/hugginsio/kv2/internal/data"
)

// mockBackend implements the data.Backend interface for testing
type mockBackend struct {
	getOrCreateEncryptionKeyFunc func(string) (*data.EncryptionKey, error)
	createSecretFunc             func(*data.Secret) (*data.Secret, error)
	listSecretsFunc              func() ([]*data.Secret, error)
	getSecretVersionsFunc        func(string) (*data.Secret, error)
	updateSecretFunc             func(*data.Version) (*data.Version, error)
}

func (m *mockBackend) GetOrCreateEncryptionKey(pubkey string) (*data.EncryptionKey, error) {
	if m.getOrCreateEncryptionKeyFunc != nil {
		return m.getOrCreateEncryptionKeyFunc(pubkey)
	}

	return &data.EncryptionKey{ID: 1, PublicKey: pubkey}, nil
}

func (m *mockBackend) CreateSecret(secret *data.Secret) (*data.Secret, error) {
	if m.createSecretFunc != nil {
		return m.createSecretFunc(secret)
	}
	secret.ID = 1
	return secret, nil
}

func (m *mockBackend) ListSecrets() ([]*data.Secret, error) {
	if m.listSecretsFunc != nil {
		return m.listSecretsFunc()
	}
	return []*data.Secret{}, nil
}

func (m *mockBackend) GetSecretVersions(title string) (*data.Secret, error) {
	if m.getSecretVersionsFunc != nil {
		return m.getSecretVersionsFunc(title)
	}
	return &data.Secret{}, nil
}

func (m *mockBackend) UpdateSecret(version *data.Version) (*data.Version, error) {
	if m.updateSecretFunc != nil {
		return m.updateSecretFunc(version)
	}
	return version, nil
}

func TestCreateSecret_Success(t *testing.T) {
	backend := &mockBackend{}
	handler := &ServerHandler{backend: backend}

	req := &connect.Request[nexusv2.CreateSecretRequest]{
		Msg: &nexusv2.CreateSecretRequest{
			Title:     "test-secret",
			Content:   []byte("encrypted-content"),
			PublicKey: "test-public-key",
		},
	}

	resp, err := handler.CreateSecret(context.Background(), req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp == nil {
		t.Fatal("Expected response, got nil")
	}

	if resp.Msg == nil {
		t.Fatal("Expected response message, got nil")
	}
}

func TestCreateSecret_MissingRequestBody(t *testing.T) {
	backend := &mockBackend{}
	handler := &ServerHandler{backend: backend}

	req := &connect.Request[nexusv2.CreateSecretRequest]{
		Msg: nil,
	}

	resp, err := handler.CreateSecret(context.Background(), req)

	if err == nil {
		t.Fatal("Expected error for missing request body")
	}

	if resp != nil {
		t.Fatal("Expected nil response for invalid request")
	}

	connectErr := err.(*connect.Error)
	if connectErr.Code() != connect.CodeInvalidArgument {
		t.Fatalf("Expected CodeInvalidArgument, got %v", connectErr.Code())
	}

	if connectErr.Message() != "missing request body" {
		t.Fatalf("Expected 'missing request body' error message, got %v", connectErr.Message())
	}
}

func TestCreateSecret_MissingTitle(t *testing.T) {
	backend := &mockBackend{}
	handler := &ServerHandler{backend: backend}

	req := &connect.Request[nexusv2.CreateSecretRequest]{
		Msg: &nexusv2.CreateSecretRequest{
			Title:     "", // Empty title
			Content:   []byte("encrypted-content"),
			PublicKey: "test-public-key",
		},
	}

	resp, err := handler.CreateSecret(context.Background(), req)

	if err == nil {
		t.Fatal("Expected error for missing title")
	}

	if resp != nil {
		t.Fatal("Expected nil response for invalid request")
	}

	connectErr := err.(*connect.Error)
	if connectErr.Code() != connect.CodeInvalidArgument {
		t.Fatalf("Expected CodeInvalidArgument, got %v", connectErr.Code())
	}

	if connectErr.Message() != "missing title" {
		t.Fatalf("Expected 'missing title' error message, got %v", connectErr.Message())
	}
}

func TestCreateSecret_MissingContent(t *testing.T) {
	backend := &mockBackend{}
	handler := &ServerHandler{backend: backend}

	req := &connect.Request[nexusv2.CreateSecretRequest]{
		Msg: &nexusv2.CreateSecretRequest{
			Title:     "test-secret",
			Content:   nil, // Missing content
			PublicKey: "test-public-key",
		},
	}

	resp, err := handler.CreateSecret(context.Background(), req)

	if err == nil {
		t.Fatal("Expected error for missing content")
	}

	if resp != nil {
		t.Fatal("Expected nil response for invalid request")
	}

	connectErr := err.(*connect.Error)
	if connectErr.Code() != connect.CodeInvalidArgument {
		t.Fatalf("Expected CodeInvalidArgument, got %v", connectErr.Code())
	}

	if connectErr.Message() != "missing content" {
		t.Fatalf("Expected 'missing content' error message, got %v", connectErr.Message())
	}
}

func TestCreateSecret_MissingPublicKey(t *testing.T) {
	backend := &mockBackend{}
	handler := &ServerHandler{backend: backend}

	req := &connect.Request[nexusv2.CreateSecretRequest]{
		Msg: &nexusv2.CreateSecretRequest{
			Title:     "test-secret",
			Content:   []byte("encrypted-content"),
			PublicKey: "", // Empty public key
		},
	}

	resp, err := handler.CreateSecret(context.Background(), req)

	if err == nil {
		t.Fatal("Expected error for missing public key")
	}

	if resp != nil {
		t.Fatal("Expected nil response for invalid request")
	}

	connectErr := err.(*connect.Error)
	if connectErr.Code() != connect.CodeInvalidArgument {
		t.Fatalf("Expected CodeInvalidArgument, got %v", connectErr.Code())
	}

	if connectErr.Message() != "missing public key" {
		t.Fatalf("Expected 'missing public key' error message, got %v", connectErr.Message())
	}
}

func TestCreateSecret_GetOrCreateEncryptionKeyError(t *testing.T) {
	backend := &mockBackend{
		getOrCreateEncryptionKeyFunc: func(string) (*data.EncryptionKey, error) {
			return nil, errors.New("database error")
		},
	}

	handler := &ServerHandler{backend: backend}

	req := &connect.Request[nexusv2.CreateSecretRequest]{
		Msg: &nexusv2.CreateSecretRequest{
			Title:     "test-secret",
			Content:   []byte("encrypted-content"),
			PublicKey: "test-public-key",
		},
	}

	resp, err := handler.CreateSecret(context.Background(), req)

	if err == nil {
		t.Fatal("Expected error from GetOrCreateEncryptionKey")
	}

	if resp != nil {
		t.Fatal("Expected nil response for backend error")
	}

	connectErr := err.(*connect.Error)
	if connectErr.Code() != connect.CodeInternal {
		t.Fatalf("Expected CodeInternal, got %v", connectErr.Code())
	}
}

func TestCreateSecret_CreateSecretError(t *testing.T) {
	backend := &mockBackend{
		createSecretFunc: func(*data.Secret) (*data.Secret, error) {
			return nil, errors.New("database error")
		},
	}

	handler := &ServerHandler{backend: backend}

	req := &connect.Request[nexusv2.CreateSecretRequest]{
		Msg: &nexusv2.CreateSecretRequest{
			Title:     "test-secret",
			Content:   []byte("encrypted-content"),
			PublicKey: "test-public-key",
		},
	}

	resp, err := handler.CreateSecret(context.Background(), req)

	if err == nil {
		t.Fatal("Expected error from CreateSecret")
	}

	if resp != nil {
		t.Fatal("Expected nil response for backend error")
	}

	connectErr := err.(*connect.Error)
	if connectErr.Code() != connect.CodeInternal {
		t.Fatalf("Expected CodeInternal, got %v", connectErr.Code())
	}
}

func TestCreateSecret_VerifySecretStructure(t *testing.T) {
	var capturedSecret *data.Secret
	backend := &mockBackend{
		createSecretFunc: func(secret *data.Secret) (*data.Secret, error) {
			capturedSecret = secret
			secret.ID = 1
			return secret, nil
		},
	}

	handler := &ServerHandler{backend: backend}

	expectedTitle := "test-secret"
	expectedContent := []byte("encrypted-content")
	expectedPublicKey := "test-public-key"

	req := &connect.Request[nexusv2.CreateSecretRequest]{
		Msg: &nexusv2.CreateSecretRequest{
			Title:     expectedTitle,
			Content:   expectedContent,
			PublicKey: expectedPublicKey,
		},
	}

	_, err := handler.CreateSecret(context.Background(), req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if capturedSecret == nil {
		t.Fatal("Expected secret to be passed to backend")
	}

	if capturedSecret.Title != expectedTitle {
		t.Fatalf("Expected title %q, got %q", expectedTitle, capturedSecret.Title)
	}

	if len(capturedSecret.Versions) != 1 {
		t.Fatalf("Expected 1 secret version, got %d", len(capturedSecret.Versions))
	}

	version := capturedSecret.Versions[0]
	if version.Version != 1 {
		t.Fatalf("Expected version 1, got %d", version.Version)
	}

	if string(version.Content) != string(expectedContent) {
		t.Fatalf("Expected content %q, got %q", expectedContent, version.Content)
	}

	if version.Key.PublicKey != expectedPublicKey {
		t.Fatalf("Expected public key %q, got %q", expectedPublicKey, version.Key.PublicKey)
	}

	if capturedSecret.CreatedSignature.CreatedBy != "unknown" {
		t.Fatalf("Expected CreatedBy 'unknown', got %q", capturedSecret.CreatedSignature.CreatedBy)
	}

	if version.CreatedSignature.CreatedBy != "unknown" {
		t.Fatalf("Expected version CreatedBy 'unknown', got %q", version.CreatedSignature.CreatedBy)
	}
}
