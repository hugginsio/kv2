// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package database

import secretsv1 "git.huggins.io/kv2/api/secrets/v1"

// Database provides methods for interacting with a secrets database.
type Database interface {
	List() (*secretsv1.ListSecretsResponse, error)
	Create(*secretsv1.CreateSecretRequest) error
	Read(*secretsv1.GetSecretRequest) (*secretsv1.GetSecretResponse, error)
	Update(*secretsv1.UpdateSecretRequest) error
	Delete(*secretsv1.DeleteSecretRequest) error
	Revert(*secretsv1.RevertSecretRequest) error
}

type Configuration struct {
	Dsn string
}

type SecretRecord struct {
	ID  uint
	Key string
}

type ValueRecord struct {
	ID             uint
	SecretRecordID uint
	Value          []byte
	Version        uint32
}
