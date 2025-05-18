// Package nodes provides a client for managing nodes in Headscale.
package nodes

import (
	"context"
	"net/http"
	"time"

	"github.com/hibare/headscale-client-go/requests"
	"github.com/hibare/headscale-client-go/v1/preauthkeys"
	"github.com/hibare/headscale-client-go/v1/users"
)

// NodeResourceInterface is an interface for managing nodes in Headscale.
type NodeResourceInterface interface {
	List(ctx context.Context, filter NodeListFilter) (NodesResponse, error)
	Get(ctx context.Context, id string) (NodeResponse, error)
	Register(ctx context.Context, user, key string) (NodeResponse, error)
	Delete(ctx context.Context, id string) error
	Expire(ctx context.Context, id string) error
	Rename(ctx context.Context, id, name string) (NodeResponse, error)
	AddTags(ctx context.Context, id string, tags []string) (NodeResponse, error)
	UpdateUser(ctx context.Context, id, user string) (NodeResponse, error)
	BackFillIP(ctx context.Context, id string) (BackfillIPsResponse, error)
}

// Node represents a node in Headscale.
type Node struct {
	ID              string                 `json:"id"`
	MachineKey      string                 `json:"machineKey"`
	NodeKey         string                 `json:"nodeKey"`
	DiscoKey        string                 `json:"discoKey"`
	IPAddresses     []string               `json:"ipAddresses"`
	Name            string                 `json:"name"`
	User            users.User             `json:"user"`
	LastSeen        time.Time              `json:"lastSeen"`
	Expiry          time.Time              `json:"expiry"`
	PreAuthKey      preauthkeys.PreAuthKey `json:"preAuthKey"`
	CreatedAt       time.Time              `json:"createdAt"`
	RegisterMethod  string                 `json:"registerMethod"`
	ForcedTags      []string               `json:"forcedTags"`
	InvalidTags     []string               `json:"invalidTags"`
	ValidTags       []string               `json:"validTags"`
	GivenName       string                 `json:"givenName"`
	Online          bool                   `json:"online"`
	ApprovedRoutes  []string               `json:"approvedRoutes"`
	AvailableRoutes []string               `json:"availableRoutes"`
	SubnetRoutes    []string               `json:"subnetRoutes"`
}

// NodeResponse represents a single node response from the API.
type NodeResponse struct {
	Node Node `json:"node"`
}

// NodesResponse represents a list of nodes response from the API.
//
//nolint:revive // This is a struct for a response from the API.
type NodesResponse struct {
	Nodes []Node `json:"nodes"`
}

// NodeListFilter represents a filter for listing nodes.
type NodeListFilter struct {
	User string `json:"user"`
}

// List returns a list of nodes from the Headscale.
func (n *NodeResource) List(ctx context.Context, filter NodeListFilter) (NodesResponse, error) {
	var nodes NodesResponse

	queryParams := map[string]any{}
	if filter.User != "" {
		queryParams["user"] = filter.User
	}

	url := n.r.BuildURL("node")
	req, err := n.r.BuildRequest(ctx, http.MethodGet, url, requests.RequestOptions{
		QueryParams: queryParams,
	})
	if err != nil {
		return nodes, err
	}

	err = n.r.Do(ctx, req, &nodes)
	return nodes, err
}

// Get retrieves a node by its ID from the Headscale.
func (n *NodeResource) Get(ctx context.Context, id string) (NodeResponse, error) {
	var node NodeResponse

	url := n.r.BuildURL("node", id)
	req, err := n.r.BuildRequest(ctx, http.MethodGet, url, requests.RequestOptions{})
	if err != nil {
		return node, err
	}

	err = n.r.Do(ctx, req, &node)
	return node, err
}

// Register registers a new node with the Headscale.
func (n *NodeResource) Register(ctx context.Context, user, key string) (NodeResponse, error) {
	var node NodeResponse

	url := n.r.BuildURL("node")
	req, err := n.r.BuildRequest(ctx, http.MethodPost, url, requests.RequestOptions{
		QueryParams: map[string]interface{}{
			"user": user,
			"key":  key,
		},
	})
	if err != nil {
		return node, err
	}

	err = n.r.Do(ctx, req, &node)
	return node, err
}

