// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

// Package data provides models and implementations for database interactions.
package data

import (
	"time"
)

type CreatedSignature struct {
	CreatedAt time.Time `gorm:"not null;autoUpdateTime"` // The time the CreatedSignature was created.
	CreatedBy string    `gorm:"not null"`                // The IP address of the client that created the Secret.
}

type EncryptionKey struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;uniqueIndex;foreignKey"` // Unique numeric identifier for the EncryptionKey. PK.
	PublicKey string `gorm:"not null"`                                        // The public key text.
	CreatedSignature
}

type Secret struct {
	ID            uint            `gorm:"primaryKey;autoIncrement;"`   // Unique numeric identifier for the Secret. PK.
	Title         string          `gorm:"not null;unique;uniqueIndex"` // Unique title of the Secret.
	SecretVersion []SecretVersion `gorm:"foreignKey:SecretID"`         // All versions of the Secret.
	CreatedSignature
}

type SecretVersion struct {
	ID       uint          `gorm:"primaryKey;autoIncrement"`        // Unique numeric identifier for the SecretVersion. PK.
	SecretID uint          `gorm:"not null;foreignKey;uniqueIndex"` // The ID of the Secret this version belongs to.
	Version  uint          `gorm:"not null"`                        // The version of the Secret.
	Content  []byte        `gorm:"not null"`                        // The encrypted contents of the SecretVersion.
	Key      EncryptionKey `gorm:"not null;foreignKey:ID"`          // The public key used to encrypt the SecretVersion contents.
	CreatedSignature
}

type Backup struct {
	// TODO
}

// Backend represents the methods available through the underlying storage system.
type Backend interface {
	ListSecrets() ([]*Secret, error)                                // List all secrets.
	CreateSecret(*Secret) (*Secret, error)                          // Create a new secret.
	GetSecretVersions(string) (*Secret, error)                      // Retrieve all versions of an existing secret by title.
	UpdateSecret(*SecretVersion) (*SecretVersion, error)            // Update an existing secret with a new version.
	GetOrCreateEncryptionKey(pubkey string) (*EncryptionKey, error) // Retrieve or create the encryption key details.
}
