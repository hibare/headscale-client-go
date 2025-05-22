package client

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/hibare/headscale-client-go/logger"
	"github.com/hibare/headscale-client-go/requests"
	"github.com/hibare/headscale-client-go/v1/apikeys"
	"github.com/hibare/headscale-client-go/v1/nodes"
	"github.com/hibare/headscale-client-go/v1/policy"
	"github.com/hibare/headscale-client-go/v1/preauthkeys"
	"github.com/hibare/headscale-client-go/v1/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNewClient_InvalidURL(t *testing.T) {
	_, err := NewClient("://invalid-url", "key", ClientOptions{})
	assert.Error(t, err)
}

func TestNewClient_EmptyAPIKey(t *testing.T) {
	_, err := NewClient("http://localhost", "", ClientOptions{})
	assert.ErrorIs(t, err, ErrAPIKeyRequired)
}

func TestNewClient_Defaults(t *testing.T) {
	client, err := NewClient("http://localhost", "key", ClientOptions{})
	require.NoError(t, err)
	require.NotNil(t, client)
}

func TestNewClient_CustomOptions(t *testing.T) {
	ua := "custom-agent"
	opt := ClientOptions{
		HTTPClient: &http.Client{},
		UserAgent:  &ua,
		Logger:     logger.NewDefaultLogger(logger.LevelDebug),
	}
	client, err := NewClient("http://localhost", "key", opt)
	require.NoError(t, err)
	require.NotNil(t, client)
}

func TestNewClient_Alias(t *testing.T) {
	client, err := NewClient("http://localhost", "key", ClientOptions{})
	require.NoError(t, err)
	require.NotNil(t, client)
}

func TestClient_ResourceAccessors(t *testing.T) {
	// Use mocks for the resource interfaces
	mockReq := new(requests.MockRequest)
	requestURL, _ := url.Parse("http://localhost")
	mockReq.On("BuildURL", mock.Anything).Return(requestURL)
	mockReq.On("BuildRequest", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&http.Request{}, nil)
	mockReq.On("Do", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	c := &Client{
		apiKeys:     &apikeys.APIKeyResource{R: mockReq},
		nodes:       &nodes.NodeResource{R: mockReq},
		policy:      &policy.PolicyResource{R: mockReq},
		users:       &users.UserResource{R: mockReq},
		preAuthKeys: &preauthkeys.PreAuthKeyResource{R: mockReq},
	}
	assert.NotNil(t, c.APIKeys())
	assert.NotNil(t, c.Nodes())
	assert.NotNil(t, c.Policy())
	assert.NotNil(t, c.Users())
	assert.NotNil(t, c.PreAuthKeys())
}

func TestClientOptions_ZeroValue(t *testing.T) {
	var opt ClientOptions
	assert.Nil(t, opt.HTTPClient)
	assert.Nil(t, opt.UserAgent)
	assert.Nil(t, opt.Logger)
}
