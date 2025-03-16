package main

import (
	"context"
	"dagger/kv-2/internal/dagger"
)

// Lint the project.
func (m *Kv2) Lint(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
) (string, error) {
	return m.devEnv(ctx, source, nil).WithExec([]string{"go", "vet", "./..."}).Stdout(ctx)
}

// Run all tests.
func (m *Kv2) Test(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
) (string, error) {
	return m.devEnv(ctx, source, nil).WithExec([]string{"go", "test", "-v", "./..."}).Stdout(ctx)
}
