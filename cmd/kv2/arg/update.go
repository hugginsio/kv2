// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package arg

import (
	"encoding/json"
	"fmt"
	"os"

	"connectrpc.com/connect"
	secretsv1 "git.huggins.io/kv2/api/secrets/v1"
	"git.huggins.io/kv2/internal/cli"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update <key>",
	Short: "Update a secret to a new version",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		req := &secretsv1.UpdateSecretRequest{
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

		res, err := kv2.UpdateSecret(cmd.Context(), &connect.Request[secretsv1.UpdateSecretRequest]{Msg: req})
		if err != nil {
			cli.PrintErrorOutput(jsonOutput, err)
			os.Exit(1)
		}

		if jsonOutput {
			json.NewEncoder(os.Stdout).Encode(res.Msg)
			return
		}

		fmt.Println("Secret updated to version", res.Msg.Version)
	},
}

func init() {
	updateCmd.Flags().String("from-literal", "", "create secret from literal value")
	updateCmd.Flags().String("from-file", "", "create secret from file")

	updateCmd.MarkFlagsMutuallyExclusive("from-literal", "from-file")
	updateCmd.MarkFlagsOneRequired("from-literal", "from-file")

	rootCmd.AddCommand(updateCmd)
}
