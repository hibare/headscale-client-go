package apikeys

import (
	"errors"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/hibare/headscale-client-go/requests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAPIKeyResource_List(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		a := &APIKeyResource{R: mockReq}
		ctx := t.Context()
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}
		fakeResp := APIKeysResponse{APIKeys: []APIKey{{ID: "1", Prefix: "prefix1"}}}

		mockReq.On("BuildURL", "apikey").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodGet, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*apikeys.APIKeysResponse")).Run(func(args mock.Arguments) {
			resp := args.Get(2).(*APIKeysResponse) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
			*resp = fakeResp
		}).Return(nil)

		resp, err := a.List(ctx)
		require.NoError(t, err)
		assert.Equal(t, fakeResp, resp)
		mockReq.AssertExpectations(t)
	})

	t.Run("build request error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		a := &APIKeyResource{R: mockReq}
		ctx := t.Context()
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "apikey").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodGet, fakeURL, mock.Anything).Return(fakeReq, errors.New("build error"))

		resp, err := a.List(ctx)
		require.Error(t, err)
		assert.Empty(t, resp.APIKeys)
		mockReq.AssertExpectations(t)
	})

	t.Run("do error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		a := &APIKeyResource{R: mockReq}
		ctx := t.Context()
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "apikey").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodGet, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*apikeys.APIKeysResponse")).Return(errors.New("do error"))

		resp, err := a.List(ctx)
		require.Error(t, err)
		assert.Empty(t, resp.APIKeys)
		mockReq.AssertExpectations(t)
	})
}

func TestAPIKeyResource_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		a := &APIKeyResource{R: mockReq}
		ctx := t.Context()
		expiration := time.Now().Add(24 * time.Hour)
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}
		fakeResp := CreateAPIKeyResponse{APIKey: "new-api-key"}

		mockReq.On("BuildURL", "apikey").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*apikeys.CreateAPIKeyResponse")).Run(func(args mock.Arguments) {
			resp := args.Get(2).(*CreateAPIKeyResponse) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
			*resp = fakeResp
		}).Return(nil)

		resp, err := a.Create(ctx, CreateAPIKeyRequest{Expiration: expiration})
		require.NoError(t, err)
		assert.Equal(t, fakeResp, resp)
		mockReq.AssertExpectations(t)
	})

	t.Run("build request error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		a := &APIKeyResource{R: mockReq}
		ctx := t.Context()
		expiration := time.Now().Add(24 * time.Hour)
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "apikey").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, errors.New("build error"))

		resp, err := a.Create(ctx, CreateAPIKeyRequest{Expiration: expiration})
		require.Error(t, err)
		assert.Empty(t, resp.APIKey)
		mockReq.AssertExpectations(t)
	})

	t.Run("do error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		a := &APIKeyResource{R: mockReq}
		ctx := t.Context()
		expiration := time.Now().Add(24 * time.Hour)
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "apikey").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*apikeys.CreateAPIKeyResponse")).Return(errors.New("do error"))

		resp, err := a.Create(ctx, CreateAPIKeyRequest{Expiration: expiration})
		require.Error(t, err)
		assert.Empty(t, resp.APIKey)
		mockReq.AssertExpectations(t)
	})
}

func TestAPIKeyResource_Expire(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		a := &APIKeyResource{R: mockReq}
		ctx := t.Context()
		prefix := "prefix1"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "apikey", "expire").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, nil).Return(nil)

		err := a.Expire(ctx, prefix)
		require.NoError(t, err)
		mockReq.AssertExpectations(t)
	})

	t.Run("build request error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		a := &APIKeyResource{R: mockReq}
		ctx := t.Context()
		prefix := "prefix1"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "apikey", "expire").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, errors.New("build error"))

		err := a.Expire(ctx, prefix)
		require.Error(t, err)
		mockReq.AssertExpectations(t)
	})

	t.Run("do error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		a := &APIKeyResource{R: mockReq}
		ctx := t.Context()
		prefix := "prefix1"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "apikey", "expire").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPost, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, nil).Return(errors.New("do error"))

		err := a.Expire(ctx, prefix)
		require.Error(t, err)
		mockReq.AssertExpectations(t)
	})
}

func TestAPIKeyResource_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		a := &APIKeyResource{R: mockReq}
		ctx := t.Context()
		prefix := "prefix1"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "apikey", prefix).Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodDelete, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, nil).Return(nil)

		err := a.Delete(ctx, prefix)
		require.NoError(t, err)
		mockReq.AssertExpectations(t)
	})

	t.Run("build request error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		a := &APIKeyResource{R: mockReq}
		ctx := t.Context()
		prefix := "prefix1"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "apikey", prefix).Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodDelete, fakeURL, mock.Anything).Return(fakeReq, errors.New("build error"))

		err := a.Delete(ctx, prefix)
		require.Error(t, err)
		mockReq.AssertExpectations(t)
	})

	t.Run("do error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		a := &APIKeyResource{R: mockReq}
		ctx := t.Context()
		prefix := "prefix1"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "apikey", prefix).Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodDelete, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, nil).Return(errors.New("do error"))

		err := a.Delete(ctx, prefix)
		require.Error(t, err)
		mockReq.AssertExpectations(t)
	})
}
