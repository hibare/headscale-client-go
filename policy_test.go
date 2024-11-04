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

func TestPolicyResource_Get(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		policyResource := &PolicyResource{Client: client}
		client.(*MockClient).On("buildURL", "policy").Return(&url.URL{Path: "http://example.com/api/v1/policy"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodGet, &url.URL{Path: "http://example.com/api/v1/policy"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			resp := args.Get(1).(*Policy)
			resp.Policy = "test-policy"
			resp.UpdatedAt = "2024-01-01T00:00:00Z"
		})

		policy, err := policyResource.Get(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, "test-policy", policy.Policy)
		assert.Equal(t, "2024-01-01T00:00:00Z", policy.UpdatedAt)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		policyResource := &PolicyResource{Client: client}
		client.(*MockClient).On("buildURL", "policy").Return(&url.URL{Path: "http://example.com/api/v1/policy"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodGet, &url.URL{Path: "http://example.com/api/v1/policy"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		policy, err := policyResource.Get(context.Background())
		assert.Error(t, err)
		assert.Empty(t, policy.Policy)
	})
}

func TestPolicyResource_Put(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		policyResource := &PolicyResource{Client: client}
		client.(*MockClient).On("buildURL", "policy").Return(&url.URL{Path: "http://example.com/api/v1/policy"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPut, &url.URL{Path: "http://example.com/api/v1/policy"}, requestOptions{
			body: UpdatePolicyRequest{Policy: "new-policy"},
		}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil)

		err := policyResource.Put(context.Background(), "new-policy")
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		policyResource := &PolicyResource{Client: client}
		client.(*MockClient).On("buildURL", "policy").Return(&url.URL{Path: "http://example.com/api/v1/policy"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPut, &url.URL{Path: "http://example.com/api/v1/policy"}, requestOptions{
			body: UpdatePolicyRequest{Policy: "new-policy"},
		}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		err := policyResource.Put(context.Background(), "new-policy")
		assert.Error(t, err)
	})
}
