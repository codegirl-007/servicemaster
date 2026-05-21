package qbo

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

const defaultPageSize = 100

// QueryPage is one page of a QuickBooks query response.
//
// TotalCount is the number of records in this page, not the total entities in
// the company. Use SELECT COUNT(*) for a global count when needed.
type QueryPage struct {
	StartPosition int
	MaxResults    int
	TotalCount    int
	Entities      map[string]json.RawMessage
}

// PageIterator walks query pages until a short page is returned.
type PageIterator struct {
	client        *Client
	baseQuery     string
	pageSize      int
	startPosition int
	done          bool
	current       QueryPage
}

// Next fetches the next page. It returns false when iteration is complete.
func (it *PageIterator) Next(ctx context.Context) (bool, error) {
	if it.done {
		return false, nil
	}

	page, err := it.client.queryPage(ctx, it.baseQuery, it.startPosition, it.pageSize)
	if err != nil {
		return false, err
	}

	it.current = page

	if page.TotalCount < it.pageSize {
		it.done = true
	} else {
		it.startPosition += page.TotalCount
	}

	return true, nil
}

// Page returns the most recently fetched page. Call only after Next returns true.
func (it *PageIterator) Page() QueryPage {
	return it.current
}

func paginatedQuery(baseQuery string, startPosition, pageSize int) string {
	trimmed := strings.TrimSpace(baseQuery)
	trimmed = strings.TrimSuffix(trimmed, ";")
	return fmt.Sprintf("%s STARTPOSITION %d MAXRESULTS %d", trimmed, startPosition, pageSize)
}

func decodeQueryPage(body []byte) (QueryPage, error) {
	var outer struct {
		QueryResponse json.RawMessage `json:"QueryResponse"`
	}
	if err := json.Unmarshal(body, &outer); err != nil {
		return QueryPage{}, fmt.Errorf("decode query envelope: %w", err)
	}

	if len(outer.QueryResponse) == 0 {
		return QueryPage{}, ErrEmptyResponse
	}

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(outer.QueryResponse, &fields); err != nil {
		return QueryPage{}, fmt.Errorf("decode query response: %w", err)
	}

	page := QueryPage{Entities: make(map[string]json.RawMessage)}
	for key, raw := range fields {
		switch key {
		case "startPosition":
			_ = json.Unmarshal(raw, &page.StartPosition)
		case "maxResults":
			_ = json.Unmarshal(raw, &page.MaxResults)
		case "totalCount":
			_ = json.Unmarshal(raw, &page.TotalCount)
		default:
			page.Entities[key] = raw
		}
	}

	if len(page.Entities) == 0 {
		return QueryPage{}, ErrEmptyResponse
	}

	return page, nil
}
