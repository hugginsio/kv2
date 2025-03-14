// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package sqlite

import (
	"log"

	"git.huggins.io/kv2/api"
	"git.huggins.io/kv2/internal/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Configuration struct {
	Dsn string
}

type SqliteDatabase struct {
	sql *gorm.DB
}

// Initialize a sqlite database for secret storage.
func Initialize(configuration Configuration) (*SqliteDatabase, error) {
	log.Println("Initializing database (sqlite)")
	db, err := gorm.Open(sqlite.Open(configuration.Dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&database.SecretRecord{}, &database.ValueRecord{}); err != nil {
		return nil, err
	}

	return &SqliteDatabase{sql: db}, nil
}

// List all secrets and their versions.
func (db *SqliteDatabase) List() ([]api.ListSecretResponse, error) {
	results := []api.ListSecretResponse{}
	var secrets []database.SecretRecord

	if err := db.sql.Find(&secrets).Error; err != nil {
		return nil, err
	}

	for _, secret := range secrets {
		var valueRecords []database.ValueRecord
		var secretVersions []api.SecretVersion

		if err := db.sql.Where("secret_record_id = ?", secret.ID).Order("version DESC").Find(&valueRecords).Error; err != nil {
			return nil, err
		}

		for _, secretValue := range valueRecords {
			secretVersions = append(secretVersions, secretValue.Version)
		}

		results = append(results, api.ListSecretResponse{
			Key:      secret.Key,
			Versions: secretVersions,
		})
	}

	return results, nil
}

// Create a new secret.
func (db *SqliteDatabase) Create(request api.CreateSecretRequest) error {
	existingSecret := db.sql.Find(&database.SecretRecord{}, "key = ?", request.Key)
	if existingSecret.Error != nil {
		return existingSecret.Error
	}

	if existingSecret.RowsAffected > 0 {
		return api.ErrorSecretAlreadyExists
	}

	return db.sql.Transaction(func(tx *gorm.DB) error {
		secretRecord := database.SecretRecord{
			Key: request.Key,
		}

		if err := tx.Create(&secretRecord).Error; err != nil {
			return err
		}

		valueRecord := database.ValueRecord{
			SecretRecordID: secretRecord.ID,
			Value:          request.Value,
			Version:        1,
		}

		if err := tx.Create(&valueRecord).Error; err != nil {
			return err
		}

		return nil
	})
}

// Retrieve the latest version of a secret.
func (db *SqliteDatabase) Read(request api.ReadSecretRequest) (api.Secret, error) {
	var requestedSecret database.SecretRecord
	var latestValue database.ValueRecord

	if err := db.sql.First(&requestedSecret, "key = ?", request.Key).Error; err != nil {
		return api.Secret{}, err
	}

	if err := db.sql.Where("secret_record_id = ?", requestedSecret.ID).Order("version DESC").Limit(1).Find(&latestValue).Error; err != nil {
		return api.Secret{}, err
	}

	return api.Secret{
		Key:     requestedSecret.Key,
		Value:   latestValue.Value,
		Version: latestValue.Version,
	}, nil
}

// Create a new version of an existing secret.
func (db *SqliteDatabase) Update(request api.UpdateSecretRequest) error {
	var secretRecord database.SecretRecord

	if err := db.sql.First(&secretRecord, "key = ?", request.Key).Error; err != nil {
		return err
	}

	return db.sql.Transaction(func(tx *gorm.DB) error {
		var latestValue database.ValueRecord
		if err := tx.Where("secret_record_id = ?", secretRecord.ID).Order("version DESC").Limit(1).Find(&latestValue).Error; err != nil {
			return err
		}

		newValue := database.ValueRecord{
			SecretRecordID: latestValue.SecretRecordID,
			Value:          request.Value,
			Version:        latestValue.Version + 1,
		}

		if err := tx.Create(&newValue).Error; err != nil {
			return err
		}

		var allValues []database.ValueRecord
		if err := tx.Where("secret_record_id = ?", secretRecord.ID).Order("version ASC").Find(&allValues).Error; err != nil {
			return err
		}

		if len(allValues) < 9 {
			return nil
		}

		if err := tx.Delete(allValues[0]).Error; err != nil {
			return err
		}

		return nil
	})
}

// Delete a secret and all its versions.
func (db *SqliteDatabase) Delete(request api.DeleteSecretRequest) error {
	var secretRecord database.SecretRecord
	if err := db.sql.Where("key = ?", request.Key).First(&secretRecord).Error; err != nil {
		return err
	}

	return db.sql.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&secretRecord).Error; err != nil {
			return err
		}

		var allValues []database.ValueRecord
		if err := tx.Find(&allValues, "secret_record_id = ?", secretRecord.ID).Error; err != nil {
			return err
		}

		if err := tx.Delete(&allValues).Error; err != nil {
			return err
		}

		return nil
	})
}

// Revert a secret to a previous version.
func (db *SqliteDatabase) Revert(request api.RevertSecretRequest) error {
	var secretRecord database.SecretRecord
	if err := db.sql.Where("key = ?", request.Key).First(&secretRecord).Error; err != nil {
		return err
	}

	return db.sql.Transaction(func(tx *gorm.DB) error {
		var allValues []database.ValueRecord
		if err := tx.Where("secret_record_id = ?", secretRecord.ID).Order("version DESC").Limit(2).Find(&allValues).Error; err != nil {
			return err
		}

		if len(allValues) == 1 {
			return api.ErrorCannotRevert
		}

		if err := tx.Delete(allValues[0]).Error; err != nil {
			return err
		}

		return nil
	})
}
