// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package data

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// SqliteBackend represents the SQLite implementation of the Backend.
type SqliteBackend struct {
	sql *gorm.DB
}

type SqliteConfiguration struct {
	Path string // The absolute path to where the database is stored.
}

func NewSqliteBackend(config *SqliteConfiguration) (*SqliteBackend, error) {
	db, err := gorm.Open(sqlite.Open(config.Path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&EncryptionKey{}, &Secret{}, &Version{}); err != nil {
		return nil, err
	}

	return &SqliteBackend{sql: db}, nil
}

func (m *SqliteBackend) ListSecrets() ([]*Secret, error) {
	var secrets []*Secret
	// TODO: add sort to preload and find
	if err := m.sql.Preload("Versions").Find(&secrets).Error; err != nil {
		return nil, err
	}

	return secrets, nil
}

func (m *SqliteBackend) CreateSecret(secret *Secret) (*Secret, error) {
	tx := m.sql.Save(&secret)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return secret, nil
}

func (m *SqliteBackend) GetSecretVersions(string) (*Secret, error) {
	panic("unimplemented")
}

func (m *SqliteBackend) UpdateSecret(*Version) (*Version, error) {
	panic("unimplemented")
}

func (m *SqliteBackend) GetOrCreateEncryptionKey(pubkey string) (*EncryptionKey, error) {
	key := EncryptionKey{
		PublicKey: pubkey,
	}

	tx := m.sql.FirstOrCreate(&key, "public_key = ?", pubkey)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &key, nil
}
