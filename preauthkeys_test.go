package headscale

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPreAuthKeyResource_List(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedKeys := []PreAuthKey{
			{
				ID:        "1",
				Key:       "test-key",
				CreatedAt: time.Date(2025, 2, 22, 10, 0, 0, 0, time.UTC),
			},
		}

		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if r.URL.Path != "/api/v1/preauthkey" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(PreAuthKeysResponse{PreAuthKeys: expectedKeys})
		})

		keys, err := client.PreAuthKeys().List(context.Background())
		require.NoError(t, err)
		require.Equal(t, expectedKeys, keys.PreAuthKeys)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		})

		keys, err := client.PreAuthKeys().List(context.Background())
		require.Error(t, err)
		require.Empty(t, keys.PreAuthKeys)
	})
}

func TestPreAuthKeyResource_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		fixedTime := time.Date(2025, 2, 22, 10, 0, 0, 0, time.UTC)
		expectedKey := PreAuthKeyResponse{
			PreAuthKey: []PreAuthKey{{
				ID:         "1",
				Key:        "test-key",
				User:       "test-user",
				Reusable:   true,
				Ephemeral:  false,
				CreatedAt:  fixedTime,
				Expiration: fixedTime.Add(24 * time.Hour),
				AclTags:    []string{"tag1", "tag2"},
			}},
		}

		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if r.URL.Path != "/api/v1/preauthkey" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			var req CreatePreAuthKeyRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(expectedKey)
		})

		key, err := client.PreAuthKeys().Create(context.Background(), "test-user", true, false, fixedTime.Add(24*time.Hour), []string{"tag1", "tag2"})
		require.NoError(t, err)
		require.Equal(t, expectedKey.PreAuthKey, key.PreAuthKey)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		})

		key, err := client.PreAuthKeys().Create(context.Background(), "test-user", true, false, time.Now().Add(24*time.Hour), []string{"tag1", "tag2"})
		require.Error(t, err)
		require.Empty(t, key.PreAuthKey)
	})
}

func TestPreAuthKeyResource_Expire(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if r.URL.Path != "/api/v1/preauthkey/expire" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			var req ExpirePreAuthKeyRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusOK)
		})

		err := client.PreAuthKeys().Expire(context.Background(), "test-user", "test-key")
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		})

		err := client.PreAuthKeys().Expire(context.Background(), "test-user", "test-key")
		require.Error(t, err)
	})
}
