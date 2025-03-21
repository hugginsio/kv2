// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package gcs

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

type Configuration struct {
	BucketName      string
	DestinationPath string
}

type GoogleCloudStorage struct {
	config Configuration
}

func Initialize(config Configuration) *GoogleCloudStorage {
	return &GoogleCloudStorage{config: config}
}

func (gcs *GoogleCloudStorage) Restore() error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create GCS client: %v", err)
	}

	defer client.Close()

	query := &storage.Query{Prefix: "kv2_"}
	query.SetAttrSelection([]string{"Name", "Updated"})

	var mostRecent *storage.ObjectAttrs
	var mostRecentTime time.Time

	it := client.Bucket(gcs.config.BucketName).Objects(ctx, query)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return fmt.Errorf("error iterating objects: %v", err)
		}

		if mostRecent == nil || attrs.Updated.After(mostRecentTime) {
			mostRecent = attrs
			mostRecentTime = attrs.Updated
		}
	}

	if mostRecent == nil {
		return fmt.Errorf("no kv2.db files found in bucket")
	}

	// TODO: if file found, retrieve its contents and persist to the disk

	return nil
}
