package main

import (
	"context"
	"dagger/kv-2/internal/dagger"
	"time"
)

type Kv2 struct{}

// Build a ready-to-use development environment.
func (m *Kv2) devEnv(
	ctx context.Context,
	source *dagger.Directory,
) *dagger.Container {
	return dag.Container().
		From("golang:1.24-alpine").
		WithDirectory("/go/src/", source).
		WithMountedCache("/go/pkg/mod/", dag.CacheVolume("go-mod-124")).
		WithEnvVariable("GOMODCACHE", "/go/pkg/mod").
		WithMountedCache("/go/build-cache", dag.CacheVolume("go-build-124")).
		WithEnvVariable("GOCACHE", "/go/build-cache").
		WithWorkdir("/go/src/").
		WithExec([]string{"go", "mod", "download"})
}

// Enable the use of the Semantic Version Utility
func (m *Kv2) svu(ctx context.Context, source *dagger.Directory) *dagger.Container {
	return dag.Container().
		From("ghcr.io/caarlos0/svu:v3.2.2").
		WithDirectory("/tmp/.git", source.Directory(".git")).
		WithWorkdir("/tmp")
}

// Build the server binary
func (m *Kv2) buildServer(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
) *dagger.File {
	return m.devEnv(ctx, source).
		WithWorkdir("cmd/server").
		WithExec([]string{"go", "build", "-o", "/app/server", "."}).
		File("/app/server")
}

// Lint the project.
func (m *Kv2) Lint(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
) (string, error) {
	return m.devEnv(ctx, source).WithExec([]string{"go", "vet", "./..."}).Stdout(ctx)
}

// Run all tests.
func (m *Kv2) Test(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
) (string, error) {
	return m.devEnv(ctx, source).WithExec([]string{"go", "test", "-v", "./..."}).Stdout(ctx)
}

// Build the server container
func (m *Kv2) BuildServerContainer(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
) *dagger.Container {
	version, err := m.svu(ctx, source).WithExec([]string{"svu", "current"}).Stdout(ctx)
	if err != nil {
		dag.Error(err.Error())
	}

	server := m.buildServer(ctx, source)
	return dag.Container().
		From("gcr.io/distroless/base-debian12").
		WithLabel("org.opencontainers.image.title", "kv2").
		WithLabel("org.opencontainers.image.version", version).
		WithLabel("org.opencontainers.image.created", time.Now().String()).
		WithLabel("org.opencontainers.image.source", "https://github.com/hugginsio/kv2").
		WithLabel("org.opencontainers.image.licenses", "BSD-3-Clause").
		WithFile("/app/server", server).
		WithEntrypoint([]string{"/app/server"}).
		WithExposedPort(8080)
}
