// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package database

import (
	"git.huggins.io/kv2/api"
)

// Database provides methods for interacting with a secrets database.
type Database interface {
	List() ([]api.ListSecretResponse, error)
	Create(api.CreateSecretRequest) error
	Read(api.ReadSecretRequest) (api.Secret, error)
	Update(api.UpdateSecretRequest) error
	Delete(api.DeleteSecretRequest) error
	Revert(api.RevertSecretRequest) error
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
	Version        api.SecretVersion
}

type ServerMetadata struct {
	CreatedAt int64
}
