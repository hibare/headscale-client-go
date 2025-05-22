package requests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/hibare/headscale-client-go/logger"
	"github.com/hibare/headscale-client-go/versions"
	"github.com/stretchr/testify/require"
)

const (
	// TestAPIKey is a mock API key used for testing.
	TestAPIKey = "test-api-key"
)

// TestBuildURL verifies that BuildURL constructs the correct URL with path and query parameters.
func TestBuildURL(t *testing.T) {
	baseURL, err := url.Parse("http://example.com/")
	require.NoError(t, err)
	r := &Request{
		baseURL:    baseURL,
		apiKey:     TestAPIKey,
		apiVersion: versions.APIVersionV1,
		userAgent:  DefaultUserAgent,
		logger:     logger.NewDefaultLogger(logger.LevelError),
		httpClient: &http.Client{},
	}
	u := r.BuildURL("foo", "bar baz", 123)
	require.Equal(t, "http://example.com/api/v1/foo/bar%20baz/123", u.String())
}

// TestBuildRequest checks that BuildRequest sets headers, encodes body, and query params correctly.
func TestBuildRequest(t *testing.T) {
	baseURL, err := url.Parse("http://example.com/")
	require.NoError(t, err)
	r := &Request{
		baseURL:    baseURL,
		apiKey:     TestAPIKey,
		apiVersion: versions.APIVersionV1,
		userAgent:  DefaultUserAgent,
		logger:     logger.NewDefaultLogger(logger.LevelError),
		httpClient: &http.Client{},
	}
	uri := r.BuildURL("test")
	opt := RequestOptions{
		Body:        map[string]string{"hello": "world"},
		Headers:     map[string]string{"X-Test": "yes"},
		ContentType: "application/json",
		QueryParams: map[string]interface{}{"q": "search"},
	}
	req, err := r.BuildRequest(t.Context(), http.MethodPost, uri, opt)
	require.NoError(t, err)
	require.Equal(t, "application/json", req.Header.Get("Content-Type"))
	require.Equal(t, DefaultUserAgent, req.Header.Get("User-Agent"))
	require.Equal(t, "Bearer "+TestAPIKey, req.Header.Get("Authorization"))
	require.Equal(t, "yes", req.Header.Get("X-Test"))
	require.Contains(t, req.URL.RawQuery, "q=search")
	var body map[string]string
	err = json.NewDecoder(req.Body).Decode(&body)
	require.NoError(t, err)
	require.Equal(t, map[string]string{"hello": "world"}, body)
}

// TestDo_Success ensures Do executes the request and decodes the response on success.
func TestDo_Success(t *testing.T) {
	called := false
	h := func(w http.ResponseWriter, _ *http.Request) {
		called = true
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"foo":"bar"}`))
	}
	ts := httptest.NewServer(http.HandlerFunc(h))
	defer ts.Close()

	baseURL, _ := url.Parse(ts.URL + "/")
	r := &Request{
		baseURL:    baseURL,
		apiKey:     TestAPIKey,
		apiVersion: versions.APIVersionV1,
		userAgent:  DefaultUserAgent,
		logger:     logger.NewDefaultLogger(logger.LevelError),
		httpClient: ts.Client(),
	}

	uri := r.BuildURL("foo")
	req, err := r.BuildRequest(t.Context(), http.MethodGet, uri, RequestOptions{})
	require.NoError(t, err)
	var resp struct{ Foo string }
	err = r.Do(t.Context(), req, &resp)
	require.NoError(t, err)
	require.True(t, called)
	require.Equal(t, "bar", resp.Foo)
}

// TestDo_ErrorStatus checks that Do returns an error for non-2xx status codes.
func TestDo_ErrorStatus(t *testing.T) {
	h := func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}
	ts := httptest.NewServer(http.HandlerFunc(h))
	defer ts.Close()

	baseURL, _ := url.Parse(ts.URL + "/")
	r := &Request{
		baseURL:    baseURL,
		apiKey:     TestAPIKey,
		apiVersion: versions.APIVersionV1,
		userAgent:  DefaultUserAgent,
		logger:     logger.NewDefaultLogger(logger.LevelError),
		httpClient: ts.Client(),
	}

	uri := r.BuildURL("foo")
	req, err := r.BuildRequest(t.Context(), http.MethodGet, uri, RequestOptions{})
	require.NoError(t, err)
	err = r.Do(t.Context(), req, nil)
	require.Error(t, err)
}

// TestBuildRequest_InvalidURL checks that BuildRequest returns an error if the URL is nil.
func TestBuildRequest_InvalidURL(t *testing.T) {
	baseURL, err := url.Parse("http://example.com/")
	require.NoError(t, err)
	r := &Request{
		baseURL:    baseURL,
		apiKey:     TestAPIKey,
		apiVersion: versions.APIVersionV1,
		userAgent:  DefaultUserAgent,
		logger:     logger.NewDefaultLogger(logger.LevelError),
		httpClient: &http.Client{},
	}
	_, err = r.BuildRequest(t.Context(), http.MethodGet, nil, RequestOptions{})
	require.Error(t, err)
}

// TestBuildRequest_InvalidBody checks that BuildRequest returns an error for an unserializable body.
func TestBuildRequest_InvalidBody(t *testing.T) {
	baseURL, err := url.Parse("http://example.com/")
	require.NoError(t, err)
	r := &Request{
		baseURL:    baseURL,
		apiKey:     TestAPIKey,
		apiVersion: versions.APIVersionV1,
		userAgent:  DefaultUserAgent,
		logger:     logger.NewDefaultLogger(logger.LevelError),
		httpClient: &http.Client{},
	}
	uri := r.BuildURL("test")
	opt := RequestOptions{
		Body: make(chan int), // not JSON serializable
	}
	_, err = r.BuildRequest(t.Context(), http.MethodPost, uri, opt)
	require.Error(t, err)
}
