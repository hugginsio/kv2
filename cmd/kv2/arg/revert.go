// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package arg

import (
	"os"

	"git.huggins.io/kv2/api"
	"git.huggins.io/kv2/internal/cli"
	"github.com/spf13/cobra"
)

var revertCmd = &cobra.Command{
	Use:   "revert <key>",
	Short: "Revert a secret to a previous version",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := kv2.Revert(api.RevertSecretRequest{Key: args[0]}); err != nil {
			cli.PrintErrorOutput(jsonOutput, err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(revertCmd)
}
