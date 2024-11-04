package headscale

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNodeResource_List(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		nodeResource := &NodeResource{Client: client}
		client.(*MockClient).On("buildURL", "node").Return(&url.URL{Path: "http://example.com/api/v1/node"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodGet, &url.URL{Path: "http://example.com/api/v1/node"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			resp := args.Get(1).(*NodesResponse)
			resp.Nodes = []Node{{ID: "1", Name: "test-node"}}
		})

		nodes, err := nodeResource.List(context.Background())
		assert.NoError(t, err)
		assert.Len(t, nodes.Nodes, 1)
		assert.Equal(t, "test-node", nodes.Nodes[0].Name)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		nodeResource := &NodeResource{Client: client}
		client.(*MockClient).On("buildURL", "node").Return(&url.URL{Path: "http://example.com/api/v1/node"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodGet, &url.URL{Path: "http://example.com/api/v1/node"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		nodes, err := nodeResource.List(context.Background())
		assert.Error(t, err)
		assert.Empty(t, nodes.Nodes)
	})
}

func TestNodeResource_Get(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		nodeResource := &NodeResource{Client: client}
		client.(*MockClient).On("buildURL", "node", "1").Return(&url.URL{Path: "http://example.com/api/v1/node/1"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodGet, &url.URL{Path: "http://example.com/api/v1/node/1"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			resp := args.Get(1).(*NodeResponse)
			resp.Node = Node{ID: "1", Name: "test-node"}
		})

		node, err := nodeResource.Get(context.Background(), "1")
		assert.NoError(t, err)
		assert.Equal(t, "test-node", node.Node.Name)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		nodeResource := &NodeResource{Client: client}
		client.(*MockClient).On("buildURL", "node", "1").Return(&url.URL{Path: "http://example.com/api/v1/node/1"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodGet, &url.URL{Path: "http://example.com/api/v1/node/1"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		node, err := nodeResource.Get(context.Background(), "1")
		assert.Error(t, err)
		assert.Empty(t, node.Node)
	})
}
func TestNodeResource_Register(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		nodeResource := &NodeResource{Client: client}
		client.(*MockClient).On("buildURL", "node").Return(&url.URL{Path: "http://example.com/api/v1/node"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPost, &url.URL{Path: "http://example.com/api/v1/node"}, requestOptions{
			queryParams: map[string]interface{}{
				"user": "test-user",
				"key":  "test-key",
			},
		}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			resp := args.Get(1).(*NodeResponse)
			resp.Node = Node{ID: "1", Name: "test-node"}
		})

		node, err := nodeResource.Register(context.Background(), "test-user", "test-key")
		assert.NoError(t, err)
		assert.Equal(t, "test-node", node.Node.Name)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		nodeResource := &NodeResource{Client: client}
		client.(*MockClient).On("buildURL", "node").Return(&url.URL{Path: "http://example.com/api/v1/node"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPost, &url.URL{Path: "http://example.com/api/v1/node"}, requestOptions{
			queryParams: map[string]interface{}{
				"user": "test-user",
				"key":  "test-key",
			},
		}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		node, err := nodeResource.Register(context.Background(), "test-user", "test-key")
		assert.Error(t, err)
		assert.Empty(t, node.Node)
	})
}
func TestNodeResource_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		nodeResource := &NodeResource{Client: client}
		client.(*MockClient).On("buildURL", "node", "1").Return(&url.URL{Path: "http://example.com/api/v1/node/1"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodDelete, &url.URL{Path: "http://example.com/api/v1/node/1"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil)

		err := nodeResource.Delete(context.Background(), "1")
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		nodeResource := &NodeResource{Client: client}
		client.(*MockClient).On("buildURL", "node", "1").Return(&url.URL{Path: "http://example.com/api/v1/node/1"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodDelete, &url.URL{Path: "http://example.com/api/v1/node/1"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		err := nodeResource.Delete(context.Background(), "1")
		assert.Error(t, err)
	})
}
func TestNodeResource_Expire(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		nodeResource := &NodeResource{Client: client}
		client.(*MockClient).On("buildURL", "node", "1", "expire").Return(&url.URL{Path: "http://example.com/api/v1/node/1/expire"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPost, &url.URL{Path: "http://example.com/api/v1/node/1/expire"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil)

		err := nodeResource.Expire(context.Background(), "1")
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		nodeResource := &NodeResource{Client: client}
		client.(*MockClient).On("buildURL", "node", "1", "expire").Return(&url.URL{Path: "http://example.com/api/v1/node/1/expire"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPost, &url.URL{Path: "http://example.com/api/v1/node/1/expire"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		err := nodeResource.Expire(context.Background(), "1")
		assert.Error(t, err)
	})
}
func TestNodeResource_Rename(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		nodeResource := &NodeResource{Client: client}
		client.(*MockClient).On("buildURL", "node", "1", "rename", "new-name").Return(&url.URL{Path: "http://example.com/api/v1/node/1/rename/new-name"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPost, &url.URL{Path: "http://example.com/api/v1/node/1/rename/new-name"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			resp := args.Get(1).(*NodeResponse)
			resp.Node = Node{ID: "1", Name: "new-name"}
		})

		node, err := nodeResource.Rename(context.Background(), "1", "new-name")
		assert.NoError(t, err)
		assert.Equal(t, "new-name", node.Node.Name)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		nodeResource := &NodeResource{Client: client}
		client.(*MockClient).On("buildURL", "node", "1", "rename", "new-name").Return(&url.URL{Path: "http://example.com/api/v1/node/1/rename/new-name"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPost, &url.URL{Path: "http://example.com/api/v1/node/1/rename/new-name"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		node, err := nodeResource.Rename(context.Background(), "1", "new-name")
		assert.Error(t, err)
		assert.Empty(t, node.Node)
	})
}
func TestNodeResource_GetRoutes(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		nodeResource := &NodeResource{Client: client}
		client.(*MockClient).On("buildURL", "node", "1", "routes").Return(&url.URL{Path: "http://example.com/api/v1/node/1/routes"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodGet, &url.URL{Path: "http://example.com/api/v1/node/1/routes"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			resp := args.Get(1).(*RoutesResponse)
			resp.Routes = []Route{{ID: "1", Prefix: "192.168.1.0/24"}}
		})

		routes, err := nodeResource.GetRoutes(context.Background(), "1")
		assert.NoError(t, err)
		assert.Len(t, routes.Routes, 1)
		assert.Equal(t, "192.168.1.0/24", routes.Routes[0].Prefix)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		nodeResource := &NodeResource{Client: client}
		client.(*MockClient).On("buildURL", "node", "1", "routes").Return(&url.URL{Path: "http://example.com/api/v1/node/1/routes"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodGet, &url.URL{Path: "http://example.com/api/v1/node/1/routes"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		routes, err := nodeResource.GetRoutes(context.Background(), "1")
		assert.Error(t, err)
		assert.Empty(t, routes.Routes)
	})
}
func TestNodeResource_AddTags(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		nodeResource := &NodeResource{Client: client}
		client.(*MockClient).On("buildURL", "node", "1", "tags").Return(&url.URL{Path: "http://example.com/api/v1/node/1/tags"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPost, &url.URL{Path: "http://example.com/api/v1/node/1/tags"}, requestOptions{
			body: AddTagsRequest{Tags: []string{"tag1", "tag2"}},
		}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			resp := args.Get(1).(*NodeResponse)
			resp.Node = Node{ID: "1", Name: "test-node"}
		})

		node, err := nodeResource.AddTags(context.Background(), "1", []string{"tag1", "tag2"})
		assert.NoError(t, err)
		assert.Equal(t, "test-node", node.Node.Name)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		nodeResource := &NodeResource{Client: client}
		client.(*MockClient).On("buildURL", "node", "1", "tags").Return(&url.URL{Path: "http://example.com/api/v1/node/1/tags"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPost, &url.URL{Path: "http://example.com/api/v1/node/1/tags"}, requestOptions{
			body: AddTagsRequest{Tags: []string{"tag1", "tag2"}},
		}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		node, err := nodeResource.AddTags(context.Background(), "1", []string{"tag1", "tag2"})
		assert.Error(t, err)
		assert.Empty(t, node.Node)
	})
}

func TestNodeResource_UpdateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		nodeResource := &NodeResource{Client: client}
		client.(*MockClient).On("buildURL", "node", "1", "user").Return(&url.URL{Path: "http://example.com/api/v1/node/1/user"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPost, &url.URL{Path: "http://example.com/api/v1/node/1/user"}, requestOptions{
			queryParams: map[string]interface{}{
				"user": "new-user",
			},
		}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			resp := args.Get(1).(*NodeResponse)
			resp.Node = Node{ID: "1", Name: "test-node"}
		})

		node, err := nodeResource.UpdateUser(context.Background(), "1", "new-user")
		assert.NoError(t, err)
		assert.Equal(t, "test-node", node.Node.Name)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		nodeResource := &NodeResource{Client: client}
		client.(*MockClient).On("buildURL", "node", "1", "user").Return(&url.URL{Path: "http://example.com/api/v1/node/1/user"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPost, &url.URL{Path: "http://example.com/api/v1/node/1/user"}, requestOptions{
			queryParams: map[string]interface{}{
				"user": "new-user",
			},
		}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		node, err := nodeResource.UpdateUser(context.Background(), "1", "new-user")
		assert.Error(t, err)
		assert.Empty(t, node.Node)
	})
}
