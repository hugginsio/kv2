// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

// Package kms provides an interface to retrieve secrets from external key management systems.
package kms

import (
	"log"
	"regexp"
	"strings"

	"git.huggins.io/kv2/internal/kms/gsm"
)

type KeyManagementSystem interface {
	Retrieve(id string) (string, error)
}

// This method determines if the provided value contains a KMS reference.
// If it does, it will attempt to fetch the key value and return it to the caller.
func KmsMiddleware(value string) string {
	r := regexp.MustCompile(`[^:]*`)
	match := r.FindString(value)

	var result string
	switch match {
	case "gsm":
		if key, err := gsm.Retrieve(strings.TrimPrefix(value, match+"://")); err != nil {
			log.Fatalf("Failed to retrieve key from GSM: %v", err)
		} else {
			result = key
		}
	default:
		result = value
	}

	return result
}
