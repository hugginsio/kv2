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
	// TODO: multi arch
	return m.devEnv(ctx, source, nil).
		WithWorkdir("cmd/server").
		WithEnvVariable("CGO_ENABLED", "0").
		WithExec([]string{"go", "build", "-ldflags", "-s -w", "-gcflags=all=-l -C", "-buildvcs", "-o", "/app/server", "."}).
		File("/app/server")
}

// Build the server container
func (m *Kv2) BuildServerContainer(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
) *dagger.Container {
	server := m.buildServer(ctx, source)
	return dag.Container().
		From("gcr.io/distroless/static-debian12").
		WithAnnotation("org.opencontainers.image.title", "kv2").
		WithAnnotation("org.opencontainers.image.created", time.Now().String()).
		WithAnnotation("org.opencontainers.image.source", "https://github.com/hugginsio/kv2").
		WithAnnotation("org.opencontainers.image.licenses", "BSD-3-Clause").
		WithFile("/app/server", server).
		WithEntrypoint([]string{"/app/server"}).
		WithExposedPort(8080).
		WithExposedPort(80) // used for development mode ONLY
}
