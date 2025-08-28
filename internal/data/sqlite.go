// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package data

import (
	"log/slog"

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
	slog.Debug("opening database", "path", config.Path)

	db, err := gorm.Open(sqlite.Open(config.Path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&EncryptionKey{}, &Secret{}, &SecretVersion{}); err != nil {
		return nil, err
	}

	return &SqliteBackend{sql: db}, nil
}
