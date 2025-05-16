package headscale

import (
	"context"
	"net/http"
)

// PolicyResource is a resource for managing policies in Headscale.
type PolicyResource struct {
	Client ClientInterface
}

// Policy represents a policy in Headscale.
type Policy struct {
	Policy    string `json:"policy"`
	UpdatedAt string `json:"updated_at"`
}

// UpdatePolicyRequest represents a request to update the policy.
type UpdatePolicyRequest struct {
	Policy string `json:"policy"`
}

// Get retrieves the current policy from Headscale.
func (p *PolicyResource) Get(ctx context.Context) (Policy, error) {
	var policy Policy

	url := p.Client.buildURL("policy")
	req, err := p.Client.buildRequest(ctx, http.MethodGet, url, requestOptions{})
	if err != nil {
		return policy, err
	}

	err = p.Client.do(ctx, req, &policy)
	return policy, err
}

// Put updates the policy in Headscale.
func (p *PolicyResource) Put(ctx context.Context, policy string) error {
	url := p.Client.buildURL("policy")
	req, err := p.Client.buildRequest(ctx, http.MethodPut, url, requestOptions{
		body: UpdatePolicyRequest{
			Policy: policy,
		},
	})
	if err != nil {
		return err
	}

	return p.Client.do(ctx, req, nil)
}
