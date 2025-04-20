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
	"git.huggins.io/kv2/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := kv2.ApplicationVersionInfo(context.Background(), &connect.Request[secretsv1.ApplicationVersionInfoRequest]{})

		if jsonOutput {
			version := struct {
				Client any `json:"client"`
				Server any `json:"server"`
			}{Client: version.VersionInfo()}

			if err == nil {
				version.Server = res.Msg.GetInfo()
			}

			json.NewEncoder(os.Stdout).Encode(version)
			return
		}

		fmt.Println("CLIENT")
		fmt.Println("    GitVersion:", version.VersionInfo().GitVersion)
		fmt.Println("    GoVersion: ", version.VersionInfo().GoVersion)
		fmt.Println("    Platform:  ", version.VersionInfo().Platform)

		fmt.Println("")
		fmt.Println("SERVER")
		if err != nil {
			fmt.Fprintln(os.Stderr, "   ", err)
			return
		}

		fmt.Println("    GitVersion:", res.Msg.GetInfo().GitVersion)
		fmt.Println("    GoVersion: ", res.Msg.GetInfo().GoVersion)
		fmt.Println("    Platform:  ", res.Msg.GetInfo().Platform)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
