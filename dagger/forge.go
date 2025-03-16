package main

import (
	"context"
	"dagger/kv-2/internal/dagger"
	"strings"
)

// Run all pull request checks.
func (m *Kv2) PullRequest(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
) (string, error) {
	lint, err := m.Lint(ctx, source)
	if err != nil {
		return "", err
	}

	test, err := m.Test(ctx, source)
	if err != nil {
		return "", err
	}

	return strings.Join([]string{lint, test}, "\n"), nil
}
