package policy

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/hibare/headscale-client-go/requests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestPolicyResource_Get(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		p := &PolicyResource{R: mockReq}
		ctx := t.Context()
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}
		fakeResp := Policy{Policy: "test-policy", UpdatedAt: "2024-01-01T00:00:00Z"}

		mockReq.On("BuildURL", "policy").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodGet, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*policy.Policy")).Run(func(args mock.Arguments) {
			resp := args.Get(2).(*Policy) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
			*resp = fakeResp
		}).Return(nil)

		resp, err := p.Get(ctx)
		require.NoError(t, err)
		assert.Equal(t, fakeResp, resp)
		mockReq.AssertExpectations(t)
	})

	t.Run("build request error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		p := &PolicyResource{R: mockReq}
		ctx := t.Context()
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "policy").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodGet, fakeURL, mock.Anything).Return(fakeReq, errors.New("build error"))

		resp, err := p.Get(ctx)
		require.Error(t, err)
		assert.Empty(t, resp.Policy)
		mockReq.AssertExpectations(t)
	})

	t.Run("do error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		p := &PolicyResource{R: mockReq}
		ctx := t.Context()
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "policy").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodGet, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*policy.Policy")).Return(errors.New("do error"))

		resp, err := p.Get(ctx)
		require.Error(t, err)
		assert.Empty(t, resp.Policy)
		mockReq.AssertExpectations(t)
	})
}

func TestPolicyResource_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		p := &PolicyResource{R: mockReq}
		ctx := t.Context()
		policyStr := "new-policy"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}
		fakeResp := UpdatePolicyResponse{
			Policy:    policyStr,
			UpdatedAt: "2024-01-01T00:00:00Z",
		}

		mockReq.On("BuildURL", "policy").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPut, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*policy.UpdatePolicyResponse")).Run(func(args mock.Arguments) {
			resp := args.Get(2).(*UpdatePolicyResponse) //nolint:errcheck // reason: type assertion on mock, error not possible/needed
			*resp = fakeResp
		}).Return(nil)

		resp, err := p.Update(ctx, policyStr)
		require.NoError(t, err)
		assert.Equal(t, fakeResp, resp)
		mockReq.AssertExpectations(t)
	})

	t.Run("build request error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		p := &PolicyResource{R: mockReq}
		ctx := t.Context()
		policyStr := "new-policy"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "policy").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPut, fakeURL, mock.Anything).Return(fakeReq, errors.New("build error"))

		resp, err := p.Update(ctx, policyStr)
		require.Error(t, err)
		assert.Empty(t, resp.Policy)
		mockReq.AssertExpectations(t)
	})

	t.Run("do error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		p := &PolicyResource{R: mockReq}
		ctx := t.Context()
		policyStr := "new-policy"
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		mockReq.On("BuildURL", "policy").Return(fakeURL)
		mockReq.On("BuildRequest", ctx, http.MethodPut, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.AnythingOfType("*policy.UpdatePolicyResponse")).Return(errors.New("do error"))

		resp, err := p.Update(ctx, policyStr)
		require.Error(t, err)
		assert.Empty(t, resp.Policy)
		mockReq.AssertExpectations(t)
	})
}
