package preauthkeys

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/hibare/headscale-client-go/requests"
	"github.com/hibare/headscale-client-go/v1/testutil"
	"github.com/hibare/headscale-client-go/v1/users"
)

func TestPreAuthKeyResource_List(t *testing.T) {
	fixture := testutil.TestFixture[PreAuthKeysResponse]{
		Endpoint: "preauthkey",
		Method:   http.MethodGet,
		SuccessResp: PreAuthKeysResponse{
			PreAuthKeys: []PreAuthKey{
				{
					ID:   "1",
					User: users.User{ID: "u1", Name: "testuser"},
				},
			},
		},
	}

	testutil.RunResourceTest(t, fixture, func(ctx context.Context, mockReq *requests.MockRequest) (PreAuthKeysResponse, error) {
		p := &PreAuthKeyResource{r: mockReq}
		return p.List(ctx)
	})
}

func TestPreAuthKeyResource_Create(t *testing.T) {
	request := CreatePreAuthKeyRequest{
		User:       "testuser",
		Reusable:   true,
		Ephemeral:  false,
		Expiration: time.Now().Add(24 * time.Hour),
		ACLTags:    []string{"tag:test"},
	}

	fixture := testutil.TestFixture[PreAuthKeyResponse]{
		Endpoint: "preauthkey",
		Method:   http.MethodPost,
		SuccessResp: PreAuthKeyResponse{
			PreAuthKey: PreAuthKey{
				ID:   "1",
				User: users.User{ID: "u1", Name: "testuser"},
			},
		},
	}

	testutil.RunResourceTest(t, fixture, func(ctx context.Context, mockReq *requests.MockRequest) (PreAuthKeyResponse, error) {
		p := &PreAuthKeyResource{r: mockReq}
		return p.Create(ctx, request)
	})
}

func TestPreAuthKeyResource_Expire(t *testing.T) {
	id := "1"
	fixture := testutil.TestFixture[struct{}]{
		Endpoint:    []any{"preauthkey", "expire"},
		Method:      http.MethodPost,
		SuccessResp: struct{}{},
	}

	testutil.RunResourceTest(t, fixture, func(ctx context.Context, mockReq *requests.MockRequest) (struct{}, error) {
		p := &PreAuthKeyResource{r: mockReq}
		err := p.Expire(ctx, id)
		return struct{}{}, err
	})
}

func TestPreAuthKeyResource_Delete(t *testing.T) {
	id := "1"
	fixture := testutil.TestFixture[struct{}]{
		Endpoint:    "preauthkey",
		Method:      http.MethodDelete,
		SuccessResp: struct{}{},
	}

	testutil.RunResourceTest(t, fixture, func(ctx context.Context, mockReq *requests.MockRequest) (struct{}, error) {
		p := &PreAuthKeyResource{r: mockReq}
		err := p.Delete(ctx, id)
		return struct{}{}, err
	})
}
