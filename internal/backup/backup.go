// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package backup

import (
	"errors"
	"regexp"
	"strings"

	"git.huggins.io/kv2/internal/backup/gcs"
)

type CloudBackup interface {
	Restore() error
	Backup(name string) error
}

func DetermineStorageProvider(value string, destinationPath string) (*CloudBackup, error) {
	r := regexp.MustCompile(`[^:]*`)
	match := r.FindString(value)
	location := strings.TrimPrefix(value, match+"://")

	var cloudBackup CloudBackup
	switch match {
	case "gcs":
		cloudBackup = gcs.Initialize(gcs.Configuration{BucketName: location, FilePath: destinationPath})
	default:
		return nil, errors.New("invalid provider \"" + match + "\"")
	}

	return &cloudBackup, nil
}
