package preauthkeys

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/hibare/headscale-client-go/requests"
	"github.com/hibare/headscale-client-go/v1/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestPreAuthKeyResource_List(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		p := &PreAuthKeyResource{r: mockReq}
		ctx := context.Background()
		filter := PreAuthKeyListFilter{User: "testuser"}
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}
		fakeResp := PreAuthKeysResponse{PreAuthKeys: []PreAuthKey{{ID: "1", User: users.User{ID: "u1", Name: "testuser"}}}}

		mockReq.On("BuildURL", "preauthkey").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodGet, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*preauthkeys.PreAuthKeysResponse")).Run(func(args mock.Arguments) {
			resp := args.Get(2).(*PreAuthKeysResponse) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
			*resp = fakeResp
		}).Return(nil)

		resp, err := p.List(ctx, filter)
		require.NoError(t, err)
		assert.Equal(t, fakeResp, resp)
		mockReq.AssertExpectations(t)
	})

	t.Run("build request error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		p := &PreAuthKeyResource{r: mockReq}
		ctx := context.Background()
		filter := PreAuthKeyListFilter{}
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "preauthkey").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodGet, fakeURL, mock.Anything).Return(fakeReq, errors.New("build error"))

		resp, err := p.List(ctx, filter)
		require.Error(t, err)
		assert.Empty(t, resp.PreAuthKeys)
		mockReq.AssertExpectations(t)
	})

	t.Run("do error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		p := &PreAuthKeyResource{r: mockReq}
		ctx := context.Background()
		filter := PreAuthKeyListFilter{}
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "preauthkey").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodGet, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*preauthkeys.PreAuthKeysResponse")).Return(errors.New("do error"))

		resp, err := p.List(ctx, filter)
		require.Error(t, err)
		assert.Empty(t, resp.PreAuthKeys)
		mockReq.AssertExpectations(t)
	})
}

func TestPreAuthKeyResource_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		p := &PreAuthKeyResource{r: mockReq}
		ctx := context.Background()
		user := "testuser"
		reusable := true
		ephemeral := false
		expiration := time.Now().Add(24 * time.Hour)
		aclTags := []string{"tag:test"}
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}
		fakeResp := PreAuthKeyResponse{PreAuthKey: []PreAuthKey{{ID: "1", User: users.User{ID: "u1", Name: user}}}}

		mockReq.On("BuildURL", "preauthkey").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*preauthkeys.PreAuthKeyResponse")).Run(func(args mock.Arguments) {
			resp := args.Get(2).(*PreAuthKeyResponse) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
			*resp = fakeResp
		}).Return(nil)

		resp, err := p.Create(ctx, user, reusable, ephemeral, expiration, aclTags)
		require.NoError(t, err)
		assert.Equal(t, fakeResp, resp)
		mockReq.AssertExpectations(t)
	})

	t.Run("build request error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		p := &PreAuthKeyResource{r: mockReq}
		ctx := context.Background()
		user := "testuser"
		reusable := true
		ephemeral := false
		expiration := time.Now().Add(24 * time.Hour)
		aclTags := []string{"tag:test"}
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "preauthkey").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, errors.New("build error"))

		resp, err := p.Create(ctx, user, reusable, ephemeral, expiration, aclTags)
		require.Error(t, err)
		assert.Empty(t, resp.PreAuthKey)
		mockReq.AssertExpectations(t)
	})

	t.Run("do error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		p := &PreAuthKeyResource{r: mockReq}
		ctx := context.Background()
		user := "testuser"
		reusable := true
		ephemeral := false
		expiration := time.Now().Add(24 * time.Hour)
		aclTags := []string{"tag:test"}
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "preauthkey").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*preauthkeys.PreAuthKeyResponse")).Return(errors.New("do error"))

		resp, err := p.Create(ctx, user, reusable, ephemeral, expiration, aclTags)
		require.Error(t, err)
		assert.Empty(t, resp.PreAuthKey)
		mockReq.AssertExpectations(t)
	})
}

func TestPreAuthKeyResource_Expire(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		p := &PreAuthKeyResource{r: mockReq}
		ctx := context.Background()
		user := "testuser"
		key := "key1"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "preauthkey", "expire").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, nil).Return(nil)

		err := p.Expire(ctx, user, key)
		require.NoError(t, err)
		mockReq.AssertExpectations(t)
	})

	t.Run("build request error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		p := &PreAuthKeyResource{r: mockReq}
		ctx := context.Background()
		user := "testuser"
		key := "key1"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "preauthkey", "expire").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, errors.New("build error"))

		err := p.Expire(ctx, user, key)
		require.Error(t, err)
		mockReq.AssertExpectations(t)
	})

	t.Run("do error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		p := &PreAuthKeyResource{r: mockReq}
		ctx := context.Background()
		user := "testuser"
		key := "key1"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "preauthkey", "expire").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, nil).Return(errors.New("do error"))

		err := p.Expire(ctx, user, key)
		require.Error(t, err)
		mockReq.AssertExpectations(t)
	})
}
