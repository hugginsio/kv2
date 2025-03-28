// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package version

import (
	version "github.com/caarlos0/go-version"
)

func VersionInfo() version.Info {
	appDetails := version.WithAppDetails("kv2", "Encrypted secrets management for the homelab.", "git.huggins.io/kv2")
	return version.GetVersionInfo(appDetails)
}
