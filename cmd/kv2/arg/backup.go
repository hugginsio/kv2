// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package arg

import (
	"context"

	"connectrpc.com/connect"
	secretsv1 "git.huggins.io/kv2/api/secrets/v1"
	"git.huggins.io/kv2/internal/cli"
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup the secrets database",
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Flag("name").Value.String()
		_, err := kv2.Backup(context.Background(), &connect.Request[secretsv1.BackupRequest]{Msg: &secretsv1.BackupRequest{Name: &name}})

		if err != nil {
			cli.PrintErrorOutput(jsonOutput, err)
		}
	},
}

func init() {
	backupCmd.Flags().StringP("name", "n", "kv2.db", "name of the backup")

	rootCmd.AddCommand(backupCmd)
}
