// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package version

import (
	"fmt"

	version "github.com/caarlos0/go-version"
)

func VersionInfo() version.Info {
	appDetails := version.WithAppDetails("kv2", "Encrypted secrets management for the homelab.", "git.huggins.io/kv2")
	return version.GetVersionInfo(appDetails)
}

func Line() string {
	return fmt.Sprintf("%s %s (%s) for %s", VersionInfo().Name, VersionInfo().GitVersion, VersionInfo().GitCommit, VersionInfo().Platform)
}

func Seymour() string {
	return fmt.Sprintf("see more at %s", VersionInfo().URL)
}
