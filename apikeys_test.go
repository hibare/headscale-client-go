package headscale

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAPIKeyResource_List(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedKeys := []APIKey{
			{
				ID:        "1",
				Prefix:    "test-prefix",
				CreatedAt: time.Date(2025, 2, 22, 10, 0, 0, 0, time.UTC),
			},
		}

		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if r.URL.Path != "/api/v1/apikey" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(APIKeysResponse{APIKeys: expectedKeys})
		})

		keys, err := client.APIKeys().List(context.Background())
		require.NoError(t, err)
		require.Equal(t, expectedKeys, keys.APIKeys)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		})

		keys, err := client.APIKeys().List(context.Background())
		require.Error(t, err)
		require.Empty(t, keys.APIKeys)
	})
}

func TestAPIKeyResource_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expiration := time.Date(2025, 2, 23, 10, 0, 0, 0, time.UTC)
		expectedKey := APIKey{
			ID:        "1",
			Prefix:    "test-prefix",
			CreatedAt: time.Date(2025, 2, 22, 10, 0, 0, 0, time.UTC),
		}

		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if r.URL.Path != "/api/v1/apikey" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			var req AddAPIKeyRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if req.Expiration != expiration {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(expectedKey)
		})

		key, err := client.APIKeys().Create(context.Background(), expiration)
		require.NoError(t, err)
		require.Equal(t, expectedKey, key)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		})

		key, err := client.APIKeys().Create(context.Background(), time.Now().Add(24*time.Hour))
		require.Error(t, err)
		require.Empty(t, key)
	})
}

func TestAPIKeyResource_Expire(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if r.URL.Path != "/api/v1/apikey/expire" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			var req ExpireAPIKeyRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if req.Prefix != "test-prefix" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusOK)
		})

		err := client.APIKeys().Expire(context.Background(), "test-prefix")
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		})

		err := client.APIKeys().Expire(context.Background(), "test-prefix")
		require.Error(t, err)
	})
}

func TestAPIKeyResource_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if r.URL.Path != "/api/v1/apikey/test-prefix" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.WriteHeader(http.StatusOK)
		})

		err := client.APIKeys().Delete(context.Background(), "test-prefix")
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		})

		err := client.APIKeys().Delete(context.Background(), "test-prefix")
		require.Error(t, err)
	})
}
