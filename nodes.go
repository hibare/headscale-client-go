package headscale

import (
	"context"
	"net/http"
	"time"
)

// Node represents a node in Headscale.
type Node struct {
	ID             string     `json:"id"`
	MachineKey     string     `json:"machineKey"`
	NodeKey        string     `json:"nodeKey"`
	DiscoKey       string     `json:"discoKey"`
	IPAddresses    []string   `json:"ipAddresses"`
	Name           string     `json:"name"`
	User           User       `json:"user"`
	LastSeen       time.Time  `json:"lastSeen"`
	Expiry         time.Time  `json:"expiry"`
	PreAuthKey     PreAuthKey `json:"preAuthKey"`
	CreatedAt      time.Time  `json:"createdAt"`
	RegisterMethod string     `json:"registerMethod"`
	ForcedTags     []string   `json:"forcedTags"`
	InvalidTags    []string   `json:"invalidTags"`
	ValidTags      []string   `json:"validTags"`
	GivenName      string     `json:"givenName"`
	Online         bool       `json:"online"`
}

// NodeResponse represents a single node response from the API.
type NodeResponse struct {
	Node Node `json:"node"`
}

// NodesResponse represents a list of nodes response from the API.
type NodesResponse struct {
	Nodes []Node `json:"nodes"`
}

// NodeResource is a struct that provides methods to interact with the nodes API of Headscale.
type NodeResource struct {
	Client ClientInterface
}

// List returns a list of nodes from the Headscale.
func (n *NodeResource) List(ctx context.Context) (NodesResponse, error) {
	var nodes NodesResponse

	url := n.Client.buildURL("node")
	req, err := n.Client.buildRequest(ctx, http.MethodGet, url, requestOptions{})
	if err != nil {
		return nodes, err
	}

	err = n.Client.do(ctx, req, &nodes)
	return nodes, err
}

// Get retrieves a node by its ID from the Headscale.
func (n *NodeResource) Get(ctx context.Context, id string) (NodeResponse, error) {
	var node NodeResponse

	url := n.Client.buildURL("node", id)
	req, err := n.Client.buildRequest(ctx, http.MethodGet, url, requestOptions{})
	if err != nil {
		return node, err
	}

	err = n.Client.do(ctx, req, &node)
	return node, err
}

// Register registers a new node with the Headscale.
func (n *NodeResource) Register(ctx context.Context, user, key string) (NodeResponse, error) {
	var node NodeResponse

	url := n.Client.buildURL("node")
	req, err := n.Client.buildRequest(ctx, http.MethodPost, url, requestOptions{
		queryParams: map[string]interface{}{
			"user": user,
			"key":  key,
		},
	})
	if err != nil {
		return node, err
	}

	err = n.Client.do(ctx, req, &node)
	return node, err
}

// Delete removes a node from the Headscale.
func (n *NodeResource) Delete(ctx context.Context, id string) error {
	url := n.Client.buildURL("node", id)
	req, err := n.Client.buildRequest(ctx, http.MethodDelete, url, requestOptions{})
	if err != nil {
		return err
	}

	return n.Client.do(ctx, req, nil)
}

// Expire marks a node as expired in the Headscale.
func (n *NodeResource) Expire(ctx context.Context, id string) error {
	url := n.Client.buildURL("node", id, "expire")
	req, err := n.Client.buildRequest(ctx, http.MethodPost, url, requestOptions{})
	if err != nil {
		return err
	}

	return n.Client.do(ctx, req, nil)
}

// Rename renames a node in the Headscale.
func (n *NodeResource) Rename(ctx context.Context, id, name string) (NodeResponse, error) {
	var node NodeResponse

	url := n.Client.buildURL("node", id, "rename", name)
	req, err := n.Client.buildRequest(ctx, http.MethodPost, url, requestOptions{})
	if err != nil {
		return node, err
	}

	err = n.Client.do(ctx, req, &node)
	return node, err
}

// GetRoutes retrieves the routes for a node by its ID from the Headscale.
func (n *NodeResource) GetRoutes(ctx context.Context, id string) (RoutesResponse, error) {
	var routes RoutesResponse

	url := n.Client.buildURL("node", id, "routes")
	req, err := n.Client.buildRequest(ctx, http.MethodGet, url, requestOptions{})
	if err != nil {
		return routes, err
	}

	err = n.Client.do(ctx, req, &routes)
	return routes, err
}

// AddTagsRequest represents a request to add tags to a node.
type AddTagsRequest struct {
	Tags []string `json:"tags"`
}

// AddTags adds tags to a node in the Headscale.
func (n *NodeResource) AddTags(ctx context.Context, id string, tags []string) (NodeResponse, error) {
	var node NodeResponse

	url := n.Client.buildURL("node", id, "tags")
	req, err := n.Client.buildRequest(ctx, http.MethodPost, url, requestOptions{
		body: AddTagsRequest{Tags: tags},
	})
	if err != nil {
		return node, err
	}

	err = n.Client.do(ctx, req, &node)
	return node, err
}

// UpdateUser updates the user associated with a node in the Headscale.
func (n *NodeResource) UpdateUser(ctx context.Context, id, user string) (NodeResponse, error) {
	var node NodeResponse

	url := n.Client.buildURL("node", id, "user")
	req, err := n.Client.buildRequest(ctx, http.MethodPost, url, requestOptions{
		queryParams: map[string]interface{}{
			"user": user,
		},
	})
	if err != nil {
		return node, err
	}

	err = n.Client.do(ctx, req, &node)
	return node, err
}
