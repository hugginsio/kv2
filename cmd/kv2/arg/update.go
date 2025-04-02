// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package arg

import (
	"os"

	"git.huggins.io/kv2/api"
	"git.huggins.io/kv2/internal/cli"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update <key>",
	Short: "Update a secret to a new version",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		req := api.UpdateSecretRequest{
			Key: args[0],
		}

		if val, _ := cmd.Flags().GetString("from-literal"); val != "" {
			req.Value = []byte(val)
		}

		if val, _ := cmd.Flags().GetString("from-file"); val != "" {
			bytes, err := os.ReadFile(val)
			if err != nil {
				cli.PrintErrorOutput(jsonOutput, err)
				os.Exit(1)
			}

			req.Value = bytes
		}

		if err := kv2.Update(req); err != nil {
			cli.PrintErrorOutput(jsonOutput, err)
			os.Exit(1)
		}
	},
}

func init() {
	updateCmd.Flags().String("from-literal", "", "create secret from literal value")
	updateCmd.Flags().String("from-file", "", "create secret from file")

	updateCmd.MarkFlagsMutuallyExclusive("from-literal", "from-file")
	updateCmd.MarkFlagsOneRequired("from-literal", "from-file")

	rootCmd.AddCommand(updateCmd)
}
