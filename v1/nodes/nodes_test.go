package nodes

import (
	"context"
	"net/http"
	"testing"

	"github.com/hibare/headscale-client-go/requests"
	"github.com/hibare/headscale-client-go/v1/testutil"
	"github.com/stretchr/testify/assert"
)

func TestNodeResource_List(t *testing.T) {
	fixture := testutil.TestFixture[NodesResponse]{
		Endpoint:    "node",
		Method:      http.MethodGet,
		SuccessResp: NodesResponse{Nodes: []Node{{ID: "1", Name: "testnode"}}},
	}

	testutil.RunResourceTest(t, fixture, func(ctx context.Context, mockReq *requests.MockRequest) (NodesResponse, error) {
		n := &NodeResource{r: mockReq}
		return n.List(ctx, NodeListFilter{User: "testuser"})
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
	id := "1"
	fixture := testutil.TestFixture[NodeResponse]{
		Endpoint:    []any{"node", id},
		Method:      http.MethodGet,
		SuccessResp: NodeResponse{Node: Node{ID: id, Name: "testnode"}},
	}

	testutil.RunResourceTest(t, fixture, func(ctx context.Context, mockReq *requests.MockRequest) (NodeResponse, error) {
		n := &NodeResource{r: mockReq}
		return n.Get(ctx, id)
	})
}

func TestNodeResource_Register(t *testing.T) {
	fixture := testutil.TestFixture[NodeResponse]{
		Endpoint:    []any{"node", "register"},
		Method:      http.MethodPost,
		SuccessResp: NodeResponse{Node: Node{ID: "1", Name: "testnode"}},
	}

	testutil.RunResourceTest(t, fixture, func(ctx context.Context, mockReq *requests.MockRequest) (NodeResponse, error) {
		n := &NodeResource{r: mockReq}
		return n.Register(ctx, "testuser", "testkey")
	})
}

func TestNodeResource_Delete(t *testing.T) {
	id := "1"
	fixture := testutil.TestFixture[struct{}]{
		Endpoint:    []any{"node", id},
		Method:      http.MethodDelete,
		SuccessResp: struct{}{},
	}

	testutil.RunResourceTest(t, fixture, func(ctx context.Context, mockReq *requests.MockRequest) (struct{}, error) {
		n := &NodeResource{r: mockReq}
		err := n.Delete(ctx, id)
		return struct{}{}, err
	})
}

func TestNodeResource_Expire(t *testing.T) {
	id := "1"
	fixture := testutil.TestFixture[struct{}]{
		Endpoint:    []any{"node", id, "expire"},
		Method:      http.MethodPost,
		SuccessResp: struct{}{},
	}

	testutil.RunResourceTest(t, fixture, func(ctx context.Context, mockReq *requests.MockRequest) (struct{}, error) {
		n := &NodeResource{r: mockReq}
		err := n.Expire(ctx, id)
		return struct{}{}, err
	})
}

func TestNodeResource_Rename(t *testing.T) {
	id := "1"
	name := "newname"
	fixture := testutil.TestFixture[NodeResponse]{
		Endpoint:    []any{"node", id, "rename", name},
		Method:      http.MethodPost,
		SuccessResp: NodeResponse{Node: Node{ID: id, Name: name}},
	}

	testutil.RunResourceTest(t, fixture, func(ctx context.Context, mockReq *requests.MockRequest) (NodeResponse, error) {
		n := &NodeResource{r: mockReq}
		return n.Rename(ctx, id, name)
	})
}

func TestNodeResource_AddTags(t *testing.T) {
	id := "1"
	tags := []string{"tag1", "tag2"}
	fixture := testutil.TestFixture[NodeResponse]{
		Endpoint:    []any{"node", id, "tags"},
		Method:      http.MethodPost,
		SuccessResp: NodeResponse{Node: Node{ID: id, Tags: tags}},
	}

	testutil.RunResourceTest(t, fixture, func(ctx context.Context, mockReq *requests.MockRequest) (NodeResponse, error) {
		n := &NodeResource{r: mockReq}
		return n.AddTags(ctx, id, tags)
	})
}

func TestNodeResource_BackfillIPs(t *testing.T) {
	fixture := testutil.TestFixture[BackfillIPsResponse]{
		Endpoint:    []any{"node", "backfillips"},
		Method:      http.MethodPost,
		SuccessResp: BackfillIPsResponse{Changes: []string{"ip1", "ip2"}},
	}

	testutil.RunResourceTest(t, fixture, func(ctx context.Context, mockReq *requests.MockRequest) (BackfillIPsResponse, error) {
		n := &NodeResource{r: mockReq}
		return n.BackfillIPs(ctx, true)
	})
}

func TestNodeResource_ApproveRoutes(t *testing.T) {
	id := "1"
	routes := []string{"10.0.0.0/24", "192.168.1.0/24"}
	fixture := testutil.TestFixture[NodeResponse]{
		Endpoint:    []any{"node", id, "approve_routes"},
		Method:      http.MethodPost,
		SuccessResp: NodeResponse{Node: Node{ID: id, ApprovedRoutes: routes}},
	}

	testutil.RunResourceTest(t, fixture, func(ctx context.Context, mockReq *requests.MockRequest) (NodeResponse, error) {
		n := &NodeResource{r: mockReq}
		return n.ApproveRoutes(ctx, id, routes)
	})
}
