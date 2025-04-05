// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package arg

import (
	"fmt"
	"os"

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
		os.Exit(1)
	},
}

func init() {
	readCmd.Flags().BoolP("decode", "d", false, "attempt to decode value base64")

	rootCmd.AddCommand(readCmd)
}
