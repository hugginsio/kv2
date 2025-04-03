package client

// API client for the kv2 secrets server.
type Client struct {
	server string
}

func NewClient(server string) *Client {
	return &Client{
		server: server,
	}
}

// TODO: start over
