package headscale

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type MockLogger struct{}

func (m *MockLogger) Debug(_ context.Context, _ string, _ ...interface{}) {}

func (m *MockLogger) Info(_ context.Context, _ string, _ ...interface{}) {}

func (m *MockLogger) Warn(_ context.Context, _ string, _ ...interface{}) {}

func (m *MockLogger) Error(_ context.Context, _ string, _ ...interface{}) {}

func TestNewClient(t *testing.T) {
	baseURL := "http://example.com"
	apiKey := "test-api-key"
	opt := ClientOptions{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		UserAgent:  "custom-user-agent",
		Logger:     &MockLogger{},
	}

	clientInterface, err := NewClient(baseURL, apiKey, opt)
	require.NoError(t, err)
	require.NotNil(t, clientInterface)

	client, ok := clientInterface.(*Client)
	require.True(t, ok, "expected client to be of type *Client")

	assert.Equal(t, "custom-user-agent", client.UserAgent)
	assert.Equal(t, apiKey, client.APIKey)
	assert.Equal(t, 10*time.Second, client.HTTP.Timeout)
}

func TestNewClientMissingAPIKey(t *testing.T) {
	baseURL := "http://example.com"
	opt := ClientOptions{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		UserAgent:  "custom-user-agent",
		Logger:     &MockLogger{},
	}

	client, err := NewClient(baseURL, "", opt)
	require.Error(t, err)
	assert.Nil(t, client)
}

func TestClient_buildURL(t *testing.T) {
	baseURL, _ := url.Parse("http://example.com")
	client := &Client{BaseURL: baseURL}

	url := client.buildURL("users", "123")
	expectedURL := "http://example.com/api/v1/users/123"
	assert.Equal(t, expectedURL, url.String())
}

func TestClient_buildRequest(t *testing.T) {
	baseURL, _ := url.Parse("http://example.com")
	client := &Client{
		BaseURL:   baseURL,
		UserAgent: DefaultUserAgent,
		APIKey:    "test-api-key",
	}

	ctx := context.Background()
	uri, _ := url.Parse("http://example.com/api/v1/test")
	opt := requestOptions{
		body:        map[string]string{"key": "value"},
		headers:     map[string]string{"Custom-Header": "HeaderValue"},
		contentType: "application/json",
		queryParams: map[string]interface{}{"param": "value"},
	}

	req, err := client.buildRequest(ctx, http.MethodPost, uri, opt)
	require.NoError(t, err)
	assert.Equal(t, "application/json", req.Header.Get("Content-Type"))
	assert.Equal(t, "HeaderValue", req.Header.Get("Custom-Header"))
	assert.Equal(t, "Bearer test-api-key", req.Header.Get("Authorization"))
	assert.Equal(t, "headscale-client-go", req.Header.Get("User-Agent"))
	assert.Equal(t, "value", req.URL.Query().Get("param"))
}

type ChangesResponse struct {
	Changes []string `json:"changes"`
}

func TestClient_do(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"changes": ["change1", "change2"]}`))
	}))
	defer server.Close()

	baseURL, _ := url.Parse(server.URL)
	client := &Client{
		BaseURL: baseURL,
		HTTP:    server.Client(),
		Logger:  &MockLogger{},
	}

	ctx := context.Background()
	uri, _ := url.Parse(server.URL)
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, uri.String(), nil)

	var resp ChangesResponse
	err := client.do(ctx, req, &resp)
	require.NoError(t, err)
	assert.Equal(t, []string{"change1", "change2"}, resp.Changes)
}
