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

// Scan for vulnerabilities and license violations.
func (m *Kv2) Scan(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
) (string, error) {
	return dag.Container().
		From("ghcr.io/google/osv-scanner:v2.0.1").
		WithWorkdir("/go/src/").
		WithMountedDirectory("/go/src/", source).
		WithExec([]string{"/osv-scanner", "scan", "--call-analysis=go", "-r", "/go/src/go.mod"}).
		Stdout(ctx)
}
