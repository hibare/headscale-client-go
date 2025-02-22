package headscale

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	TestAPIKey              = "test-api-key"
	ExpectedTestBearerToken = "Bearer test-api-key"
)

func setupTestServer(t *testing.T, handler http.HandlerFunc) HeadscaleClientInterface {
	server := httptest.NewServer(handler)
	t.Cleanup(func() { server.Close() })

	client, err := NewClient(server.URL, TestAPIKey, HeadscaleClientOptions{})
	require.NoError(t, err)

	return client
}
