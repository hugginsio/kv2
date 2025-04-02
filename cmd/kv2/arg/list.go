// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package arg

import (
	"encoding/json"
	"fmt"
	"os"

	"git.huggins.io/kv2/internal/cli"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available secrets and versions",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := kv2.List()
		if err != nil {
			cli.PrintErrorOutput(jsonOutput, err)
			os.Exit(1)
		}

		if jsonOutput {
			json.NewEncoder(os.Stdout).Encode(res)
			return
		}

		var data [][]string

		for _, s := range res {
			data = append(data, []string{s.Key, fmt.Sprintf("%d", len(s.Versions))})
		}

		cli.PrintTable([]string{"KEY", "VERSION"}, data)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
