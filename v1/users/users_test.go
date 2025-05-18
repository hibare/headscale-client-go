package users

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/hibare/headscale-client-go/requests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUserResource_List(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		u := &UserResource{r: mockReq}
		ctx := context.Background()
		filter := UserListFilter{ID: "1", Name: "test", Email: "test@example.com"}
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}
		fakeResp := UsersResponse{Users: []User{{ID: "1", Name: "test"}}}

		mockReq.On("BuildURL", "user").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodGet, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*users.UsersResponse")).Run(func(args mock.Arguments) {
			resp := args.Get(2).(*UsersResponse) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
			*resp = fakeResp
		}).Return(nil)

		resp, err := u.List(ctx, filter)
		require.NoError(t, err)
		assert.Equal(t, fakeResp, resp)
		mockReq.AssertExpectations(t)
	})

	t.Run("build request error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		u := &UserResource{r: mockReq}
		ctx := context.Background()
		filter := UserListFilter{}
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "user").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodGet, fakeURL, mock.Anything).Return(fakeReq, errors.New("build error"))

		resp, err := u.List(ctx, filter)
		require.Error(t, err)
		assert.Empty(t, resp.Users)
		mockReq.AssertExpectations(t)
	})

	t.Run("do error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		u := &UserResource{r: mockReq}
		ctx := context.Background()
		filter := UserListFilter{}
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "user").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodGet, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*users.UsersResponse")).Return(errors.New("do error"))

		resp, err := u.List(ctx, filter)
		require.Error(t, err)
		assert.Empty(t, resp.Users)
		mockReq.AssertExpectations(t)
	})
}

func TestUserResource_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		u := &UserResource{r: mockReq}
		ctx := context.Background()
		name := "test"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}
		fakeResp := UserResponse{User: User{ID: "1", Name: name}}

		mockReq.On("BuildURL", "user").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*users.UserResponse")).Run(func(args mock.Arguments) {
			resp := args.Get(2).(*UserResponse) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
			*resp = fakeResp
		}).Return(nil)

		resp, err := u.Create(ctx, name)
		require.NoError(t, err)
		assert.Equal(t, fakeResp, resp)
		mockReq.AssertExpectations(t)
	})

	t.Run("build request error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		u := &UserResource{r: mockReq}
		ctx := context.Background()
		name := "test"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "user").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, errors.New("build error"))

		resp, err := u.Create(ctx, name)
		require.Error(t, err)
		assert.Empty(t, resp.User)
		mockReq.AssertExpectations(t)
	})

	t.Run("do error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		u := &UserResource{r: mockReq}
		ctx := context.Background()
		name := "test"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "user").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*users.UserResponse")).Return(errors.New("do error"))

		resp, err := u.Create(ctx, name)
		require.Error(t, err)
		assert.Empty(t, resp.User)
		mockReq.AssertExpectations(t)
	})
}

func TestUserResource_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		u := &UserResource{r: mockReq}
		ctx := context.Background()
		id := "1"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "user", id).Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodDelete, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, nil).Return(nil)

		err := u.Delete(ctx, id)
		require.NoError(t, err)
		mockReq.AssertExpectations(t)
	})

	t.Run("build request error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		u := &UserResource{r: mockReq}
		ctx := context.Background()
		id := "1"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "user", id).Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodDelete, fakeURL, mock.Anything).Return(fakeReq, errors.New("build error"))

		err := u.Delete(ctx, id)
		require.Error(t, err)
		mockReq.AssertExpectations(t)
	})

	t.Run("do error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		u := &UserResource{r: mockReq}
		ctx := context.Background()
		id := "1"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "user", id).Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodDelete, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, nil).Return(errors.New("do error"))

		err := u.Delete(ctx, id)
		require.Error(t, err)
		mockReq.AssertExpectations(t)
	})
}

func TestUserResource_Rename(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		u := &UserResource{r: mockReq}
		ctx := context.Background()
		id := "1"
		newName := "new-name"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}
		fakeResp := UserResponse{User: User{ID: id, Name: newName}}

		mockReq.On("BuildURL", "user", id, "rename", newName).Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*users.UserResponse")).Run(func(args mock.Arguments) {
			resp := args.Get(2).(*UserResponse) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
			*resp = fakeResp
		}).Return(nil)

		resp, err := u.Rename(ctx, id, newName)
		require.NoError(t, err)
		assert.Equal(t, fakeResp, resp)
		mockReq.AssertExpectations(t)
	})

	t.Run("build request error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		u := &UserResource{r: mockReq}
		ctx := context.Background()
		id := "1"
		newName := "new-name"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "user", id, "rename", newName).Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, errors.New("build error"))

		resp, err := u.Rename(ctx, id, newName)
		require.Error(t, err)
		assert.Empty(t, resp.User)
		mockReq.AssertExpectations(t)
	})

	t.Run("do error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		u := &UserResource{r: mockReq}
		ctx := context.Background()
		id := "1"
		newName := "new-name"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "user", id, "rename", newName).Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*users.UserResponse")).Return(errors.New("do error"))

		resp, err := u.Rename(ctx, id, newName)
		require.Error(t, err)
		assert.Empty(t, resp.User)
		mockReq.AssertExpectations(t)
	})
}
