package nodes

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/hibare/headscale-client-go/requests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNodeResource_List(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		filter := NodeListFilter{User: "testuser"}
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}
		fakeResp := NodesResponse{Nodes: []Node{{ID: "1", Name: "testnode"}}}

		mockReq.On("BuildURL", "node").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodGet, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*nodes.NodesResponse")).Run(func(args mock.Arguments) {
			resp := args.Get(2).(*NodesResponse) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
			*resp = fakeResp
		}).Return(nil)

		resp, err := n.List(ctx, filter)
		require.NoError(t, err)
		assert.Equal(t, fakeResp, resp)
		mockReq.AssertExpectations(t)
	})

	t.Run("build request error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		filter := NodeListFilter{}
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "node").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodGet, fakeURL, mock.Anything).Return(fakeReq, errors.New("build error"))

		resp, err := n.List(ctx, filter)
		require.Error(t, err)
		assert.Empty(t, resp.Nodes)
		mockReq.AssertExpectations(t)
	})

	t.Run("do error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		filter := NodeListFilter{}
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "node").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodGet, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*nodes.NodesResponse")).Return(errors.New("do error"))

		resp, err := n.List(ctx, filter)
		require.Error(t, err)
		assert.Empty(t, resp.Nodes)
		mockReq.AssertExpectations(t)
	})
}

func TestIsExitNode(t *testing.T) {
	tests := []struct {
		Name             string
		N                Node
		ExpectedExitNode bool
	}{
		{
			Name: "is exit node (IPv4)",
			N: Node{
				ID: "1",
				ApprovedRoutes: []string{
					"0.0.0.0/0",
				},
			},
			ExpectedExitNode: true,
		},
		{
			Name: "is Exit node (IPv6)",
			N: Node{
				ID: "1",
				ApprovedRoutes: []string{
					"::/0",
				},
			},
			ExpectedExitNode: true,
		},
		{
			Name: "is exit node",
			N: Node{
				ID: "1",
				ApprovedRoutes: []string{
					"::/0",
					"0.0.0.0/0",
				},
			},
			ExpectedExitNode: true,
		},
		{
			Name: "is not exit node",
			N: Node{
				ID:             "1",
				ApprovedRoutes: []string{},
			},
			ExpectedExitNode: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if tt.ExpectedExitNode {
				assert.True(t, tt.N.IsExitNode())
			} else {
				assert.False(t, tt.N.IsExitNode())
			}
		})
	}
}

func TestNodeResource_Get(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		id := "1"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}
		fakeResp := NodeResponse{Node: Node{ID: id, Name: "testnode"}}

		mockReq.On("BuildURL", "node", id).Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodGet, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*nodes.NodeResponse")).Run(func(args mock.Arguments) {
			resp := args.Get(2).(*NodeResponse) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
			*resp = fakeResp
		}).Return(nil)

		resp, err := n.Get(ctx, id)
		require.NoError(t, err)
		assert.Equal(t, fakeResp, resp)
		mockReq.AssertExpectations(t)
	})

	t.Run("build request error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		id := "1"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "node", id).Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodGet, fakeURL, mock.Anything).Return(fakeReq, errors.New("build error"))

		resp, err := n.Get(ctx, id)
		require.Error(t, err)
		assert.Empty(t, resp.Node)
		mockReq.AssertExpectations(t)
	})

	t.Run("do error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		id := "1"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "node", id).Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodGet, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*nodes.NodeResponse")).Return(errors.New("do error"))

		resp, err := n.Get(ctx, id)
		require.Error(t, err)
		assert.Empty(t, resp.Node)
		mockReq.AssertExpectations(t)
	})
}

