package headscale

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserResource_List(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		userResource := &UserResource{Client: client}
		client.(*MockClient).On("buildURL", "user").Return(&url.URL{Path: "http://example.com/api/v1/user"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodGet, &url.URL{Path: "http://example.com/api/v1/user"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			resp := args.Get(1).(*UsersResponse)
			resp.Users = []User{{ID: "1", Name: "test-user"}}
		})

		users, err := userResource.List(context.Background())
		assert.NoError(t, err)
		assert.Len(t, users.Users, 1)
		assert.Equal(t, "test-user", users.Users[0].Name)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		userResource := &UserResource{Client: client}
		client.(*MockClient).On("buildURL", "user").Return(&url.URL{Path: "http://example.com/api/v1/user"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodGet, &url.URL{Path: "http://example.com/api/v1/user"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		users, err := userResource.List(context.Background())
		assert.Error(t, err)
		assert.Empty(t, users.Users)
	})
}

func TestUserResource_Get(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		userResource := &UserResource{Client: client}
		client.(*MockClient).On("buildURL", "user", "1").Return(&url.URL{Path: "http://example.com/api/v1/user/1"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodGet, &url.URL{Path: "http://example.com/api/v1/user/1"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			resp := args.Get(1).(*UserResponse)
			resp.User = User{ID: "1", Name: "test-user"}
		})

		user, err := userResource.Get(context.Background(), "1")
		assert.NoError(t, err)
		assert.Equal(t, "test-user", user.User.Name)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		userResource := &UserResource{Client: client}
		client.(*MockClient).On("buildURL", "user", "1").Return(&url.URL{Path: "http://example.com/api/v1/user/1"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodGet, &url.URL{Path: "http://example.com/api/v1/user/1"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		user, err := userResource.Get(context.Background(), "1")
		assert.Error(t, err)
		assert.Empty(t, user.User)
	})
}

func TestUserResource_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		userResource := &UserResource{Client: client}
		client.(*MockClient).On("buildURL", "user").Return(&url.URL{Path: "http://example.com/api/v1/user"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPost, &url.URL{Path: "http://example.com/api/v1/user"}, requestOptions{
			body: CreateUserRequest{Name: "new-user"},
		}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			resp := args.Get(1).(*User)
			resp.ID = "1"
			resp.Name = "new-user"
		})

		user, err := userResource.Create(context.Background(), "new-user")
		assert.NoError(t, err)
		assert.Equal(t, "new-user", user.Name)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		userResource := &UserResource{Client: client}
		client.(*MockClient).On("buildURL", "user").Return(&url.URL{Path: "http://example.com/api/v1/user"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPost, &url.URL{Path: "http://example.com/api/v1/user"}, requestOptions{
			body: CreateUserRequest{Name: "new-user"},
		}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		user, err := userResource.Create(context.Background(), "new-user")
		assert.Error(t, err)
		assert.Empty(t, user.ID)
	})
}

func TestUserResource_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		userResource := &UserResource{Client: client}
		client.(*MockClient).On("buildURL", "user", "1").Return(&url.URL{Path: "http://example.com/api/v1/user/1"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodDelete, &url.URL{Path: "http://example.com/api/v1/user/1"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil)

		err := userResource.Delete(context.Background(), "1")
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		userResource := &UserResource{Client: client}
		client.(*MockClient).On("buildURL", "user", "1").Return(&url.URL{Path: "http://example.com/api/v1/user/1"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodDelete, &url.URL{Path: "http://example.com/api/v1/user/1"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		err := userResource.Delete(context.Background(), "1")
		assert.Error(t, err)
	})
}

func TestUserResource_Rename(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		userResource := &UserResource{Client: client}
		client.(*MockClient).On("buildURL", "user", "1", "rename", "new-name").Return(&url.URL{Path: "http://example.com/api/v1/user/1/rename/new-name"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPost, &url.URL{Path: "http://example.com/api/v1/user/1/rename/new-name"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil)

		err := userResource.Rename(context.Background(), "1", "new-name")
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		userResource := &UserResource{Client: client}
		client.(*MockClient).On("buildURL", "user", "1", "rename", "new-name").Return(&url.URL{Path: "http://example.com/api/v1/user/1/rename/new-name"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPost, &url.URL{Path: "http://example.com/api/v1/user/1/rename/new-name"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		err := userResource.Rename(context.Background(), "1", "new-name")
		assert.Error(t, err)
	})
}
