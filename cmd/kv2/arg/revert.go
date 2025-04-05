// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package arg

import (
	"os"

	"github.com/spf13/cobra"
)

var revertCmd = &cobra.Command{
	Use:   "revert <key>",
	Short: "Revert a secret to a previous version",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(revertCmd)
}
