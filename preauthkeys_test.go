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

func TestPreAuthKeyResource_List(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		preAuthKeyResource := &PreAuthKeyResource{Client: client}
		client.(*MockClient).On("buildURL", "preauthkey").Return(&url.URL{Path: "http://example.com/api/v1/preauthkey"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodGet, &url.URL{Path: "http://example.com/api/v1/preauthkey"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			resp := args.Get(1).(*PreAuthKeysResponse)
			resp.PreAuthKeys = []PreAuthKey{{ID: "1", Key: "test-key"}}
		})

		keys, err := preAuthKeyResource.List(context.Background())
		assert.NoError(t, err)
		assert.Len(t, keys.PreAuthKeys, 1)
		assert.Equal(t, "test-key", keys.PreAuthKeys[0].Key)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		preAuthKeyResource := &PreAuthKeyResource{Client: client}
		client.(*MockClient).On("buildURL", "preauthkey").Return(&url.URL{Path: "http://example.com/api/v1/preauthkey"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodGet, &url.URL{Path: "http://example.com/api/v1/preauthkey"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		keys, err := preAuthKeyResource.List(context.Background())
		assert.Error(t, err)
		assert.Empty(t, keys.PreAuthKeys)
	})
}

func TestPreAuthKeyResource_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		preAuthKeyResource := &PreAuthKeyResource{Client: client}
		expiration := time.Now().Add(24 * time.Hour)
		client.(*MockClient).On("buildURL", "preauthkey").Return(&url.URL{Path: "http://example.com/api/v1/preauthkey"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPost, &url.URL{Path: "http://example.com/api/v1/preauthkey"}, requestOptions{
			body: CreatePreAuthKeyRequest{
				User:       "test-user",
				Reusable:   true,
				Ephemeral:  false,
				Expiration: expiration,
				AclTags:    []string{"tag1", "tag2"},
			},
		}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			resp := args.Get(1).(*PreAuthKeyResponse)
			resp.PreAuthKey = []PreAuthKey{{ID: "1", Key: "test-key"}}
		})

		key, err := preAuthKeyResource.Create(context.Background(), "test-user", true, false, expiration, []string{"tag1", "tag2"})
		assert.NoError(t, err)
		assert.Len(t, key.PreAuthKey, 1)
		assert.Equal(t, "test-key", key.PreAuthKey[0].Key)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		preAuthKeyResource := &PreAuthKeyResource{Client: client}
		expiration := time.Now().Add(24 * time.Hour)
		client.(*MockClient).On("buildURL", "preauthkey").Return(&url.URL{Path: "http://example.com/api/v1/preauthkey"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPost, &url.URL{Path: "http://example.com/api/v1/preauthkey"}, requestOptions{
			body: CreatePreAuthKeyRequest{
				User:       "test-user",
				Reusable:   true,
				Ephemeral:  false,
				Expiration: expiration,
				AclTags:    []string{"tag1", "tag2"},
			},
		}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		key, err := preAuthKeyResource.Create(context.Background(), "test-user", true, false, expiration, []string{"tag1", "tag2"})
		assert.Error(t, err)
		assert.Empty(t, key.PreAuthKey)
	})
}

func TestPreAuthKeyResource_Expire(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		preAuthKeyResource := &PreAuthKeyResource{Client: client}
		client.(*MockClient).On("buildURL", "preauthkey", "expire").Return(&url.URL{Path: "http://example.com/api/v1/preauthkey/expire"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPost, &url.URL{Path: "http://example.com/api/v1/preauthkey/expire"}, requestOptions{
			body: ExpirePreAuthKeyRequest{
				User: "test-user",
				Key:  "test-key",
			},
		}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil)

		err := preAuthKeyResource.Expire(context.Background(), "test-user", "test-key")
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		preAuthKeyResource := &PreAuthKeyResource{Client: client}
		client.(*MockClient).On("buildURL", "preauthkey", "expire").Return(&url.URL{Path: "http://example.com/api/v1/preauthkey/expire"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPost, &url.URL{Path: "http://example.com/api/v1/preauthkey/expire"}, requestOptions{
			body: ExpirePreAuthKeyRequest{
				User: "test-user",
				Key:  "test-key",
			},
		}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		err := preAuthKeyResource.Expire(context.Background(), "test-user", "test-key")
		assert.Error(t, err)
	})
}
