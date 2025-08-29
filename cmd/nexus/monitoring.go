// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"log/slog"
	"net/http"
)

func ServeHealthEndpoint() {
	// TODO: get knowledge of tailscale connection where configured for tailnet access
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})

		port := "8081"
		slog.Info("health listening on", "port", port)
		if err := http.ListenAndServe(":"+port, mux); err != nil {
			slog.Error("health server failed", "error", err)
		}
	}()
}
