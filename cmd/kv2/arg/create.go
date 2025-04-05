// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package arg

import (
	"os"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create <key>",
	Short: "Create a new secret",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		os.Exit(1)
	},
}

func init() {
	createCmd.Flags().String("from-literal", "", "create secret from literal value")
	createCmd.Flags().String("from-file", "", "create secret from file")

	createCmd.MarkFlagsMutuallyExclusive("from-literal", "from-file")
	createCmd.MarkFlagsOneRequired("from-literal", "from-file")

	rootCmd.AddCommand(createCmd)
}