func TestNodeResource_Register(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		user := "testuser"
		key := "testkey"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}
		fakeResp := NodeResponse{Node: Node{ID: "1", Name: "testnode"}}

		mockReq.On("BuildURL", "node", "register").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*nodes.NodeResponse")).Run(func(args mock.Arguments) {
			resp := args.Get(2).(*NodeResponse) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
			*resp = fakeResp
		}).Return(nil)

		resp, err := n.Register(ctx, user, key)
		require.NoError(t, err)
		assert.Equal(t, fakeResp, resp)
		mockReq.AssertExpectations(t)
	})

	t.Run("build request error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		user := "testuser"
		key := "testkey"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "node", "register").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, errors.New("build error"))

		resp, err := n.Register(ctx, user, key)
		require.Error(t, err)
		assert.Empty(t, resp.Node)
		mockReq.AssertExpectations(t)
	})

	t.Run("do error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		user := "testuser"
		key := "testkey"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "node", "register").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*nodes.NodeResponse")).Return(errors.New("do error"))

		resp, err := n.Register(ctx, user, key)
		require.Error(t, err)
		assert.Empty(t, resp.Node)
		mockReq.AssertExpectations(t)
	})
}

func TestNodeResource_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		id := "1"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "node", id).Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodDelete, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, nil).Return(nil)

		err := n.Delete(ctx, id)
		require.NoError(t, err)
		mockReq.AssertExpectations(t)
	})

	t.Run("build request error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		id := "1"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "node", id).Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodDelete, fakeURL, mock.Anything).Return(fakeReq, errors.New("build error"))

		err := n.Delete(ctx, id)
		require.Error(t, err)
		mockReq.AssertExpectations(t)
	})

	t.Run("do error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		id := "1"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "node", id).Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodDelete, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, nil).Return(errors.New("do error"))

		err := n.Delete(ctx, id)
		require.Error(t, err)
		mockReq.AssertExpectations(t)
	})
}

func TestNodeResource_Expire(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		id := "1"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "node", id, "expire").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, nil).Return(nil)

		err := n.Expire(ctx, id)
		require.NoError(t, err)
		mockReq.AssertExpectations(t)
	})

	t.Run("build request error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		id := "1"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "node", id, "expire").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, errors.New("build error"))

		err := n.Expire(ctx, id)
		require.Error(t, err)
		mockReq.AssertExpectations(t)
	})

	t.Run("do error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		id := "1"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "node", id, "expire").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, nil).Return(errors.New("do error"))

		err := n.Expire(ctx, id)
		require.Error(t, err)
		mockReq.AssertExpectations(t)
	})
}

func TestNodeResource_Rename(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		id := "1"
		name := "newname"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}
		fakeResp := NodeResponse{Node: Node{ID: id, Name: name}}

		mockReq.On("BuildURL", "node", id, "rename", name).Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*nodes.NodeResponse")).Run(func(args mock.Arguments) {
			resp := args.Get(2).(*NodeResponse) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
			*resp = fakeResp
		}).Return(nil)

		resp, err := n.Rename(ctx, id, name)
		require.NoError(t, err)
		assert.Equal(t, fakeResp, resp)
		mockReq.AssertExpectations(t)
	})

	t.Run("build request error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		id := "1"
		name := "newname"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "node", id, "rename", name).Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, errors.New("build error"))

		resp, err := n.Rename(ctx, id, name)
		require.Error(t, err)
		assert.Empty(t, resp.Node)
		mockReq.AssertExpectations(t)
	})

	t.Run("do error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		id := "1"
		name := "newname"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "node", id, "rename", name).Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*nodes.NodeResponse")).Return(errors.New("do error"))

		resp, err := n.Rename(ctx, id, name)
		require.Error(t, err)
		assert.Empty(t, resp.Node)
		mockReq.AssertExpectations(t)
	})
}

