// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

// Package data provides models and implementations for database interactions.
package data

import (
	"time"
)

type CreatedSignature struct {
	CreatedAt time.Time `gorm:"not null;autoUpdateTime"` // The time the CreatedSignature was created.
	CreatedBy string    `gorm:"not null"`                // The program code of the client that created the Secret.
}

type EncryptionKey struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;uniqueIndex"` // Unique numeric identifier for the EncryptionKey. PK.
	PublicKey string `gorm:"not null"`                             // The public key text.
	CreatedSignature
}

type Secret struct {
	ID            uint            `gorm:"primaryKey;autoIncrement;"`   // Unique numeric identifier for the Secret. PK.
	Title         string          `gorm:"not null;unique;uniqueIndex"` // Unique title of the Secret.
	SecretVersion []SecretVersion `gorm:"foreignKey:SecretID"`         // All versions of the Secret.
	CreatedSignature
}

type SecretVersion struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`        // Unique numeric identifier for the SecretVersion. PK.
	SecretID uint   `gorm:"not null;foreignKey;uniqueIndex"` // The ID of the Secret this version belongs to.
	Version  uint   `gorm:"not null"`                        // The version of the Secret.
	Content  []byte `gorm:"not null"`                        // The encrypted contents of the SecretVersion.
	CreatedSignature
}

type Backup struct {
	// TODO
}

// Backend represents the methods available through the underlying storage system.
type Backend interface {
	CreateSecret(*Secret) (*Secret, error)               // Create a new secret.
	GetSecretVersions(string) (*Secret, error)           // Retrieve all versions of an existing secret by title.
	UpdateSecret(*SecretVersion) (*SecretVersion, error) // Update an existing secret with a new version.
}
