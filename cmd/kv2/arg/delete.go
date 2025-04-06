// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package arg

import (
	"os"

	"connectrpc.com/connect"
	secretsv1 "git.huggins.io/kv2/api/secrets/v1"
	"git.huggins.io/kv2/internal/cli"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <key>",
	Short: "Delete a secret and all its versions",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		req := &secretsv1.DeleteSecretRequest{
			Key: args[0],
		}

		if _, err := kv2.DeleteSecret(cmd.Context(), &connect.Request[secretsv1.DeleteSecretRequest]{Msg: req}); err != nil {
			cli.PrintErrorOutput(jsonOutput, err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
