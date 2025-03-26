// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package nobackup

import "errors"

var errNoBackupProvider = errors.New("no backup provider configured")

type NoBackup struct{}

func Initialize() *NoBackup {
	return &NoBackup{}
}

func (n *NoBackup) Restore() error {
	return errNoBackupProvider
}

func (n *NoBackup) Backup(_ string) error {
	return errNoBackupProvider
}
