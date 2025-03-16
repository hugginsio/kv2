package main

import (
	"context"
	"dagger/kv-2/internal/dagger"
)

type Kv2 struct{}

// Build a ready-to-use development environment.
func (m *Kv2) devEnv(
	ctx context.Context,
	source *dagger.Directory,
	// +optional
	platform *dagger.Platform,
) *dagger.Container {
	if platform == nil {
		enginePlatform, err := dag.DefaultPlatform(ctx)
		if err != nil {
			panic(err)
		}

		platform = &enginePlatform
	}

	return dag.Container(dagger.ContainerOpts{Platform: *platform}).
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
