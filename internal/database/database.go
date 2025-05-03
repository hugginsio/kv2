// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package database

import (
	"time"

	secretsv1 "git.huggins.io/kv2/api/secrets/v1"
)

// Database provides methods for interacting with a secrets database.
type Database interface {
	List() (*secretsv1.ListSecretsResponse, error)
	Create(*secretsv1.CreateSecretRequest) error
	Read(*secretsv1.GetSecretRequest) (*secretsv1.GetSecretResponse, error)
	Update(*secretsv1.UpdateSecretRequest) (*secretsv1.UpdateSecretResponse, error)
	Delete(*secretsv1.DeleteSecretRequest) error
	Revert(*secretsv1.RevertSecretRequest) (*secretsv1.RevertSecretResponse, error)
}

type Configuration struct {
	Dsn string
}

// Secret is a collection of Versions uniquely identified by a key. Only one version can be active at a time.
type Secret struct {
	ID        uint
	Key       string    // the name uniquely identifying the string
	Active    uint      // the currently active Version of the Secret
	Versions  []Version // all versions of the provided secret
	CreatedAt time.Time // the time the Secret was created
	UpdatedAt time.Time // the time the Secret was last updated
	CreatedBy string    // the peer who created the Secret
	UpdatedBy string    // the peer who last updated the Secret
}

// Version represents a single instance of a Secret value, including the Recipient who can decode it.
type Version struct {
	ID        uint
	Value     []byte    // the encrypted value of the Version
	Recipient Recipient // the Recipient of the Version
	CreatedAt time.Time // the time the Version was created
	CreatedBy string    // the peer who created the Version
}

// Recipient represents the identity that can decode a Secret Version.
type Recipient struct {
	ID        uint
	Value     []byte // the version Recipient, typically a public key
	CreatedAt time.Time
	CreatedBy string // the peer who created the Recipient
}
