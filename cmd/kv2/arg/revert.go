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

var revertCmd = &cobra.Command{
	Use:   "revert <key>",
	Short: "Revert a secret to a previous version",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		req := secretsv1.RevertSecretRequest{
			Key: args[0],
		}

		res, err := kv2.RevertSecret(cmd.Context(), &connect.Request[secretsv1.RevertSecretRequest]{Msg: &req})
		if err != nil {
			cli.PrintErrorOutput(jsonOutput, err)
			os.Exit(1)
		}

		if jsonOutput {
			json.NewEncoder(os.Stdout).Encode(res.Msg)
			return
		}

		fmt.Println("Secret reverted to version", res.Msg.Version)
	},
}

func init() {
	rootCmd.AddCommand(revertCmd)
}
