package headscale

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAPIKeyResource_List(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		apiKeyResource := &APIKeyResource{Client: client}
		client.(*MockClient).On("buildURL", "apikey").Return(&url.URL{Path: "http://example.com/api/v1/apikey"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodGet, &url.URL{Path: "http://example.com/api/v1/apikey"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			resp := args.Get(1).(*APIKeysResponse)
			resp.APIKeys = []APIKey{{ID: "1", Prefix: "test-prefix"}}
		})

		keys, err := apiKeyResource.List(context.Background())
		assert.NoError(t, err)
		assert.Len(t, keys.APIKeys, 1)
		assert.Equal(t, "test-prefix", keys.APIKeys[0].Prefix)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		apiKeyResource := &APIKeyResource{Client: client}
		client.(*MockClient).On("buildURL", "apikey").Return(&url.URL{Path: "http://example.com/api/v1/apikey"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodGet, &url.URL{Path: "http://example.com/api/v1/apikey"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		keys, err := apiKeyResource.List(context.Background())
		assert.Error(t, err)
		assert.Empty(t, keys.APIKeys)
	})
}

func TestAPIKeyResource_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		apiKeyResource := &APIKeyResource{Client: client}
		expiration := time.Now().Add(24 * time.Hour)
		client.(*MockClient).On("buildURL", "apikey").Return(&url.URL{Path: "http://example.com/api/v1/apikey"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPost, &url.URL{Path: "http://example.com/api/v1/apikey"}, requestOptions{
			body: AddAPIKeyRequest{Expiration: expiration},
		}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			resp := args.Get(1).(*APIKey)
			resp.ID = "1"
			resp.Prefix = "test-prefix"
		})

		key, err := apiKeyResource.Create(context.Background(), expiration)
		assert.NoError(t, err)
		assert.Equal(t, "test-prefix", key.Prefix)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		apiKeyResource := &APIKeyResource{Client: client}
		expiration := time.Now().Add(24 * time.Hour)
		client.(*MockClient).On("buildURL", "apikey").Return(&url.URL{Path: "http://example.com/api/v1/apikey"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPost, &url.URL{Path: "http://example.com/api/v1/apikey"}, requestOptions{
			body: AddAPIKeyRequest{Expiration: expiration},
		}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		key, err := apiKeyResource.Create(context.Background(), expiration)
		assert.Error(t, err)
		assert.Empty(t, key.ID)
	})
}

func TestAPIKeyResource_Expire(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		apiKeyResource := &APIKeyResource{Client: client}
		client.(*MockClient).On("buildURL", "apikey", "expire").Return(&url.URL{Path: "http://example.com/api/v1/apikey/expire"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPost, &url.URL{Path: "http://example.com/api/v1/apikey/expire"}, requestOptions{
			body: ExpireAPIKeyRequest{Prefix: "test-prefix"},
		}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil)

		err := apiKeyResource.Expire(context.Background(), "test-prefix")
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		apiKeyResource := &APIKeyResource{Client: client}
		client.(*MockClient).On("buildURL", "apikey", "expire").Return(&url.URL{Path: "http://example.com/api/v1/apikey/expire"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPost, &url.URL{Path: "http://example.com/api/v1/apikey/expire"}, requestOptions{
			body: ExpireAPIKeyRequest{Prefix: "test-prefix"},
		}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		err := apiKeyResource.Expire(context.Background(), "test-prefix")
		assert.Error(t, err)
	})
}

func TestAPIKeyResource_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		apiKeyResource := &APIKeyResource{Client: client}
		client.(*MockClient).On("buildURL", "apikey", "test-prefix").Return(&url.URL{Path: "http://example.com/api/v1/apikey/test-prefix"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodDelete, &url.URL{Path: "http://example.com/api/v1/apikey/test-prefix"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil)

		err := apiKeyResource.Delete(context.Background(), "test-prefix")
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		apiKeyResource := &APIKeyResource{Client: client}
		client.(*MockClient).On("buildURL", "apikey", "test-prefix").Return(&url.URL{Path: "http://example.com/api/v1/apikey/test-prefix"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodDelete, &url.URL{Path: "http://example.com/api/v1/apikey/test-prefix"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		err := apiKeyResource.Delete(context.Background(), "test-prefix")
		assert.Error(t, err)
	})
}
