// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package arg

import (
	"errors"
	"net"
	"net/http"
	"os"
	"time"

	"connectrpc.com/connect"
	"git.huggins.io/kv2/api/secrets/v1/secretsv1connect"
	"github.com/spf13/cobra"
)

var jsonOutput bool
var kv2 secretsv1connect.Kv2ServiceClient
var rootCmd = &cobra.Command{
	Use:   "kv2",
	Short: "kv2 provides an interface for your secrets manager",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		serverUrlEnv, exists := os.LookupEnv("KV2_SERVER_URL")
		if !exists {
			// TODO: attempt to use tailscale CLI to automatically determine URL?
			return errors.New("KV2_SERVER_URL not found in environment")
		}

		opts := connect.WithClientOptions(
			connect.WithCompressMinBytes(1280),
		)

		client := &http.Client{
			CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Timeout: 20 * time.Second,
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout: 3 * time.Second,
				}).Dial,
				TLSHandshakeTimeout:   3 * time.Second,
				ResponseHeaderTimeout: 3 * time.Second,
				ExpectContinueTimeout: 3 * time.Second,
			},
		}

		kv2 = secretsv1connect.NewKv2ServiceClient(client, serverUrlEnv, opts)

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&jsonOutput, "json", "j", false, "JSON output")
	rootCmd.PersistentFlags().String("config", "", "Path to config file")
}
