package qbo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	productionAPIHost = "https://quickbooks.api.intuit.com"
	sandboxAPIHost    = "https://sandbox-quickbooks.api.intuit.com"
)

// Config configures a QuickBooks API client for one realm (company).
type Config struct {
	BaseURL      string
	RealmID      string
	MinorVersion int
	HTTPClient   *http.Client
	TokenSource  TokenSource
	RetryPolicy  RetryPolicy
}

// Client performs authenticated QuickBooks Online API v3 requests.
type Client struct {
	baseURL      string
	realmID      string
	minorVersion int
	httpClient   *http.Client
	tokenSource  TokenSource
	retryPolicy  RetryPolicy
}

// NewClient validates config and returns a Client.
func NewClient(cfg Config) (*Client, error) {
	if cfg.TokenSource == nil {
		return nil, fmt.Errorf("qbo: TokenSource is required")
	}
	if strings.TrimSpace(cfg.RealmID) == "" {
		return nil, fmt.Errorf("qbo: RealmID is required")
	}

	baseURL := strings.TrimSpace(cfg.BaseURL)
	if baseURL == "" {
		return nil, fmt.Errorf("qbo: BaseURL is required")
	}

	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		baseURL:      strings.TrimRight(baseURL, "/"),
		realmID:      cfg.RealmID,
		minorVersion: cfg.MinorVersion,
		httpClient:   httpClient,
		tokenSource:  cfg.TokenSource,
		retryPolicy:  cfg.RetryPolicy.normalized(),
	}, nil
}

// ProductionBaseURL returns the company API base URL for production.
func ProductionBaseURL(realmID string) string {
	return fmt.Sprintf("%s/v3/company/%s", productionAPIHost, realmID)
}

// SandboxBaseURL returns the company API base URL for sandbox.
func SandboxBaseURL(realmID string) string {
	return fmt.Sprintf("%s/v3/company/%s", sandboxAPIHost, realmID)
}

// Do executes an authenticated HTTP request with retry boundaries.
//
// The caller owns closing resp.Body on success. On error, the body is drained
// and closed internally.
func (c *Client) Do(ctx context.Context, method, path string, query url.Values, body io.Reader) (*http.Response, []byte, error) {
	token, err := c.tokenSource.AccessToken(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("qbo: access token: %w", err)
	}

	endpoint, err := c.buildURL(path, query)
	if err != nil {
		return nil, nil, err
	}

	var lastErr error

	for attempt := 1; attempt <= c.retryPolicy.MaxAttempts; attempt++ {
		req, err := http.NewRequestWithContext(ctx, method, endpoint, body)
		if err != nil {
			return nil, nil, fmt.Errorf("qbo: build request: %w", err)
		}

		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Accept", "application/json")
		if body != nil {
			req.Header.Set("Content-Type", "application/json")
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, nil, fmt.Errorf("qbo: http request: %w", err)
		}

		respBody, readErr := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if readErr != nil {
			return nil, nil, fmt.Errorf("qbo: read response: %w", readErr)
		}

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return resp, respBody, nil
		}

		apiErr := newAPIError(resp.StatusCode, resp.Header.Get("intuit_tid"), respBody)
		lastErr = apiErr

		if !c.retryPolicy.shouldRetry(resp.StatusCode, attempt) {
			return resp, respBody, apiErr
		}

		if err := c.retryPolicy.wait(ctx, attempt, resp.Header.Get("Retry-After")); err != nil {
			return resp, respBody, err
		}
	}

	return nil, nil, lastErr
}

// Query runs a QuickBooks SQL query and returns the first page.
func (c *Client) Query(ctx context.Context, sql string) (QueryPage, error) {
	return c.queryPage(ctx, sql, 1, defaultPageSize)
}

// QueryPages returns an iterator over all pages for a base SQL query.
func (c *Client) QueryPages(ctx context.Context, sql string, pageSize int) (*PageIterator, error) {
	if strings.TrimSpace(sql) == "" {
		return nil, fmt.Errorf("qbo: query is required")
	}

	if pageSize <= 0 {
		pageSize = defaultPageSize
	}

	return &PageIterator{
		client:        c,
		baseQuery:     sql,
		pageSize:      pageSize,
		startPosition: 1,
	}, nil
}

func (c *Client) queryPage(ctx context.Context, sql string, startPosition, pageSize int) (QueryPage, error) {
	querySQL := paginatedQuery(sql, startPosition, pageSize)

	values := url.Values{}
	values.Set("query", querySQL)
	if c.minorVersion > 0 {
		values.Set("minorversion", fmt.Sprintf("%d", c.minorVersion))
	}

	_, body, err := c.Do(ctx, http.MethodGet, "/query", values, nil)
	if err != nil {
		return QueryPage{}, err
	}

	return decodeQueryPage(body)
}

func (c *Client) buildURL(path string, query url.Values) (string, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	endpoint := c.baseURL + path
	if len(query) == 0 {
		return endpoint, nil
	}

	parsed, err := url.Parse(endpoint)
	if err != nil {
		return "", fmt.Errorf("qbo: parse endpoint: %w", err)
	}

	parsed.RawQuery = query.Encode()
	return parsed.String(), nil
}

// DecodeEntities unmarshals one entity key from a query page into dest.
func DecodeEntities(page QueryPage, entityKey string, dest any) error {
	raw, ok := page.Entities[entityKey]
	if !ok {
		return fmt.Errorf("qbo: entity %q not present in page", entityKey)
	}

	if err := json.Unmarshal(raw, dest); err != nil {
		return fmt.Errorf("qbo: decode %s: %w", entityKey, err)
	}

	return nil
}
