// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package arg

import (
	"encoding/json"
	"fmt"
	"os"

	"git.huggins.io/kv2/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		if jsonOutput {
			json.NewEncoder(os.Stdout).Encode(version.VersionInfo())
			return
		}

		fmt.Println(version.VersionInfo())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
