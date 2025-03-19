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

	version, err := client.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{Name: fmt.Sprintf("%s/versions/latest", id)})
	if err != nil {
		return "", err
	}

	return string(version.Payload.Data), nil
}
