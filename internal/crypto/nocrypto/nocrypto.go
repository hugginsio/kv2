// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

// Package nocrypto provides a no-op implementation of the crypto interface for development mode.
package nocrypto

type NoCrypto struct{}

func Initialize() *NoCrypto {
	return &NoCrypto{}
}

// Encrypt some bytes.
func (nc *NoCrypto) Encrypt(data []byte) ([]byte, error) {
	return data, nil
}

// Decrypt some bytes.
func (nc *NoCrypto) Decrypt(data []byte) ([]byte, error) {
	return data, nil
}
