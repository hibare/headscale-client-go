// Package requests provides a HTTP client for the Headscale API.
package requests

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/hibare/headscale-client-go/logger"
	"github.com/hibare/headscale-client-go/versions"
)

const (
	// DefaultUserAgent is the default user agent for the Headscale client.
	DefaultUserAgent = "headscale-client-go"

	// DefaultHTTPClientTimeout is the default timeout for HTTP requests.
	DefaultHTTPClientTimeout = time.Minute
)

var (
	// ErrURIRequired is returned when a URI is required but not provided.
	ErrURIRequired = errors.New("uri cannot be nil")
)

// RequestInterface defines the interface for building and executing HTTP requests.
type RequestInterface interface {
	BuildURL(pathParts ...any) *url.URL
	BuildRequest(ctx context.Context, method string, uri *url.URL, opt RequestOptions) (*http.Request, error)
	Do(ctx context.Context, req *http.Request, v interface{}) error
}

// Request represents an HTTP request builder and executor.
type Request struct {
	baseURL    *url.URL
	apiKey     string
	apiVersion versions.APIVersion
	userAgent  string
	logger     logger.Logger
	httpClient *http.Client
}

// BuildURL constructs a URL from the base URL, API version, and additional path parts.
func (r *Request) BuildURL(pathParts ...any) *url.URL {
	parts := make([]string, 1, len(pathParts)+1)
	parts[0] = r.apiVersion.GetBasePath()
	for _, p := range pathParts {
		parts = append(parts, url.PathEscape(fmt.Sprint(p)))
	}

	return r.baseURL.JoinPath(parts...)
}

// RequestOptions contains options for building an HTTP request.
type RequestOptions struct {
	Body        interface{}
	Headers     map[string]string
	ContentType string
	QueryParams map[string]interface{}
}

// BuildRequest creates an HTTP request with the specified method, URL, and options.
func (r *Request) BuildRequest(ctx context.Context, method string, uri *url.URL, opt RequestOptions) (*http.Request, error) {
	if uri == nil {
		return nil, ErrURIRequired
	}

	var bodyBytes []byte
	if opt.Body != nil {
		var err error

		switch body := opt.Body.(type) {
		case []byte:
			bodyBytes = body
		case string:
			bodyBytes = []byte(body)
		default:
			bodyBytes, err = json.Marshal(opt.Body)
			if err != nil {
				return nil, err
			}
		}
	}

	query := uri.Query()
	for k, v := range opt.QueryParams {
		query.Add(k, fmt.Sprint(v))
	}
	uri.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, method, uri.String(), bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", r.userAgent)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.apiKey))

	if opt.ContentType != "" {
		req.Header.Set("Content-Type", opt.ContentType)
	}

	for k, v := range opt.Headers {
		req.Header.Set(k, v)
	}

	return req, nil
}

// Do executes the HTTP request and decodes the response into v if provided.
func (r *Request) Do(ctx context.Context, req *http.Request, v interface{}) error {
	r.logger.Debug(ctx, "Request: ", "method", req.Method, "url", req.URL.String())
	resp, err := r.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		// Read response body for logging
		respBody, rErr := io.ReadAll(resp.Body)
		if rErr == nil {
			r.logger.Error(ctx, "Response: ", "status", resp.StatusCode, "body", string(respBody))
		}
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

// RequestConfig contains configuration options for creating a new Request.
type RequestConfig struct {
	UserAgent  *string
	Logger     logger.Logger
	HTTPClient *http.Client
}

// NewRequest creates a new Request instance with the given configuration.
func NewRequest(baseURL *url.URL, apiKey string, apiVersion versions.APIVersion, opt RequestConfig) RequestInterface {
	if opt.UserAgent == nil {
		userAgent := DefaultUserAgent
		opt.UserAgent = &userAgent
	}

	if opt.HTTPClient == nil {
		httpClient := &http.Client{
			Timeout: DefaultHTTPClientTimeout,
		}
		opt.HTTPClient = httpClient
	}

	return &Request{
		baseURL:    baseURL,
		apiKey:     apiKey,
		apiVersion: apiVersion,
		userAgent:  *opt.UserAgent,
		logger:     opt.Logger,
		httpClient: opt.HTTPClient,
	}
}
