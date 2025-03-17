package main

import (
	"context"
	"dagger/kv-2/internal/dagger"
	"fmt"
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

func (m *Kv2) Release(
	ctx context.Context,
	tag string,
	registry string,
	imageName string,
	username string,
	password *dagger.Secret,
) error {
	source := dag.Git("https://github.com/hugginsio/kv2.git", dagger.GitOpts{KeepGitDir: true}).Tag(tag).Tree()
	serverContainer := m.BuildServerContainer(ctx, source).
		WithLabel("org.opencontainers.image.version", tag).
		WithRegistryAuth(registry, username, password)

	if _, err := serverContainer.Publish(ctx, fmt.Sprintf("%s/%s:%s", registry, imageName, tag)); err != nil {
		return err
	}

	if _, err := serverContainer.Publish(ctx, fmt.Sprintf("%s/%s:latest", registry, imageName)); err != nil {
		return err
	}

	return nil
}
