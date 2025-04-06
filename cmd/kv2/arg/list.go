// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package arg

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"connectrpc.com/connect"
	secretsv1 "git.huggins.io/kv2/api/secrets/v1"
	"git.huggins.io/kv2/internal/cli"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available secrets and versions",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := kv2.ListSecrets(context.Background(), &connect.Request[secretsv1.ListSecretsRequest]{})
		if err != nil {
			cli.PrintErrorOutput(jsonOutput, err)
			os.Exit(1)
		}

		if jsonOutput {
			json.NewEncoder(os.Stdout).Encode(res)
			return
		}

		var data [][]string

		for _, s := range res.Msg.Secrets {
			data = append(data, []string{s.Key, fmt.Sprintf("%d", len(s.Version))})
		}

		cli.PrintTable([]string{"KEY", "VERSION"}, data)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
