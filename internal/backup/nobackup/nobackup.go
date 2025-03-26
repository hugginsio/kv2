// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package nobackup

import "errors"

var errNoBackupProvider = errors.New("no backup provider configured")

func Restore() error {
	return errNoBackupProvider
}

func Backup(_ string) error {
	return errNoBackupProvider
}
