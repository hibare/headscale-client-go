package headscale

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const DefaultUserAgent = "headscale-client-go"
const DefaultHTTPClientTimeout = time.Minute
const basePath = "/api/v1"

type HeadscaleClientInterface interface {
	APIKeys() *APIKeyResource
	Nodes() *NodeResource
	Policy() *PolicyResource
	Routes() *RoutesResource
	Users() *UserResource
	buildURL(pathParts ...any) *url.URL
	buildRequest(ctx context.Context, method string, uri *url.URL, opt requestOptions) (*http.Request, error)
	do(ctx context.Context, req *http.Request, v interface{}) error
}

type Client struct {
	BaseURL   *url.URL
	UserAgent string
	APIKey    string
	HTTP      *http.Client
	Logger    Logger

	// Specific resources
	apiKeys *APIKeyResource
	nodes   *NodeResource
	policy  *PolicyResource
	routes  *RoutesResource
	users   *UserResource
}

func (c *Client) APIKeys() *APIKeyResource {
	return c.apiKeys
}

func (c *Client) Nodes() *NodeResource {
	return c.nodes
}

func (c *Client) Policy() *PolicyResource {
	return c.policy
}

func (c *Client) Users() *UserResource {
	return c.users
}

func (c *Client) Routes() *RoutesResource {
	return c.routes
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

	defer resp.Body.Close()

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

type ChangesResponse struct {
	Changes []string `json:"changes"`
}

type HeadscaleClientOptions struct {
	Http      *http.Client
	UserAgent string
	Logger    Logger
}

func NewClient(baseURL, apiKey string, opt HeadscaleClientOptions) (HeadscaleClientInterface, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	if apiKey == "" {
		return nil, fmt.Errorf("API key is required")
	}

	c := &Client{
		BaseURL:   u,
		UserAgent: DefaultUserAgent,
		APIKey:    apiKey,
		HTTP:      &http.Client{Timeout: DefaultHTTPClientTimeout},
		Logger:    NewDefaultLogger(LevelDebug),
	}

	if opt.Http != nil {
		c.HTTP = opt.Http
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

	return c, nil
}
