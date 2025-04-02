// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package arg

import (
	"os"

	"git.huggins.io/kv2/api"
	"git.huggins.io/kv2/internal/cli"
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup the secrets database",
	Run: func(cmd *cobra.Command, args []string) {
		if err := kv2.Backup(api.BackupRequest{Name: "kv2.db"}); err != nil {
			cli.PrintErrorOutput(jsonOutput, err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
}
