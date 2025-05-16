package headscale

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	// TestAPIKey is a mock API key used for testing.
	TestAPIKey = "test-api-key"

	// ExpectedTestBearerToken is the expected Authorization header value for the test API key.
	// This is a mock value used only for testing purposes.
	// #nosec G101
	ExpectedTestBearerToken = "Bearer test-api-key"
)

func setupTestServer(t *testing.T, handler http.HandlerFunc) ClientInterface {
	server := httptest.NewServer(handler)
	t.Cleanup(func() { server.Close() })

	client, err := NewClient(server.URL, TestAPIKey, ClientOptions{})
	require.NoError(t, err)

	return client
}
