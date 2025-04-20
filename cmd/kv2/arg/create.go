// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package arg

import (
	"fmt"
	"os"

	"connectrpc.com/connect"
	secretsv1 "git.huggins.io/kv2/api/secrets/v1"
	"git.huggins.io/kv2/internal/cli"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create <key>",
	Short: "Create a new secret",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		req := &secretsv1.CreateSecretRequest{
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

		if req.Value == nil {
			fmt.Print("Enter value: ")
			fmt.Scanln(&req.Value)
		}

		if _, err := kv2.CreateSecret(cmd.Context(), &connect.Request[secretsv1.CreateSecretRequest]{Msg: req}); err != nil {
			cli.PrintErrorOutput(jsonOutput, err)
			os.Exit(1)
		}
	},
}

func init() {
	createCmd.Flags().String("from-literal", "", "create secret from literal value")
	createCmd.Flags().String("from-file", "", "create secret from file")

	createCmd.MarkFlagsMutuallyExclusive("from-literal", "from-file")

	RootCmd.AddCommand(createCmd)
}
