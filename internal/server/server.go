// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package server

import (
	"encoding/json"
	"io"
	"net/http"

	"git.huggins.io/kv2/api"
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

	config.Mux.HandleFunc("/secrets", server.list)
	config.Mux.HandleFunc("/secrets/create", server.create)
	config.Mux.HandleFunc("/secrets/read", server.read)
	config.Mux.HandleFunc("/secrets/update", server.update)
	config.Mux.HandleFunc("/secrets/delete", server.delete)
	config.Mux.HandleFunc("/secrets/revert", server.revert)
	config.Mux.HandleFunc("/secrets/backup", server.createBackup)

	return server
}

// Lists all secrets in the database.
func (hs *HttpServer) list(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	secrets, err := hs.database.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(secrets)
}

// Create a new secret in the database.
func (hs *HttpServer) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	var request api.CreateSecretRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	enc, err := hs.crypto.Encrypt(request.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	request.Value = enc

	if err := hs.database.Create(request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Retrieve a single secret from the database.
func (hs *HttpServer) read(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	var request api.ReadSecretRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	secret, err := hs.database.Read(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dec, err := hs.crypto.Decrypt(secret.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	secret.Value = dec

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(secret)
}

// Update an existing secret with a new version.
func (hs *HttpServer) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	var request api.UpdateSecretRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := hs.database.Update(request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete a secret and all of its versions.
func (hs *HttpServer) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	var request api.DeleteSecretRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := hs.database.Delete(request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Revert a secret to the previous version.
func (hs *HttpServer) revert(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	var request api.RevertSecretRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := hs.database.Revert(request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Create a backup using the configured provider.
func (hs *HttpServer) createBackup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	var request api.BackupRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil && err != io.EOF {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if request.Name == "" {
		request.Name = "kv2.db"
		return
	}

	if err := hs.backup.Backup(request.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