//nolint:dupl // This is identical to the TestNodeResource_ApproveRoutes test but with a different endpoint so it's not a duplicate.
func TestNodeResource_AddTags(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		id := "1"
		tags := []string{"tag1", "tag2"}
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}
		fakeResp := NodeResponse{Node: Node{ID: id, Tags: tags}}

		mockReq.On("BuildURL", "node", id, "tags").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*nodes.NodeResponse")).Run(func(args mock.Arguments) {
			resp := args.Get(2).(*NodeResponse) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
			*resp = fakeResp
		}).Return(nil)

		resp, err := n.AddTags(ctx, id, tags)
		require.NoError(t, err)
		assert.Equal(t, fakeResp, resp)
		mockReq.AssertExpectations(t)
	})

	t.Run("build request error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		id := "1"
		tags := []string{"tag1", "tag2"}
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "node", id, "tags").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, errors.New("build error"))

		resp, err := n.AddTags(ctx, id, tags)
		require.Error(t, err)
		assert.Empty(t, resp.Node)
		mockReq.AssertExpectations(t)
	})

	t.Run("do error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		id := "1"
		tags := []string{"tag1", "tag2"}
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "node", id, "tags").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*nodes.NodeResponse")).Return(errors.New("do error"))

		resp, err := n.AddTags(ctx, id, tags)
		require.Error(t, err)
		assert.Empty(t, resp.Node)
		mockReq.AssertExpectations(t)
	})
}

func TestNodeResource_BackFillIP(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		id := "1"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}
		fakeResp := BackfillIPsResponse{Changes: []string{"ip1", "ip2"}}

		mockReq.On("BuildURL", "node", id, "backfill_ip").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*nodes.BackfillIPsResponse")).Run(func(args mock.Arguments) {
			resp := args.Get(2).(*BackfillIPsResponse) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
			*resp = fakeResp
		}).Return(nil)

		resp, err := n.BackFillIP(ctx, id)
		require.NoError(t, err)
		assert.Equal(t, fakeResp, resp)
		mockReq.AssertExpectations(t)
	})

	t.Run("build request error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		id := "1"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "node", id, "backfill_ip").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, errors.New("build error"))

		resp, err := n.BackFillIP(ctx, id)
		require.Error(t, err)
		assert.Empty(t, resp.Changes)
		mockReq.AssertExpectations(t)
	})

	t.Run("do error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		id := "1"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "node", id, "backfill_ip").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*nodes.BackfillIPsResponse")).Return(errors.New("do error"))

		resp, err := n.BackFillIP(ctx, id)
		require.Error(t, err)
		assert.Empty(t, resp.Changes)
		mockReq.AssertExpectations(t)
	})
}

//nolint:dupl // This is identical to the TestNodeResource_AddTags test but with a different endpoint so it's not a duplicate.
func TestNodeResource_ApproveRoutes(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		id := "1"
		routes := []string{"10.0.0.0/24", "192.168.1.0/24"}
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}
		fakeResp := NodeResponse{Node: Node{ID: id, ApprovedRoutes: routes}}

		mockReq.On("BuildURL", "node", id, "approve_routes").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*nodes.NodeResponse")).Run(func(args mock.Arguments) {
			resp := args.Get(2).(*NodeResponse) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
			*resp = fakeResp
		}).Return(nil)

		resp, err := n.ApproveRoutes(ctx, id, routes)
		require.NoError(t, err)
		assert.Equal(t, fakeResp, resp)
		mockReq.AssertExpectations(t)
	})

	t.Run("build request error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		id := "1"
		routes := []string{"10.0.0.0/24", "192.168.1.0/24"}
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "node", id, "approve_routes").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, errors.New("build error"))

		resp, err := n.ApproveRoutes(ctx, id, routes)
		require.Error(t, err)
		assert.Empty(t, resp.Node)
		mockReq.AssertExpectations(t)
	})

	t.Run("do error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		n := &NodeResource{R: mockReq}
		ctx := t.Context()
		id := "1"
		routes := []string{"10.0.0.0/24", "192.168.1.0/24"}
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "node", id, "approve_routes").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*nodes.NodeResponse")).Return(errors.New("do error"))

		resp, err := n.ApproveRoutes(ctx, id, routes)
		require.Error(t, err)
		assert.Empty(t, resp.Node)
		mockReq.AssertExpectations(t)
	})
}
