// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package gcs

import (
	"context"
	"io"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/rs/zerolog/log"
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
		return err
	}

	defer client.Close()

	var fileBytes []byte
	if bytes, err := os.ReadFile(gcs.config.FilePath); err != nil {
		return err
	} else {
		fileBytes = bytes
	}

	bucket := client.Bucket(gcs.config.BucketName)
	writer := bucket.Object(name).NewWriter(ctx)

	writer.ContentType = "application/x-sqlite3"
	if _, err := writer.Write(fileBytes); err != nil {
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}

	return nil
}

func (gcs *GoogleCloudStorage) Restore() error {
	log.Debug().Msg("database restore started")

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
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
			return err
		}

		if mostRecent == nil || attrs.Updated.After(mostRecentTime) {
			mostRecent = attrs
			mostRecentTime = attrs.Updated
		}
	}

	if mostRecent == nil {
		return err
	}

	reader, err := client.Bucket(gcs.config.BucketName).Object(mostRecent.Name).NewReader(ctx)
	if err != nil {
		return err
	}

	defer reader.Close()

	fileBytes, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	if err := os.WriteFile(gcs.config.FilePath, fileBytes, 0644); err != nil {
		return err
	}

	log.Debug().Msg("database restore completed")
	return nil
}
