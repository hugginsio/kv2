package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"git.huggins.io/kv2/api"
)

// API client for the kv2 secrets server.
type Client struct {
	server string
}

func NewClient(server string) *Client {
	return &Client{
		server: server,
	}
}

func parseResponse[RESP any](res *http.Response) (RESP, error) {
	var result RESP

	defer res.Body.Close()

	err := json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return result, fmt.Errorf("response parsing failed: %w", err)
	}

	return result, nil
}

// TODO: this is not good
func post[REQ any, RESP any](url string, body REQ, expectResponse bool) (RESP, error) {
	var result RESP

	request, err := json.Marshal(body)
	if err != nil {
		return result, fmt.Errorf("request parsing failed: %w", err)
	}

	res, err := http.Post(url, "application/json", bytes.NewReader(request))
	if err != nil {
		return result, fmt.Errorf("request failed: %w", err)
	}

	defer res.Body.Close()

	if expectResponse {
		return parseResponse[RESP](res)
	}

	return result, nil
}

func get[RESP any](url string) (RESP, error) {
	var result RESP

	res, err := http.Get(url)
	if err != nil {
		return result, fmt.Errorf("request failed: %w", err)
	}

	defer res.Body.Close()
	return parseResponse[RESP](res)
}

func (c *Client) Backup(request api.BackupRequest) error {
	url := fmt.Sprintf("%s/secrets/backup", c.server)

	if _, err := post[api.BackupRequest, struct{}](url, request, false); err != nil {
		return fmt.Errorf("failed to trigger backup: %w", err)
	}

	return nil
}

func (c *Client) Create(request api.CreateSecretRequest) error {
	url := fmt.Sprintf("%s/secrets/create", c.server)

	if _, err := post[api.CreateSecretRequest, struct{}](url, request, false); err != nil {
		return fmt.Errorf("failed to create secret: %w", err)
	}

	return nil
}

func (c *Client) Delete(request api.DeleteSecretRequest) error {
	url := fmt.Sprintf("%s/secrets/delete", c.server)

	if _, err := post[api.DeleteSecretRequest, struct{}](url, request, false); err != nil {
		return fmt.Errorf("failed to delete secret: %w", err)
	}

	return nil
}

func (c *Client) List() ([]api.ListSecretResponse, error) {
	var secrets []api.ListSecretResponse
	url := fmt.Sprintf("%s/secrets", c.server)

	secrets, err := get[[]api.ListSecretResponse](url)
	if err != nil {
		return nil, fmt.Errorf("failed to list secrets: %w", err)
	}

	return secrets, nil
}

func (c *Client) Read(request api.ReadSecretRequest) (api.Secret, error) {
	var secret api.Secret
	url := fmt.Sprintf("%s/secrets/read", c.server)

	secret, err := post[api.ReadSecretRequest, api.Secret](url, request, true)
	if err != nil {
		return secret, fmt.Errorf("failed to read secret: %w", err)
	}

	return secret, nil
}

func (c *Client) Revert(request api.RevertSecretRequest) error {
	url := fmt.Sprintf("%s/secrets/revert", c.server)

	if _, err := post[api.RevertSecretRequest, struct{}](url, request, false); err != nil {
		return fmt.Errorf("failed to revert secret: %w", err)
	}

	return nil
}

func (c *Client) Update(request api.UpdateSecretRequest) error {
	url := fmt.Sprintf("%s/secrets/update", c.server)

	if _, err := post[api.UpdateSecretRequest, struct{}](url, request, false); err != nil {
		return fmt.Errorf("failed to update secret: %w", err)
	}

	return nil
}
