package e2e

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	hsClient "github.com/hibare/headscale-client-go/v1/client"
	"github.com/hibare/headscale-client-go/v1/preauthkeys"
	"github.com/hibare/headscale-client-go/v1/users"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/network"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	headscaleImage   = "headscale/headscale:v0.28.0"
	tailscaleImage   = "tailscale/tailscale:latest"
	containerTimeout = 5 * time.Minute
	testUser         = "testuser"
	nodeWaitTime     = 15 * time.Second
)

type E2ESuite struct {
	suite.Suite
	headscaleContainer testcontainers.Container
	tsNodeContainers   []testcontainers.Container
	client             hsClient.ClientInterface
	apiKey             string
	serverURL          string
	network            *testcontainers.DockerNetwork
	testUser           users.User
}

func TestE2E(t *testing.T) {
	suite.Run(t, new(E2ESuite))
}

func (s *E2ESuite) SetupSuite() {
	ctx, cancel := context.WithTimeout(s.T().Context(), containerTimeout)
	defer cancel()

	var err error

	s.network, err = s.createNetwork(ctx)
	s.Require().NoError(err, "Failed to create network")

	s.headscaleContainer, err = s.startHeadscale(ctx)
	s.Require().NoError(err, "Failed to start Headscale container")

	s.serverURL, err = s.getServerURL(ctx)
	s.Require().NoError(err, "Failed to get server URL")

	s.apiKey, err = s.createAPIKey(ctx)
	s.Require().NoError(err, "Failed to create API key")

	s.client, err = hsClient.NewClient(s.serverURL, s.apiKey, hsClient.ClientOptions{})
	s.Require().NoError(err, "Failed to create client")

	s.testUser, err = s.createTestUser(ctx)
	s.Require().NoError(err, "Failed to create test user")

	err = s.setInitialPolicy(ctx)
	s.Require().NoError(err, "Failed to set initial policy")

	preAuthKey, err := s.createPreAuthKey(ctx)
	s.Require().NoError(err, "Failed to create pre-auth key")

	tsNode1, err := s.startTailscaleNode(ctx, "test-node-1", preAuthKey)
	s.Require().NoError(err, "Failed to start Tailscale node 1")

	tsNode2, err := s.startTailscaleNode(ctx, "test-node-2", preAuthKey)
	s.Require().NoError(err, "Failed to start Tailscale node 2")

	s.tsNodeContainers = []testcontainers.Container{tsNode1, tsNode2}

	time.Sleep(nodeWaitTime)
}

func (s *E2ESuite) TearDownSuite() {
	ctx := s.T().Context()

	for _, c := range s.tsNodeContainers {
		if c != nil {
			_ = c.Terminate(ctx)
		}
	}

	if s.headscaleContainer != nil {
		_ = s.headscaleContainer.Terminate(ctx)
	}

	if s.network != nil {
		_ = s.network.Remove(ctx)
	}
}

func (s *E2ESuite) createNetwork(ctx context.Context) (*testcontainers.DockerNetwork, error) {
	return network.New(ctx, network.WithDriver("bridge"))
}

func (s *E2ESuite) startHeadscale(ctx context.Context) (testcontainers.Container, error) {
	configPath, err := filepath.Abs("config")
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path for config: %w", err)
	}

	req := testcontainers.ContainerRequest{
		Image:        headscaleImage,
		Cmd:          []string{"serve"},
		ExposedPorts: []string{"8080/tcp", "50443/tcp"},
		Networks:     []string{s.network.Name},
		NetworkAliases: map[string][]string{
			s.network.Name: {"headscale"},
		},
		WaitingFor: wait.ForHTTP("/health").WithPort("8080/tcp").WithStartupTimeout(containerTimeout),
	}

	if _, statErr := os.Stat(configPath); statErr == nil {
		req.Files = []testcontainers.ContainerFile{
			{
				HostFilePath:      filepath.Join(configPath, "config.yaml"),
				ContainerFilePath: "/etc/headscale/config.yaml",
			},
		}
	}

	return testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}

func (s *E2ESuite) getServerURL(ctx context.Context) (string, error) {
	port, err := s.headscaleContainer.MappedPort(ctx, "8080")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("http://localhost:%s", port.Port()), nil
}

func (s *E2ESuite) createAPIKey(ctx context.Context) (string, error) {
	exitCode, reader, err := s.headscaleContainer.Exec(ctx, []string{
		"headscale", "apikeys", "create", "--expiration", "999d",
	})
	if err != nil {
		return "", err
	}
	if exitCode != 0 {
		return "", fmt.Errorf("failed to create API key, exit code: %d", exitCode)
	}

	var buf bytes.Buffer
	_, err = buf.ReadFrom(reader)
	if err != nil {
		return "", err
	}

	output := buf.String()
	output = strings.TrimSpace(output)

	var apiKey string
	if strings.HasPrefix(output, "hskey-api-") {
		apiKey = output
	} else {
		idx := strings.Index(output, "hskey-api-")
		if idx >= 0 {
			apiKey = output[idx:]
		} else {
			apiKey = output
		}
	}
	apiKey = strings.TrimSpace(apiKey)

	s.T().Logf("API key: %s", apiKey)

	if apiKey == "" {
		return "", fmt.Errorf("API key is empty, raw output: %s", output)
	}

	return apiKey, nil
}

func (s *E2ESuite) createTestUser(ctx context.Context) (users.User, error) {
	user, err := s.client.Users().Create(ctx, users.CreateUserRequest{
		Name: testUser,
	})
	if err != nil {
		return users.User{}, err
	}
	return user.User, nil
}

func (s *E2ESuite) setInitialPolicy(ctx context.Context) error {
	acl := `{
		"tagOwners": {
			"tag:e2e-test": ["testuser@headscale.test"]
		},
		"acls": [
			{
				"action": "accept",
				"src": ["*"],
				"dst": ["*:*"]
			}
		]
	}`
	_, err := s.client.Policy().Update(ctx, acl)
	return err
}

func (s *E2ESuite) createPreAuthKey(ctx context.Context) (string, error) {
	key, err := s.client.PreAuthKeys().Create(ctx, preauthkeys.CreatePreAuthKeyRequest{
		User:       s.testUser.ID,
		Reusable:   true,
		Ephemeral:  false,
		Expiration: time.Now().Add(24 * time.Hour),
	})
	if err != nil {
		return "", err
	}
	return key.PreAuthKey.Key, nil
}

func (s *E2ESuite) startTailscaleNode(ctx context.Context, hostname, authKey string) (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Image:    tailscaleImage,
		Hostname: hostname,
		Env: map[string]string{
			"TS_STATE_DIR":  "/var/lib/tailscale",
			"TS_USERSPACE":  "false",
			"TS_AUTHKEY":    authKey,
			"TS_EXTRA_ARGS": fmt.Sprintf("--login-server=%s", "http://headscale:8080"),
		},
		CapAdd:   []string{"NET_ADMIN", "SYS_MODULE"},
		Networks: []string{s.network.Name},
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	return container, nil
}
