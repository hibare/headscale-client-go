package policy

import (
	"context"
	"testing"

	"github.com/hibare/headscale-client-go/requests"
	"github.com/hibare/headscale-client-go/v1/testutil"
)

func TestPolicyResource_Get(t *testing.T) {
	fixture := testutil.TestFixture[Policy]{
		Endpoint:    "policy",
		Method:      "GET",
		SuccessResp: Policy{Policy: "test-policy", UpdatedAt: "2024-01-01T00:00:00Z"},
	}

	testutil.RunResourceTest(t, fixture, func(ctx context.Context, mockReq *requests.MockRequest) (Policy, error) {
		p := &PolicyResource{r: mockReq}
		return p.Get(ctx)
	})
}

func TestPolicyResource_Update(t *testing.T) {
	fixture := testutil.TestFixture[UpdatePolicyResponse]{
		Endpoint: "policy",
		Method:   "PUT",
		SuccessResp: UpdatePolicyResponse{
			Policy:    "new-policy",
			UpdatedAt: "2024-01-01T00:00:00Z",
		},
	}

	testutil.RunResourceTest(t, fixture, func(ctx context.Context, mockReq *requests.MockRequest) (UpdatePolicyResponse, error) {
		p := &PolicyResource{r: mockReq}
		return p.Update(ctx, "new-policy")
	})
}
