// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package arg

import (
	"context"
	"fmt"
	"os"

	"connectrpc.com/connect"
	secretsv1 "git.huggins.io/kv2/api/secrets/v1"
	"git.huggins.io/kv2/internal/cli"
	"git.huggins.io/kv2/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := kv2.ApplicationVersionInfo(context.Background(), &connect.Request[secretsv1.ApplicationVersionInfoRequest]{})
		if err != nil {
			cli.PrintErrorOutput(jsonOutput, err)
			os.Exit(1)
		}

		// TODO: version info JSON

		// if jsonOutput {
		// 	json.NewEncoder(os.Stdout).Encode(version.VersionInfo())
		// 	return
		// }

		fmt.Println("CLIENT")
		fmt.Println(version.VersionInfo())
		fmt.Println("SERVER")
		fmt.Println(res.Msg.GetInfo())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