// Delete removes a node from the Headscale.
func (n *NodeResource) Delete(ctx context.Context, id string) error {
	url := n.r.BuildURL("node", id)
	req, err := n.r.BuildRequest(ctx, http.MethodDelete, url, requests.RequestOptions{})
	if err != nil {
		return err
	}

	return n.r.Do(ctx, req, nil)
}

// ApproveRoutesRequest represents a request to approve routes for a node.
type ApproveRoutesRequest struct {
	Routes []string `json:"routes"`
}

// ApproveRoutes approves routes for a node in the Headscale.
func (n *NodeResource) ApproveRoutes(ctx context.Context, id string, routes []string) (NodeResponse, error) {
	var node NodeResponse

	url := n.r.BuildURL("node", id, "approve_routes")
	req, err := n.r.BuildRequest(ctx, http.MethodPost, url, requests.RequestOptions{
		Body: ApproveRoutesRequest{Routes: routes},
	})
	if err != nil {
		return node, err
	}

	err = n.r.Do(ctx, req, &node)
	return node, err
}

// Expire marks a node as expired in the Headscale.
func (n *NodeResource) Expire(ctx context.Context, id string) error {
	url := n.r.BuildURL("node", id, "expire")
	req, err := n.r.BuildRequest(ctx, http.MethodPost, url, requests.RequestOptions{})
	if err != nil {
		return err
	}

	return n.r.Do(ctx, req, nil)
}

// Rename renames a node in the Headscale.
func (n *NodeResource) Rename(ctx context.Context, id, name string) (NodeResponse, error) {
	var node NodeResponse

	url := n.r.BuildURL("node", id, "rename", name)
	req, err := n.r.BuildRequest(ctx, http.MethodPost, url, requests.RequestOptions{})
	if err != nil {
		return node, err
	}

	err = n.r.Do(ctx, req, &node)
	return node, err
}

// AddTagsRequest represents a request to add tags to a node.
type AddTagsRequest struct {
	Tags []string `json:"tags"`
}

// AddTags adds tags to a node in the Headscale.
func (n *NodeResource) AddTags(ctx context.Context, id string, tags []string) (NodeResponse, error) {
	var node NodeResponse

	url := n.r.BuildURL("node", id, "tags")
	req, err := n.r.BuildRequest(ctx, http.MethodPost, url, requests.RequestOptions{
		Body: AddTagsRequest{Tags: tags},
	})
	if err != nil {
		return node, err
	}

	err = n.r.Do(ctx, req, &node)
	return node, err
}

// UpdateUserRequest represents a request to update the user associated with a node.
type UpdateUserRequest struct {
	User string `json:"user"`
}

// UpdateUser updates the user associated with a node in the Headscale.
func (n *NodeResource) UpdateUser(ctx context.Context, id, user string) (NodeResponse, error) {
	var node NodeResponse

	url := n.r.BuildURL("node", id, "user")
	req, err := n.r.BuildRequest(ctx, http.MethodPost, url, requests.RequestOptions{
		Body: UpdateUserRequest{User: user},
	})
	if err != nil {
		return node, err
	}

	err = n.r.Do(ctx, req, &node)
	return node, err
}

// BackfillIPsResponse represents a response from the backfill IP endpoint.
type BackfillIPsResponse struct {
	Changes []string `json:"changes"`
}

// BackFillIP backfills the IP address for a node in the Headscale.
func (n *NodeResource) BackFillIP(ctx context.Context, id string) (BackfillIPsResponse, error) {
	url := n.r.BuildURL("node", id, "backfill_ip")
	req, err := n.r.BuildRequest(ctx, http.MethodPost, url, requests.RequestOptions{})
	if err != nil {
		return BackfillIPsResponse{}, err
	}

	var backfillIPs BackfillIPsResponse
	err = n.r.Do(ctx, req, &backfillIPs)
	return backfillIPs, err
}

// NodeResource is a struct that provides methods to interact with the nodes API of Headscale.
type NodeResource struct {
	r requests.RequestInterface
}

// NewNodeResource creates a new NodeResource.
func NewNodeResource(r requests.RequestInterface) *NodeResource {
	return &NodeResource{
		r: r,
	}
}
