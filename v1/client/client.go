// Package client provides a Headscale client for the Headscale API.
package client

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/hibare/headscale-client-go/logger"
	"github.com/hibare/headscale-client-go/requests"
	"github.com/hibare/headscale-client-go/v1/apikeys"
	"github.com/hibare/headscale-client-go/v1/nodes"
	"github.com/hibare/headscale-client-go/v1/policy"
	"github.com/hibare/headscale-client-go/v1/preauthkeys"
	"github.com/hibare/headscale-client-go/v1/users"
	"github.com/hibare/headscale-client-go/versions"
)

var (
	// ErrAPIKeyRequired is returned when an API key is required but not provided.
	ErrAPIKeyRequired = errors.New("API key is required")
)

// ClientInterface defines the methods that a Headscale client must implement.
//
//nolint:revive // Name is intentionally to avoid confusions
type ClientInterface interface {
	APIKeys() apikeys.APIKeyResourceInterface
	Nodes() nodes.NodeResourceInterface
	Policy() policy.PolicyResourceInterface
	Users() users.UserResourceInterface
	PreAuthKeys() preauthkeys.PreAuthKeyResourceInterface
}

// Client is a struct that implements the HeadscaleClientInterface.
type Client struct {
	apiKeys     apikeys.APIKeyResourceInterface
	nodes       nodes.NodeResourceInterface
	policy      policy.PolicyResourceInterface
	users       users.UserResourceInterface
	preAuthKeys preauthkeys.PreAuthKeyResourceInterface
}

// APIKeys returns the APIKeyResource for managing API keys.
func (c *Client) APIKeys() apikeys.APIKeyResourceInterface {
	return c.apiKeys
}

// Nodes returns the NodeResource for managing nodes.
func (c *Client) Nodes() nodes.NodeResourceInterface {
	return c.nodes
}

// Policy returns the PolicyResource for managing policies.
func (c *Client) Policy() policy.PolicyResourceInterface {
	return c.policy
}

// Users returns the UserResource for managing users.
func (c *Client) Users() users.UserResourceInterface {
	return c.users
}

// PreAuthKeys returns the PreAuthKeyResource for managing pre-auth keys.
func (c *Client) PreAuthKeys() preauthkeys.PreAuthKeyResourceInterface {
	return c.preAuthKeys
}

// ClientOptions contains options for the Headscale client.
//
//nolint:revive // Name is intentionally to avoid confusions
type ClientOptions struct {
	HTTPClient *http.Client
	UserAgent  *string
	Logger     logger.Logger
}

// NewClient creates a new Headscale client with the specified base URL and API key.
func NewClient(baseURL, apiKey string, opt ClientOptions) (ClientInterface, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	if apiKey == "" {
		return nil, ErrAPIKeyRequired
	}

	// Set default values if not provided
	if opt.HTTPClient == nil {
		opt.HTTPClient = &http.Client{Timeout: requests.DefaultHTTPClientTimeout}
	} else if opt.HTTPClient.Timeout == 0 {
		// Add a default timeout if not provided
		opt.HTTPClient.Timeout = requests.DefaultHTTPClientTimeout
	}

	if opt.UserAgent == nil {
		userAgent := requests.DefaultUserAgent
		opt.UserAgent = &userAgent
	}

	if opt.Logger == nil {
		opt.Logger = logger.NewDefaultLogger(logger.LevelInfo)
	}

	// Create a new request with the given base URL, API key, and options
	request := requests.NewRequest(u, apiKey, versions.APIVersionV1, requests.RequestConfig{
		UserAgent:  opt.UserAgent,
		Logger:     opt.Logger,
		HTTPClient: opt.HTTPClient,
	})

	c := &Client{
		apiKeys:     apikeys.NewAPIKeyResource(request),
		nodes:       nodes.NewNodeResource(request),
		policy:      policy.NewPolicyResource(request),
		users:       users.NewUserResource(request),
		preAuthKeys: preauthkeys.NewPreAuthKeyResource(request),
	}

	return c, nil
}
