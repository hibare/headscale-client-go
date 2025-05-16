package headscale

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	// DefaultUserAgent is the default user agent for the Headscale client.
	DefaultUserAgent = "headscale-client-go"

	// DefaultHTTPClientTimeout is the default timeout for HTTP requests.
	DefaultHTTPClientTimeout = time.Minute

	// basePath is the base path for the Headscale API.
	basePath = "/api/v1"
)

// ClientInterface defines the methods that a Headscale client must implement.
type ClientInterface interface {
	APIKeys() *APIKeyResource
	Nodes() *NodeResource
	Policy() *PolicyResource
	Routes() *RoutesResource
	Users() *UserResource
	PreAuthKeys() *PreAuthKeyResource
	buildURL(pathParts ...any) *url.URL
	buildRequest(ctx context.Context, method string, uri *url.URL, opt requestOptions) (*http.Request, error)
	do(ctx context.Context, req *http.Request, v interface{}) error
}

// Client is a struct that implements the HeadscaleClientInterface.
type Client struct {
	BaseURL   *url.URL
	UserAgent string
	APIKey    string
	HTTP      *http.Client
	Logger    Logger

	// Specific resources
	apiKeys     *APIKeyResource
	nodes       *NodeResource
	policy      *PolicyResource
	routes      *RoutesResource
	users       *UserResource
	preAuthKeys *PreAuthKeyResource
}

// APIKeys returns the APIKeyResource for managing API keys.
func (c *Client) APIKeys() *APIKeyResource {
	return c.apiKeys
}

// Nodes returns the NodeResource for managing nodes.
func (c *Client) Nodes() *NodeResource {
	return c.nodes
}

// Policy returns the PolicyResource for managing policies.
func (c *Client) Policy() *PolicyResource {
	return c.policy
}

// Users returns the UserResource for managing users.
func (c *Client) Users() *UserResource {
	return c.users
}

// Routes returns the RoutesResource for managing routes.
func (c *Client) Routes() *RoutesResource {
	return c.routes
}

// PreAuthKeys returns the PreAuthKeyResource for managing pre-auth keys.
func (c *Client) PreAuthKeys() *PreAuthKeyResource {
	return c.preAuthKeys
}

func (c *Client) buildURL(pathParts ...any) *url.URL {
	parts := make([]string, 1, len(pathParts)+1)
	parts[0] = basePath
	for _, p := range pathParts {
		parts = append(parts, url.PathEscape(fmt.Sprint(p)))
	}

	return c.BaseURL.JoinPath(parts...)
}

type requestOptions struct {
	body        interface{}
	headers     map[string]string
	contentType string
	queryParams map[string]interface{}
}

func (c *Client) buildRequest(ctx context.Context, method string, uri *url.URL, opt requestOptions) (*http.Request, error) {
	var bodyBytes []byte
	if opt.body != nil {
		var err error

		switch body := opt.body.(type) {
		case []byte:
			bodyBytes = body
		case string:
			bodyBytes = []byte(body)
		default:
			bodyBytes, err = json.Marshal(opt.body)
			if err != nil {
				return nil, err
			}
		}
	}

	query := uri.Query()
	for k, v := range opt.queryParams {
		query.Add(k, fmt.Sprint(v))
	}
	uri.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, method, uri.String(), bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))

	if opt.contentType != "" {
		req.Header.Set("Content-Type", opt.contentType)
	}

	for k, v := range opt.headers {
		req.Header.Set(k, v)
	}

	return req, nil
}

func (c *Client) do(ctx context.Context, req *http.Request, v interface{}) error {
	c.Logger.Debug(ctx, "Request: ", "method", req.Method, "url", req.URL.String())
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return err
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
		if err != nil {
			return err
		}
	}

	return nil
}

// ClientOptions contains options for the Headscale client.
type ClientOptions struct {
	HTTPClient *http.Client
	UserAgent  string
	Logger     Logger
}

// NewClient creates a new Headscale client with the specified base URL and API key.
func NewClient(baseURL, apiKey string, opt ClientOptions) (ClientInterface, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	if apiKey == "" {
		return nil, errors.New("API key is required")
	}

	c := &Client{
		BaseURL:   u,
		UserAgent: DefaultUserAgent,
		APIKey:    apiKey,
		HTTP:      &http.Client{Timeout: DefaultHTTPClientTimeout},
		Logger:    NewDefaultLogger(LevelDebug),
	}

	if opt.HTTPClient != nil {
		c.HTTP = opt.HTTPClient
	}

	if opt.UserAgent != "" {
		c.UserAgent = opt.UserAgent
	}

	if opt.Logger != nil {
		c.Logger = opt.Logger
	}

	c.apiKeys = &APIKeyResource{c}
	c.nodes = &NodeResource{c}
	c.policy = &PolicyResource{c}
	c.routes = &RoutesResource{c}
	c.users = &UserResource{c}
	c.preAuthKeys = &PreAuthKeyResource{c}

	return c, nil
}
