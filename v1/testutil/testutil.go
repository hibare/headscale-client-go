package testutil

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

const (
	// responseArgIndex is the index of the response argument in Do(ctx, req, v).
	responseArgIndex = 2
)

type TestFixture[R any] struct {
	Endpoint    any
	Method      string
	SuccessResp R
	BuildErr    error
	DoErr       error
}

func getEndpointArgs(endpoint any) []any {
	if endpoint == nil {
		return nil
	}
	switch v := endpoint.(type) {
	case []any:
		return v
	case []string:
		args := make([]any, len(v))
		for i, s := range v {
			args[i] = s
		}
		return args
	default:
		return []any{v}
	}
}

func RunResourceTest[R any](t *testing.T, fix TestFixture[R], action func(ctx context.Context, mockReq *requests.MockRequest) (R, error)) {
	t.Run("success", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		ctx := t.Context()
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		urlArgs := getEndpointArgs(fix.Endpoint)
		mockReq.On("BuildURL", urlArgs...).Return(fakeURL)
		mockReq.On("BuildRequest", ctx, fix.Method, fakeURL, mock.Anything).Return(fakeReq, nil)
		mockReq.On("Do", ctx, fakeReq, mock.Anything).Run(func(args mock.Arguments) {
			if len(args) > responseArgIndex && args.Get(responseArgIndex) != nil {
				if resp, ok := args.Get(responseArgIndex).(*R); ok {
					*resp = fix.SuccessResp
				}
			}
		}).Return(nil)

		resp, err := action(ctx, mockReq)
		require.NoError(t, err)
		assert.Equal(t, fix.SuccessResp, resp)
		mockReq.AssertExpectations(t)
	})

	t.Run("build request error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		ctx := t.Context()
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		urlArgs := getEndpointArgs(fix.Endpoint)
		mockReq.On("BuildURL", urlArgs...).Return(fakeURL)

		buildErr := fix.BuildErr
		if buildErr == nil {
			buildErr = errors.New("build error")
		}
		mockReq.On("BuildRequest", ctx, fix.Method, fakeURL, mock.Anything).Return(fakeReq, buildErr)

		resp, err := action(ctx, mockReq)
		require.Error(t, err)
		var zero R
		assert.Equal(t, zero, resp)
		mockReq.AssertExpectations(t)
	})

	t.Run("do error", func(t *testing.T) {
		mockReq := new(requests.MockRequest)
		ctx := t.Context()
		fakeURL := &url.URL{Scheme: "http", Host: "example.com"}
		fakeReq := &http.Request{}

		urlArgs := getEndpointArgs(fix.Endpoint)
		mockReq.On("BuildURL", urlArgs...).Return(fakeURL)
		mockReq.On("BuildRequest", ctx, fix.Method, fakeURL, mock.Anything).Return(fakeReq, nil)

		doErr := fix.DoErr
		if doErr == nil {
			doErr = errors.New("do error")
		}
		mockReq.On("Do", ctx, fakeReq, mock.Anything).Return(doErr)

		resp, err := action(ctx, mockReq)
		require.Error(t, err)
		var zero R
		assert.Equal(t, zero, resp)
		mockReq.AssertExpectations(t)
	})
}
