package headscale

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRoutesResource_List(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedRoutes := []Route{
			{
				ID:        "1",
				Prefix:    "192.168.1.0/24",
				CreatedAt: time.Date(2025, 2, 22, 10, 0, 0, 0, time.UTC),
			},
		}

		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if r.URL.Path != "/api/v1/routes" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(RoutesResponse{Routes: expectedRoutes})
		})

		routes, err := client.Routes().List(context.Background())
		require.NoError(t, err)
		require.Equal(t, expectedRoutes, routes.Routes)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		})

		routes, err := client.Routes().List(context.Background())
		require.Error(t, err)
		require.Empty(t, routes.Routes)
	})
}

func TestRoutesResource_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if r.URL.Path != "/api/v1/routes/1" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.WriteHeader(http.StatusOK)
		})

		err := client.Routes().Delete(context.Background(), "1")
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		})

		err := client.Routes().Delete(context.Background(), "1")
		require.Error(t, err)
	})
}

func TestRoutesResource_Enable(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if r.URL.Path != "/api/v1/routes/1/enable" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.WriteHeader(http.StatusOK)
		})

		err := client.Routes().Enable(context.Background(), "1")
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		})

		err := client.Routes().Enable(context.Background(), "1")
		require.Error(t, err)
	})
}

func TestRoutesResource_Disable(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if r.URL.Path != "/api/v1/routes/1/disable" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.WriteHeader(http.StatusOK)
		})

		err := client.Routes().Disable(context.Background(), "1")
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		})

		err := client.Routes().Disable(context.Background(), "1")
		require.Error(t, err)
	})
}
