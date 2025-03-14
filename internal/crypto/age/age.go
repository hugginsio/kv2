package age

import (
	"bytes"
	"io"

	"filippo.io/age"
)

type Configuration struct {
	PrivateKey string // AGE-SECRET-KEY
	PublicKey  string // age1
}

type AgeCrypto struct {
	identity  *age.X25519Identity
	recipient *age.X25519Recipient
}

func Initialize(config Configuration) (*AgeCrypto, error) {
	recipient, err := age.ParseX25519Recipient(config.PublicKey)
	if err != nil {
		return nil, err
	}

	identity, err := age.ParseX25519Identity(config.PrivateKey)
	if err != nil {
		return nil, err
	}

	return &AgeCrypto{
		identity:  identity,
		recipient: recipient,
	}, nil
}

// Encrypt some bytes.
func (ac *AgeCrypto) Encrypt(data []byte) ([]byte, error) {
	out := &bytes.Buffer{}
	w, err := age.Encrypt(out, ac.recipient)
	if err != nil {
		return nil, err
	}

	if _, err := w.Write(data); err != nil {
		return nil, err
	}

	if err := w.Close(); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

// Decrypt some bytes.
func (ac *AgeCrypto) Decrypt(data []byte) ([]byte, error) {
	br := bytes.NewReader(data)
	r, err := age.Decrypt(br, ac.identity)
	if err != nil {
		return nil, err
	}

	out := &bytes.Buffer{}
	if _, err := io.Copy(out, r); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}
