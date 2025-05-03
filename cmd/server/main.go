// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"git.huggins.io/kv2/internal/version"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Str("version", version.VersionInfo().GitVersion).Str("platform", version.VersionInfo().Platform).Msg("kv2 is starting")
}
