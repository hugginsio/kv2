// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package arg

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"connectrpc.com/connect"
	secretsv1 "git.huggins.io/kv2/api/secrets/v1"
	"git.huggins.io/kv2/internal/cli"
	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read <secret>",
	Short: "Read the latest version of a secret",
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Flag("decode").Changed && cmd.Flag("json").Changed {
			return fmt.Errorf("--decode and --json are mutually exclusive")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		req := &secretsv1.GetSecretRequest{
			Key: args[0],
		}

		res, err := kv2.GetSecret(cmd.Context(), &connect.Request[secretsv1.GetSecretRequest]{Msg: req})
		if err != nil {
			cli.PrintErrorOutput(jsonOutput, err)
			os.Exit(1)
		}

		if jsonOutput {
			json.NewEncoder(os.Stdout).Encode(res.Msg.Secret)
			return
		}

		if cmd.Flag("decode").Value.String() == "true" {
			fmt.Println(string(res.Msg.Secret.Value))
			return
		}

		fmt.Println(base64.StdEncoding.EncodeToString(res.Msg.Secret.Value))
	},
}

func init() {
	readCmd.Flags().BoolP("decode", "d", false, "attempt to decode value base64")

	rootCmd.AddCommand(readCmd)
}
