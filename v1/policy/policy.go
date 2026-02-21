// Package policy provides a client for managing policies in Headscale.
package policy

import (
	"context"
	"net/http"

	"github.com/hibare/headscale-client-go/requests"
)

// PolicyResourceInterface is an interface for managing policies in Headscale.
//
//nolint:revive // This is an interface for the policy resource.
type PolicyResourceInterface interface {
	Get(ctx context.Context) (Policy, error)
	Update(ctx context.Context, policy string) (UpdatePolicyResponse, error)
}

// Policy represents a policy in Headscale.
type Policy struct {
	Policy    string `json:"policy"`
	UpdatedAt string `json:"updatedAt"`
}

// UpdatePolicyRequest represents a request to update the policy.
type UpdatePolicyRequest struct {
	Policy string `json:"policy"`
}

// Get retrieves the current policy from Headscale.
func (p *PolicyResource) Get(ctx context.Context) (Policy, error) {
	var policy Policy

	url := p.R.BuildURL("policy")
	req, err := p.R.BuildRequest(ctx, http.MethodGet, url, requests.RequestOptions{})
	if err != nil {
		return policy, err
	}

	err = p.R.Do(ctx, req, &policy)
	return policy, err
}

// UpdatePolicyResponse represents a response from the update policy endpoint.
type UpdatePolicyResponse struct {
	Policy    string `json:"policy"`
	UpdatedAt string `json:"updatedAt"`
}

// Update updates the policy in Headscale.
func (p *PolicyResource) Update(ctx context.Context, policy string) (UpdatePolicyResponse, error) {
	var updatePolicy UpdatePolicyResponse

	url := p.R.BuildURL("policy")
	req, err := p.R.BuildRequest(ctx, http.MethodPut, url, requests.RequestOptions{
		Body: UpdatePolicyRequest{
			Policy: policy,
		},
	})
	if err != nil {
		return updatePolicy, err
	}

	err = p.R.Do(ctx, req, &updatePolicy)
	return updatePolicy, err
}

// PolicyResource is a struct that implements the PolicyResourceInterface.
//
//nolint:revive // This is a resource for the policy.
type PolicyResource struct {
	R requests.RequestInterface
}
