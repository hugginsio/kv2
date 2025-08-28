// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package data

// SqliteBackend represents the SQLite implementation of the Backend.
type SqliteBackend struct{}

func NewSqliteBackend() (error, *SqliteBackend) {
	return nil, &SqliteBackend{}
}
