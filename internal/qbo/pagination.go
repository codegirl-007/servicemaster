package qbo

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

// defaultPageSize matches Intuit's commonly recommended query page size (100).
// Larger pages mean fewer requests but bigger payloads; smaller pages reduce
// memory per page at the cost of more 429 exposure.
const defaultPageSize = 100

// QueryPage is one page of a QuickBooks query response.
//
// QuickBooks returns entity arrays under dynamic keys (Customer, Invoice, …).
// We keep them as json.RawMessage in Entities so this package does not import
// internal/types. Import jobs decode into []types.Customer etc. at the edge.
//
// TotalCount is the number of records in *this page*, not the company-wide
// total. That naming trips people up constantly. End-of-data detection is:
// keep paging while totalCount == pageSize; stop on a short page.
// For a global count, run SELECT COUNT(*) FROM <Entity> separately.
type QueryPage struct {
	StartPosition int
	MaxResults    int
	TotalCount    int
	Entities      map[string]json.RawMessage
}

// PageIterator lazily fetches query pages.
//
// We use an iterator instead of returning [][]byte upfront so a single import
// job can process and stage one page at a time (lower memory, earlier partial
// progress). The job should persist startPosition on the batch if it fails
// mid-iteration so resume does not restart from 1.
type PageIterator struct {
	client        *Client
	baseQuery     string
	pageSize      int
	startPosition int
	done          bool
	current       QueryPage
}

// Next fetches the next page. Returns (false, nil) when iteration is complete.
func (it *PageIterator) Next(ctx context.Context) (bool, error) {
	if it.done {
		return false, nil
	}

	page, err := it.client.queryPage(ctx, it.baseQuery, it.startPosition, it.pageSize)
	if err != nil {
		return false, err
	}

	it.current = page

	// Short page means we reached the last page (or an empty result set encoded
	// as totalCount < pageSize). Advance startPosition only when more may exist.
	if page.TotalCount < it.pageSize {
		it.done = true
	} else {
		it.startPosition += page.TotalCount
	}

	return true, nil
}

// Page returns the most recently fetched page. Only valid after Next returns true.
func (it *PageIterator) Page() QueryPage {
	return it.current
}

// paginatedQuery appends SQL pagination clauses to the caller's base query.
//
// The caller should pass a query *without* STARTPOSITION/MAXRESULTS; this
// function owns pagination math. Trailing semicolons are stripped because
// Intuit's query parser is picky and imports often copy SQL from tools that
// add them.
func paginatedQuery(baseQuery string, startPosition, pageSize int) string {
	trimmed := strings.TrimSpace(baseQuery)
	trimmed = strings.TrimSuffix(trimmed, ";")
	return fmt.Sprintf("%s STARTPOSITION %d MAXRESULTS %d", trimmed, startPosition, pageSize)
}

// decodeQueryPage splits Intuit's QueryResponse envelope into metadata + entities.
//
// We decode into map[string]json.RawMessage instead of a struct with Customer,
// Invoice, etc. fields because the entity key varies per query and a struct
// would need dozens of optional fields or custom UnmarshalJSON anyway.
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
			// Everything else is assumed to be an entity array for this page.
			page.Entities[key] = raw
		}
	}

	// A QueryResponse with only metadata and no entity arrays is treated as empty.
	// That should be rare for SELECT * but protects jobs from silent no-ops.
	if len(page.Entities) == 0 {
		return QueryPage{}, ErrEmptyResponse
	}

	return page, nil
}
