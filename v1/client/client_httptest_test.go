package client

import (
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/hibare/headscale-client-go/v1/preauthkeys"
	"github.com/hibare/headscale-client-go/v1/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type requestRecord struct {
	Method string
	Path   string
	Query  string
	Body   string
	Header http.Header
}

func TestClient_HTTPRequests(t *testing.T) {
	var mu sync.Mutex
	var requests []requestRecord

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		_ = r.Body.Close()

		mu.Lock()
		requests = append(requests, requestRecord{
			Method: r.Method,
			Path:   r.URL.Path,
			Query:  r.URL.RawQuery,
			Body:   string(body),
			Header: r.Header.Clone(),
		})
		mu.Unlock()

		w.Header().Set("Content-Type", "application/json")

		switch r.Method + " " + r.URL.Path {
		case "GET /api/v1/apikey":
			_, _ = w.Write([]byte(`{"apiKeys":[]}`))
		case "GET /api/v1/user":
			_, _ = w.Write([]byte(`{"users":[]}`))
		case "GET /api/v1/node/node-1":
			_, _ = w.Write([]byte(`{"node":{"id":"node-1"}}`))
		case "PUT /api/v1/policy":
			_, _ = w.Write([]byte(`{"policy":"","updatedAt":""}`))
		case "POST /api/v1/preauthkey":
			_, _ = w.Write([]byte(`{"preAuthKey":{"id":"key-1","key":"abc","reusable":true}}`))
		default:
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"message":"not found"}`))
		}
	}))
	defer srv.Close()

	client, err := NewClient(srv.URL, "test-api-key", ClientOptions{})
	require.NoError(t, err)

	getRequests := func() []requestRecord {
		mu.Lock()
		defer mu.Unlock()
		out := make([]requestRecord, len(requests))
		copy(out, requests)
		return out
	}

	resetRequests := func() {
		mu.Lock()
		defer mu.Unlock()
		requests = nil
	}

	t.Run("APIKeys.List", func(t *testing.T) {
		ctx := t.Context()
		resetRequests()
		_, listErr := client.APIKeys().List(ctx)
		require.NoError(t, listErr)
		reqs := getRequests()
		require.Len(t, reqs, 1)
		assert.Equal(t, http.MethodGet, reqs[0].Method)
		assert.Equal(t, "/api/v1/apikey", reqs[0].Path)
		assert.Empty(t, reqs[0].Query)
		assert.Empty(t, reqs[0].Body)
		assert.Contains(t, reqs[0].Header.Get("Authorization"), "Bearer test-api-key")
	})

	t.Run("Users.List with filter", func(t *testing.T) {
		ctx := t.Context()
		resetRequests()
		_, listErr := client.Users().List(ctx, users.UserListFilter{Name: "test"})
		require.NoError(t, listErr)
		reqs := getRequests()
		require.Len(t, reqs, 1)
		assert.Equal(t, http.MethodGet, reqs[0].Method)
		assert.Equal(t, "/api/v1/user", reqs[0].Path)
		assert.Equal(t, "name=test", reqs[0].Query)
		assert.Empty(t, reqs[0].Body)
	})

	t.Run("Nodes.Get", func(t *testing.T) {
		ctx := t.Context()
		resetRequests()
		_, getErr := client.Nodes().Get(ctx, "node-1")
		require.NoError(t, getErr)
		reqs := getRequests()
		require.Len(t, reqs, 1)
		assert.Equal(t, http.MethodGet, reqs[0].Method)
		assert.Equal(t, "/api/v1/node/node-1", reqs[0].Path)
		assert.Empty(t, reqs[0].Query)
		assert.Empty(t, reqs[0].Body)
	})

	t.Run("Policy.Update", func(t *testing.T) {
		ctx := t.Context()
		resetRequests()
		_, updateErr := client.Policy().Update(ctx, `{"acls":[]}`)
		require.NoError(t, updateErr)
		reqs := getRequests()
		require.Len(t, reqs, 1)
		assert.Equal(t, http.MethodPut, reqs[0].Method)
		assert.Equal(t, "/api/v1/policy", reqs[0].Path)
		assert.Empty(t, reqs[0].Query)
		assert.Contains(t, reqs[0].Body, `acls`)
	})

	t.Run("PreAuthKeys.Create", func(t *testing.T) {
		ctx := t.Context()
		resetRequests()
		_, createErr := client.PreAuthKeys().Create(ctx, preauthkeys.CreatePreAuthKeyRequest{
			User:       "u1",
			Reusable:   true,
			Expiration: time.Now().Add(1 * time.Hour),
		})
		require.NoError(t, createErr)
		reqs := getRequests()
		require.Len(t, reqs, 1)
		assert.Equal(t, http.MethodPost, reqs[0].Method)
		assert.Equal(t, "/api/v1/preauthkey", reqs[0].Path)
		assert.Empty(t, reqs[0].Query)
		assert.Contains(t, reqs[0].Body, `"user":"u1"`)
		assert.Contains(t, reqs[0].Body, `"reusable":true`)
		assert.Contains(t, reqs[0].Body, `"expiration"`)
	})
}
