// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package backup

import "errors"

type CloudBackup interface {
	Restore() bool
}

func DetermineStorageProvider(path string) (*CloudBackup, error) {
	// TODO: impl (switch? fail on invalid?)
	// there should be a config string by the time we get here
	return nil, errors.New("not implemented")
}
