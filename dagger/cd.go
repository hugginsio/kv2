package main

import (
	"context"
	"dagger/kv-2/internal/dagger"
	"time"
)

// Build the server binary
func (m *Kv2) buildServer(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
) *dagger.File {
	return m.devEnv(ctx, source, nil).
		WithWorkdir("cmd/server").
		WithExec([]string{"go", "build", "-o", "/app/server", "."}).
		File("/app/server")
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
