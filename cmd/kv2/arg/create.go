// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package arg

import (
	"os"

	"git.huggins.io/kv2/api"
	"git.huggins.io/kv2/internal/cli"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create <key>",
	Short: "Create a new secret",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		req := api.CreateSecretRequest{
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

		if err := kv2.Create(req); err != nil {
			cli.PrintErrorOutput(jsonOutput, err)
			os.Exit(1)
		}
	},
}

func init() {
	createCmd.Flags().String("from-literal", "", "create secret from literal value")
	createCmd.Flags().String("from-file", "", "create secret from file")

	createCmd.MarkFlagsMutuallyExclusive("from-literal", "from-file")
	createCmd.MarkFlagsOneRequired("from-literal", "from-file")

	rootCmd.AddCommand(createCmd)
}
