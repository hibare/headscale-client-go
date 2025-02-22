package headscale

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNodeResource_List(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedNodes := []Node{
			{
				ID:        "1",
				Name:      "test-node",
				CreatedAt: time.Date(2025, 2, 22, 10, 0, 0, 0, time.UTC),
			},
		}

		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if r.URL.Path != "/api/v1/node" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(NodesResponse{Nodes: expectedNodes})
		})

		nodes, err := client.Nodes().List(context.Background())
		require.NoError(t, err)
		require.Equal(t, expectedNodes, nodes.Nodes)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		})

		nodes, err := client.Nodes().List(context.Background())
		require.Error(t, err)
		require.Empty(t, nodes.Nodes)
	})
}

func TestNodeResource_Get(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedNode := Node{
			ID:        "1",
			Name:      "test-node",
			CreatedAt: time.Date(2025, 2, 22, 10, 0, 0, 0, time.UTC),
		}

		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if r.URL.Path != "/api/v1/node/1" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(NodeResponse{Node: expectedNode})
		})

		node, err := client.Nodes().Get(context.Background(), "1")
		require.NoError(t, err)
		require.Equal(t, expectedNode, node.Node)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		})

		node, err := client.Nodes().Get(context.Background(), "1")
		require.Error(t, err)
		require.Empty(t, node.Node)
	})
}

func TestNodeResource_Register(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedNode := Node{
			ID:        "1",
			Name:      "test-node",
			CreatedAt: time.Date(2025, 2, 22, 10, 0, 0, 0, time.UTC),
		}

		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if r.URL.Path != "/api/v1/node" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if r.URL.Query().Get("user") != "test-user" || r.URL.Query().Get("key") != "test-key" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(NodeResponse{Node: expectedNode})
		})

		node, err := client.Nodes().Register(context.Background(), "test-user", "test-key")
		require.NoError(t, err)
		require.Equal(t, expectedNode, node.Node)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		})

		node, err := client.Nodes().Register(context.Background(), "test-user", "test-key")
		require.Error(t, err)
		require.Empty(t, node.Node)
	})
}

func TestNodeResource_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if r.URL.Path != "/api/v1/node/1" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.WriteHeader(http.StatusOK)
		})

		err := client.Nodes().Delete(context.Background(), "1")
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		})

		err := client.Nodes().Delete(context.Background(), "1")
		require.Error(t, err)
	})
}

func TestNodeResource_Expire(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if r.URL.Path != "/api/v1/node/1/expire" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.WriteHeader(http.StatusOK)
		})

		err := client.Nodes().Expire(context.Background(), "1")
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		})

		err := client.Nodes().Expire(context.Background(), "1")
		require.Error(t, err)
	})
}

func TestNodeResource_Rename(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedNode := Node{
			ID:        "1",
			Name:      "new-name",
			CreatedAt: time.Date(2025, 2, 22, 10, 0, 0, 0, time.UTC),
		}

		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if r.URL.Path != "/api/v1/node/1/rename/new-name" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(NodeResponse{Node: expectedNode})
		})

		node, err := client.Nodes().Rename(context.Background(), "1", "new-name")
		require.NoError(t, err)
		require.Equal(t, expectedNode, node.Node)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		})

		node, err := client.Nodes().Rename(context.Background(), "1", "new-name")
		require.Error(t, err)
		require.Empty(t, node.Node)
	})
}

func TestNodeResource_GetRoutes(t *testing.T) {
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
			if r.URL.Path != "/api/v1/node/1/routes" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(RoutesResponse{Routes: expectedRoutes})
		})

		routes, err := client.Nodes().GetRoutes(context.Background(), "1")
		require.NoError(t, err)
		require.Equal(t, expectedRoutes, routes.Routes)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		})

		routes, err := client.Nodes().GetRoutes(context.Background(), "1")
		require.Error(t, err)
		require.Empty(t, routes.Routes)
	})
}

func TestNodeResource_AddTags(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedNode := Node{
			ID:        "1",
			Name:      "test-node",
			CreatedAt: time.Date(2025, 2, 22, 10, 0, 0, 0, time.UTC),
		}

		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if r.URL.Path != "/api/v1/node/1/tags" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			var req AddTagsRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(NodeResponse{Node: expectedNode})
		})

		node, err := client.Nodes().AddTags(context.Background(), "1", []string{"tag1", "tag2"})
		require.NoError(t, err)
		require.Equal(t, expectedNode, node.Node)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		})

		node, err := client.Nodes().AddTags(context.Background(), "1", []string{"tag1", "tag2"})
		require.Error(t, err)
		require.Empty(t, node.Node)
	})
}

func TestNodeResource_UpdateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedNode := Node{
			ID:        "1",
			Name:      "test-node",
			CreatedAt: time.Date(2025, 2, 22, 10, 0, 0, 0, time.UTC),
		}

		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if r.URL.Path != "/api/v1/node/1/user" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if r.Header.Get("Authorization") != ExpectedTestBearerToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if r.URL.Query().Get("user") != "new-user" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(NodeResponse{Node: expectedNode})
		})

		node, err := client.Nodes().UpdateUser(context.Background(), "1", "new-user")
		require.NoError(t, err)
		require.Equal(t, expectedNode, node.Node)
	})

	t.Run("error", func(t *testing.T) {
		client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		})

		node, err := client.Nodes().UpdateUser(context.Background(), "1", "new-user")
		require.Error(t, err)
		require.Empty(t, node.Node)
	})
}
