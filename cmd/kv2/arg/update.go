// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package arg

import (
	"os"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update <key>",
	Short: "Update a secret to a new version",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		os.Exit(1)
	},
}

func init() {
	updateCmd.Flags().String("from-literal", "", "create secret from literal value")
	updateCmd.Flags().String("from-file", "", "create secret from file")

	updateCmd.MarkFlagsMutuallyExclusive("from-literal", "from-file")
	updateCmd.MarkFlagsOneRequired("from-literal", "from-file")

	rootCmd.AddCommand(updateCmd)
}
