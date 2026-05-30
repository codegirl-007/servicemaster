package client

// Client will hold HTTP behavior for QBO requests.
type Client struct {
	cfg Config
}

// New constructs a Client with the provided configuration.
func New(cfg Config) *Client {
	return &Client{cfg: cfg}
}
