package users

import (
	"context"
	"net/http"
	"testing"

	"github.com/hibare/headscale-client-go/requests"
	"github.com/hibare/headscale-client-go/v1/testutil"
)

func TestUserResource_List(t *testing.T) {
	fixture := testutil.TestFixture[UsersResponse]{
		Endpoint:    "user",
		Method:      http.MethodGet,
		SuccessResp: UsersResponse{Users: []User{{ID: "1", Name: "test"}}},
	}

	testutil.RunResourceTest(t, fixture, func(ctx context.Context, mockReq *requests.MockRequest) (UsersResponse, error) {
		u := &UserResource{r: mockReq}
		return u.List(ctx, UserListFilter{})
	})
}

func TestUserResource_Create(t *testing.T) {
	fixture := testutil.TestFixture[UserResponse]{
		Endpoint:    "user",
		Method:      http.MethodPost,
		SuccessResp: UserResponse{User: User{ID: "1", Name: "test"}},
	}

	testutil.RunResourceTest(t, fixture, func(ctx context.Context, mockReq *requests.MockRequest) (UserResponse, error) {
		u := &UserResource{r: mockReq}
		return u.Create(ctx, CreateUserRequest{
			Name:        "test",
			DisplayName: "test",
			Email:       "test@example.com",
			PictureURL:  "https://example.com/picture.png",
		})
	})
}

func TestUserResource_Delete(t *testing.T) {
	fixture := testutil.TestFixture[struct{}]{
		Endpoint:    []any{"user", "1"},
		Method:      http.MethodDelete,
		SuccessResp: struct{}{},
	}

	testutil.RunResourceTest(t, fixture, func(ctx context.Context, mockReq *requests.MockRequest) (struct{}, error) {
		u := &UserResource{r: mockReq}
		err := u.Delete(ctx, "1")
		return struct{}{}, err
	})
}

func TestUserResource_Rename(t *testing.T) {
	fixture := testutil.TestFixture[UserResponse]{
		Endpoint:    []any{"user", "1", "rename", "new-name"},
		Method:      http.MethodPost,
		SuccessResp: UserResponse{User: User{ID: "1", Name: "new-name"}},
	}

	testutil.RunResourceTest(t, fixture, func(ctx context.Context, mockReq *requests.MockRequest) (UserResponse, error) {
		u := &UserResource{r: mockReq}
		return u.Rename(ctx, "1", "new-name")
	})
}
