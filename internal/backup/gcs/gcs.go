// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package gcs

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

type Configuration struct {
	BucketName   string
	FilePath     string
	ObjectSuffix string
}

type GoogleCloudStorage struct {
	config Configuration
}

func Initialize(config Configuration) *GoogleCloudStorage {
	return &GoogleCloudStorage{config: config}
}

func (gcs *GoogleCloudStorage) Backup(name string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create GCS client: %v", err)
	}

	defer client.Close()

	var fileBytes []byte
	if bytes, err := os.ReadFile(gcs.config.FilePath); err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	} else {
		fileBytes = bytes
	}

	bucket := client.Bucket(gcs.config.BucketName)
	writer := bucket.Object(name).NewWriter(ctx)

	writer.ContentType = "application/x-sqlite3"
	if _, err := writer.Write(fileBytes); err != nil {
		return fmt.Errorf("failed to write file to GCS: %v", err)
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("failed to close GCS writer: %v", err)
	}

	return nil
}

func (gcs *GoogleCloudStorage) Restore() error {
	log.Println("Attempting database restore")

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create GCS client: %v", err)
	}

	defer client.Close()

	query := &storage.Query{Prefix: "kv2"}
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

	reader, err := client.Bucket(gcs.config.BucketName).Object(mostRecent.Name).NewReader(ctx)
	if err != nil {
		return fmt.Errorf("failed to create GCS reader: %v", err)
	}

	defer reader.Close()

	fileBytes, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("failed to read file from GCS: %v", err)
	}

	if err := os.WriteFile(gcs.config.FilePath, fileBytes, 0644); err != nil {
		return fmt.Errorf("failed to write file to disk: %v", err)
	}

	log.Println("Database restored successfully")

	return nil
}
