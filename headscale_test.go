package headscale

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MockLogger struct{}

func (m *MockLogger) Debug(ctx context.Context, msg string, keysAndValues ...interface{}) {}

func (m *MockLogger) Info(ctx context.Context, msg string, keysAndValues ...interface{}) {}

func (m *MockLogger) Warn(ctx context.Context, msg string, keysAndValues ...interface{}) {}

func (m *MockLogger) Error(ctx context.Context, msg string, keysAndValues ...interface{}) {}

func TestNewClient(t *testing.T) {
	baseURL := "http://example.com"
	apiKey := "test-api-key"
	opt := HeadscaleClientOptions{
		Http:      &http.Client{Timeout: 10 * time.Second},
		UserAgent: "custom-user-agent",
		Logger:    &MockLogger{},
	}

	client, err := NewClient(baseURL, apiKey, opt)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "custom-user-agent", client.(*Client).UserAgent)
	assert.Equal(t, apiKey, client.(*Client).APIKey)
	assert.Equal(t, 10*time.Second, client.(*Client).HTTP.Timeout)
}

func TestNewClientMissingAPIKey(t *testing.T) {
	baseURL := "http://example.com"
	opt := HeadscaleClientOptions{
		Http:      &http.Client{Timeout: 10 * time.Second},
		UserAgent: "custom-user-agent",
		Logger:    &MockLogger{},
	}

	client, err := NewClient(baseURL, "", opt)
	assert.Error(t, err)
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
	assert.NoError(t, err)
	assert.Equal(t, "application/json", req.Header.Get("Content-Type"))
	assert.Equal(t, "HeaderValue", req.Header.Get("Custom-Header"))
	assert.Equal(t, "Bearer test-api-key", req.Header.Get("Authorization"))
	assert.Equal(t, "headscale-client-go", req.Header.Get("User-Agent"))
	assert.Equal(t, "value", req.URL.Query().Get("param"))
}

func TestClient_do(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	assert.NoError(t, err)
	assert.Equal(t, []string{"change1", "change2"}, resp.Changes)
}
