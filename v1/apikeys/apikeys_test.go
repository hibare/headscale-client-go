package apikeys

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/hibare/headscale-client-go/requests"
	"github.com/hibare/headscale-client-go/v1/testutil"
)

func TestAPIKeyResource_List(t *testing.T) {
	fixture := testutil.TestFixture[APIKeysResponse]{
		Endpoint:    "apikey",
		Method:      http.MethodGet,
		SuccessResp: APIKeysResponse{APIKeys: []APIKey{{ID: "1", Prefix: "prefix1"}}},
	}

	testutil.RunResourceTest(t, fixture, func(ctx context.Context, mockReq *requests.MockRequest) (APIKeysResponse, error) {
		a := &APIKeyResource{r: mockReq}
		return a.List(ctx)
	})
}

func TestAPIKeyResource_Create(t *testing.T) {
	expiration := time.Now().Add(24 * time.Hour)
	fixture := testutil.TestFixture[CreateAPIKeyResponse]{
		Endpoint:    "apikey",
		Method:      http.MethodPost,
		SuccessResp: CreateAPIKeyResponse{APIKey: "new-api-key"},
	}

	testutil.RunResourceTest(t, fixture, func(ctx context.Context, mockReq *requests.MockRequest) (CreateAPIKeyResponse, error) {
		a := &APIKeyResource{r: mockReq}
		return a.Create(ctx, CreateAPIKeyRequest{Expiration: expiration})
	})
}

func TestAPIKeyResource_Expire(t *testing.T) {
	prefix := "prefix1"
	fixture := testutil.TestFixture[struct{}]{
		Endpoint:    []any{"apikey", "expire"},
		Method:      http.MethodPost,
		SuccessResp: struct{}{},
	}

	testutil.RunResourceTest(t, fixture, func(ctx context.Context, mockReq *requests.MockRequest) (struct{}, error) {
		a := &APIKeyResource{r: mockReq}
		err := a.Expire(ctx, prefix)
		return struct{}{}, err
	})
}

func TestAPIKeyResource_Delete(t *testing.T) {
	prefix := "prefix1"
	fixture := testutil.TestFixture[struct{}]{
		Endpoint:    []any{"apikey", prefix},
		Method:      http.MethodDelete,
		SuccessResp: struct{}{},
	}

	testutil.RunResourceTest(t, fixture, func(ctx context.Context, mockReq *requests.MockRequest) (struct{}, error) {
		a := &APIKeyResource{r: mockReq}
		err := a.Delete(ctx, prefix)
		return struct{}{}, err
	})
}

func TestAPIKeyResource_ExpireByID(t *testing.T) {
	id := "1"
	fixture := testutil.TestFixture[struct{}]{
		Endpoint:    []any{"apikey", "expire"},
		Method:      http.MethodPost,
		SuccessResp: struct{}{},
	}

	testutil.RunResourceTest(t, fixture, func(ctx context.Context, mockReq *requests.MockRequest) (struct{}, error) {
		a := &APIKeyResource{r: mockReq}
		err := a.ExpireByID(ctx, id)
		return struct{}{}, err
	})
}
