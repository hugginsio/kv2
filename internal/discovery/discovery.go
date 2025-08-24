// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

// Package discovery contains methods for environment and configuration discovery.
package discovery

import (
	"os"
	"strings"
)

// Determines if the process is running in development mode (KV2_DEVEL is set to true).
func IsDevel() bool {
	return os.Getenv("KV2_DEVEL") == "true"
}

// Determines if the process is running in a container (via dockerenv, cgroups, or env vars).
func IsContainerized() bool {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}

	data, err := os.ReadFile("/proc/1/cgroup")
	if err != nil {
		return false
	}

	content := string(data)
	matches := []string{"docker", "containerd", "/lxc/", "kubepods"}
	for _, match := range matches {
		if strings.Contains(content, match) {
			return true
		}
	}

	containerVars := []string{"container", "DOCKER_CONTAINER", "KUBERNETES_SERVICE_HOST"}
	for _, envVar := range containerVars {
		if os.Getenv(envVar) != "" {
			return true
		}
	}

	return false
}
