// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package version

import (
	secretsv1 "git.huggins.io/kv2/api/secrets/v1"
	version "github.com/caarlos0/go-version"
)

func VersionInfo() *secretsv1.ApplicationVersionInfo {
	v := version.GetVersionInfo()

	return &secretsv1.ApplicationVersionInfo{
		GitVersion: v.GitVersion,
		GoVersion:  v.GoVersion,
		Platform:   v.Platform,
	}
}
