// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package arg

import (
	"os"

	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup the secrets database",
	Run: func(cmd *cobra.Command, args []string) {
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
}
