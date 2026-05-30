package client

// Page represents a single paginated response payload.
type Page struct {
	StartPosition int
	MaxResults    int
	TotalCount    int
	Entities      map[string]any
}
