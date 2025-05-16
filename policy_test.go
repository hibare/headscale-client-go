package headscale

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPolicyResource_Get(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedPolicy := Policy{
			Policy:    "test-policy",
			UpdatedAt: "2025-02-22T10:00:00Z",
		}

		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if r.URL.Path != "/api/v1/policy" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(expectedPolicy)
		})

		policy, err := client.Policy().Get(context.Background())
		require.NoError(t, err)
		require.Equal(t, expectedPolicy, policy)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		})

		policy, err := client.Policy().Get(context.Background())
		require.Error(t, err)
		require.Empty(t, policy.Policy)
	})
}

func TestPolicyResource_Put(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if r.URL.Path != "/api/v1/policy" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			var req UpdatePolicyRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if req.Policy != "new-policy" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusOK)
		})

		err := client.Policy().Put(context.Background(), "new-policy")
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		})

		err := client.Policy().Put(context.Background(), "new-policy")
		require.Error(t, err)
	})
}
