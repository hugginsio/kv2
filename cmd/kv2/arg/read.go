// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package arg

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"git.huggins.io/kv2/api"
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
		res, err := kv2.Read(api.ReadSecretRequest{Key: args[0]})
		if err != nil {
			cli.PrintErrorOutput(jsonOutput, err)
			os.Exit(1)
		}

		if jsonOutput {
			json.NewEncoder(os.Stdout).Encode(res)
			return
		}

		if cmd.Flag("decode").Value.String() == "true" {
			fmt.Println(string(res.Value))
			return
		}

		fmt.Println(base64.StdEncoding.EncodeToString(res.Value))
	},
}

func init() {
	readCmd.Flags().BoolP("decode", "d", false, "attempt to decode value base64")

	rootCmd.AddCommand(readCmd)
}
