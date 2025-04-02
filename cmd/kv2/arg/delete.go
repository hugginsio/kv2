// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package arg

import (
	"git.huggins.io/kv2/api"
	"git.huggins.io/kv2/internal/cli"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <key>",
	Short: "Delete a secret and all its versions",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := kv2.Delete(api.DeleteSecretRequest{Key: args[0]}); err != nil {
			cli.PrintErrorOutput(jsonOutput, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
