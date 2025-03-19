// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

// Package gsm provides integration with Google Cloud's Secret Manager.
package gsm

import (
	"context"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

func Retrieve(id string) (string, error) {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", err
	}

	defer client.Close()

	secret, err := client.GetSecret(ctx, &secretmanagerpb.GetSecretRequest{Name: id})
	if err != nil {
		return "", err
	}

	version, err := client.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{Name: fmt.Sprintf("%s/versions/latest", secret.Name)})
	if err != nil {
		return "", err
	}

	return string(version.Payload.Data), nil
}
