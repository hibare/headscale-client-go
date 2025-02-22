package headscale

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestUserResource_List(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedUsers := []User{
			{
				ID:        "1",
				Name:      "test-user",
				CreatedAt: time.Date(2025, 2, 22, 10, 0, 0, 0, time.UTC),
			},
		}

		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if r.URL.Path != "/api/v1/user" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(UsersResponse{Users: expectedUsers})
		})

		users, err := client.Users().List(context.Background())
		require.NoError(t, err)
		require.Equal(t, expectedUsers, users.Users)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		})

		users, err := client.Users().List(context.Background())
		require.Error(t, err)
		require.Empty(t, users.Users)
	})
}

func TestUserResource_Get(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedUser := User{
			ID:        "1",
			Name:      "test-user",
			CreatedAt: time.Date(2025, 2, 22, 10, 0, 0, 0, time.UTC),
		}

		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if r.URL.Path != "/api/v1/user/1" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(UserResponse{User: expectedUser})
		})

		user, err := client.Users().Get(context.Background(), "1")
		require.NoError(t, err)
		require.Equal(t, expectedUser, user.User)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		})

		user, err := client.Users().Get(context.Background(), "1")
		require.Error(t, err)
		require.Empty(t, user.User)
	})
}

func TestUserResource_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedUser := User{
			ID:        "1",
			Name:      "new-user",
			CreatedAt: time.Date(2025, 2, 22, 10, 0, 0, 0, time.UTC),
		}

		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if r.URL.Path != "/api/v1/user" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			var req CreateUserRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if req.Name != "new-user" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(expectedUser)
		})

		user, err := client.Users().Create(context.Background(), "new-user")
		require.NoError(t, err)
		require.Equal(t, expectedUser, user)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		})

		user, err := client.Users().Create(context.Background(), "new-user")
		require.Error(t, err)
		require.Empty(t, user)
	})
}

func TestUserResource_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if r.URL.Path != "/api/v1/user/1" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.WriteHeader(http.StatusOK)
		})

		err := client.Users().Delete(context.Background(), "1")
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		})

		err := client.Users().Delete(context.Background(), "1")
		require.Error(t, err)
	})
}

func TestUserResource_Rename(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if r.URL.Path != "/api/v1/user/1/rename/new-name" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.WriteHeader(http.StatusOK)
		})

		err := client.Users().Rename(context.Background(), "1", "new-name")
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		})

		err := client.Users().Rename(context.Background(), "1", "new-name")
		require.Error(t, err)
	})
}
