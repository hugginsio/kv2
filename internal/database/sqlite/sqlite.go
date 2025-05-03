// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package sqlite

import (
	"gorm.io/gorm"
)

type Configuration struct {
	Dsn string
}

type SqliteDatabase struct {
	sql *gorm.DB
}

// Initialize a sqlite database for secret storage.
// func Initialize(configuration Configuration) (*SqliteDatabase, error) {
// 	log.Info().Str("path", configuration.Dsn).Msg("initializing database state")

// 	db, err := gorm.Open(sqlite.Open(configuration.Dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := db.AutoMigrate(&database.SecretRecord{}, &database.ValueRecord{}); err != nil {
// 		return nil, err
// 	}

// 	return &SqliteDatabase{sql: db}, nil
// }
