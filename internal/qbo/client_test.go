package qbo_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	"servicemaster/internal/qbo"
	"servicemaster/internal/types"
)

func TestClient_Do_SendsBearerToken(t *testing.T) {
	t.Parallel()

	var gotAuth string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	t.Cleanup(server.Close)

	client, err := qbo.NewClient(qbo.Config{
		BaseURL:     server.URL,
		RealmID:     "realm-1",
		TokenSource: qbo.StaticTokenSource{Token: "access-abc"},
	})
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}

	_, _, err = client.Do(context.Background(), http.MethodGet, "/companyinfo/1", nil, nil)
	if err != nil {
		t.Fatalf("Do: %v", err)
	}

	if gotAuth != "Bearer access-abc" {
		t.Fatalf("Authorization = %q, want Bearer access-abc", gotAuth)
	}
}

func TestClient_Do_RetriesRateLimit(t *testing.T) {
	t.Parallel()

	var attempts atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if attempts.Add(1) == 1 {
			w.Header().Set("Retry-After", "0")
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	t.Cleanup(server.Close)

	client, err := qbo.NewClient(qbo.Config{
		BaseURL:     server.URL,
		RealmID:     "realm-1",
		TokenSource: qbo.StaticTokenSource{Token: "token"},
		RetryPolicy: qbo.RetryPolicy{MaxAttempts: 3, BaseDelay: 1, MaxDelay: 10},
	})
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}

	_, _, err = client.Do(context.Background(), http.MethodGet, "/ping", nil, nil)
	if err != nil {
		t.Fatalf("Do: %v", err)
	}

	if attempts.Load() != 2 {
		t.Fatalf("attempts = %d, want 2", attempts.Load())
	}
}

func TestClient_Do_DoesNotRetryUnauthorized(t *testing.T) {
	t.Parallel()

	var attempts atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts.Add(1)
		w.Header().Set("intuit_tid", "tid-401")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"Fault":{"type":"AUTHENTICATION","Error":[{"Message":"Token expired"}]}}`))
	}))
	t.Cleanup(server.Close)

	client, err := qbo.NewClient(qbo.Config{
		BaseURL:     server.URL,
		RealmID:     "realm-1",
		TokenSource: qbo.StaticTokenSource{Token: "stale"},
		RetryPolicy: qbo.RetryPolicy{MaxAttempts: 3, BaseDelay: 1, MaxDelay: 10},
	})
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}

	_, _, err = client.Do(context.Background(), http.MethodGet, "/ping", nil, nil)
	if err == nil {
		t.Fatal("expected error")
	}

	if attempts.Load() != 1 {
		t.Fatalf("attempts = %d, want 1 (no retry on 401)", attempts.Load())
	}

	apiErr, ok := err.(*qbo.APIError)
	if !ok {
		t.Fatalf("error type = %T, want *qbo.APIError", err)
	}

	if apiErr.IntuitTID != "tid-401" {
		t.Fatalf("IntuitTID = %q", apiErr.IntuitTID)
	}

	if apiErr.Err != qbo.ErrUnauthorized {
		t.Fatalf("wrapped = %v, want ErrUnauthorized", apiErr.Err)
	}
}

func TestClient_QueryPages(t *testing.T) {
	t.Parallel()

	const page1 = `{
		"QueryResponse": {
			"startPosition": 1,
			"maxResults": 2,
			"totalCount": 2,
			"Customer": [
				{"Id": "1", "DisplayName": "Alpha"},
				{"Id": "2", "DisplayName": "Beta"}
			]
		}
	}`

	const page2 = `{
		"QueryResponse": {
			"startPosition": 3,
			"maxResults": 2,
			"totalCount": 1,
			"Customer": [
				{"Id": "3", "DisplayName": "Gamma"}
			]
		}
	}`

	var calls atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		switch calls.Add(1) {
		case 1:
			if !strings.Contains(query, "STARTPOSITION 1") || !strings.Contains(query, "MAXRESULTS 2") {
				t.Fatalf("first query = %q", query)
			}
			_, _ = w.Write([]byte(page1))
		case 2:
			if !strings.Contains(query, "STARTPOSITION 3") {
				t.Fatalf("second query = %q", query)
			}
			_, _ = w.Write([]byte(page2))
		default:
			t.Fatalf("unexpected extra call: %q", query)
		}
	}))
	t.Cleanup(server.Close)

	client, err := qbo.NewClient(qbo.Config{
		BaseURL:     server.URL,
		RealmID:     "realm-1",
		TokenSource: qbo.StaticTokenSource{Token: "token"},
	})
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}

	iter, err := client.QueryPages(context.Background(), "SELECT * FROM Customer", 2)
	if err != nil {
		t.Fatalf("QueryPages: %v", err)
	}

	var names []string
	for {
		ok, err := iter.Next(context.Background())
		if err != nil {
			t.Fatalf("Next: %v", err)
		}
		if !ok {
			break
		}

		page := iter.Page()
		var customers []types.Customer
		if err := qbo.DecodeEntities(page, "Customer", &customers); err != nil {
			t.Fatalf("DecodeEntities: %v", err)
		}

		for _, customer := range customers {
			names = append(names, customer.DisplayName)
		}
	}

	if calls.Load() != 2 {
		t.Fatalf("calls = %d, want 2", calls.Load())
	}

	want := []string{"Alpha", "Beta", "Gamma"}
	if len(names) != len(want) {
		t.Fatalf("names = %v, want %v", names, want)
	}

	for i := range want {
		if names[i] != want[i] {
			t.Fatalf("names[%d] = %q, want %q", i, names[i], want[i])
		}
	}
}

func TestIsRetryable(t *testing.T) {
	t.Parallel()

	retryErr := &qbo.APIError{StatusCode: http.StatusServiceUnavailable, Retryable: true, Err: qbo.ErrServerFault}
	if !qbo.IsRetryable(retryErr) {
		t.Fatal("expected retryable")
	}

	noRetry := &qbo.APIError{StatusCode: http.StatusBadRequest, Err: qbo.ErrBadRequest}
	if qbo.IsRetryable(noRetry) {
		t.Fatal("did not expect retryable")
	}
}

func TestClient_Query_EmptyQueryResponse(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"QueryResponse":{}}`))
	}))
	t.Cleanup(server.Close)

	client, err := qbo.NewClient(qbo.Config{
		BaseURL:     server.URL,
		RealmID:     "realm",
		TokenSource: qbo.StaticTokenSource{Token: "t"},
	})
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}

	_, err = client.Query(context.Background(), "SELECT * FROM Customer")
	if err != qbo.ErrEmptyResponse {
		t.Fatalf("err = %v, want ErrEmptyResponse", err)
	}
}

func TestClient_Do_ParsesValidationFault(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"Fault":{"type":"ValidationFault","Error":[{"Message":"Invalid","code":"6000"}]}}`))
	}))
	t.Cleanup(server.Close)

	client, err := qbo.NewClient(qbo.Config{
		BaseURL:     server.URL,
		RealmID:     "realm",
		TokenSource: qbo.StaticTokenSource{Token: "t"},
	})
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}

	_, _, err = client.Do(context.Background(), http.MethodGet, "/x", nil, nil)
	apiErr, ok := err.(*qbo.APIError)
	if !ok {
		t.Fatalf("type = %T", err)
	}

	if apiErr.Fault == nil || apiErr.Fault.Error[0].Code != "6000" {
		t.Fatalf("fault = %#v", apiErr.Fault)
	}
}
