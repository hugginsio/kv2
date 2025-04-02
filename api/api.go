// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package api

import "errors"

var ErrorCannotRevert error = errors.New("cannot rollback version 1 of secret")
var ErrorSecretAlreadyExists error = errors.New("secret already exists")
var ErrorSecretNotFound error = errors.New("secret not found")

type SecretValue []byte
type SecretVersion uint16

type HealthResponse struct {
	Status string
}

type Secret struct {
	Key     string
	Value   SecretValue
	Version SecretVersion
}

type ErrorResponse struct {
	Message string
}

type ListSecretResponse struct {
	Key      string
	Versions []SecretVersion
}

type CreateSecretRequest struct {
	Key   string
	Value SecretValue
}

type ReadSecretRequest struct {
	Key string
}

type UpdateSecretRequest struct {
	Key   string
	Value SecretValue
}

type DeleteSecretRequest struct {
	Key string
}

type RevertSecretRequest struct {
	Key string
}

type BackupRequest struct {
	Name string
}
