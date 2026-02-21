package e2e

import (
	"time"

	"github.com/hibare/headscale-client-go/v1/apikeys"
)

func (s *E2ESuite) TestAPIKeys_List() {
	keys, err := s.client.APIKeys().List(s.T().Context())
	s.Require().NoError(err)
	s.NotEmpty(keys.APIKeys, "Expected at least one API key")
}

func (s *E2ESuite) TestAPIKeys_CreateAndDelete() {
	ctx := s.T().Context()

	resp, err := s.client.APIKeys().Create(ctx, apikeys.CreateAPIKeyRequest{
		Expiration: time.Now().Add(24 * time.Hour),
	})
	s.Require().NoError(err)
	s.NotEmpty(resp.APIKey, "Expected API key to be created")

	keysBefore, err := s.client.APIKeys().List(ctx)
	s.Require().NoError(err)

	err = s.client.APIKeys().Delete(ctx, keysBefore.APIKeys[len(keysBefore.APIKeys)-1].Prefix)
	s.Require().NoError(err)

	keysAfter, err := s.client.APIKeys().List(ctx)
	s.Require().NoError(err)
	s.Less(len(keysAfter.APIKeys), len(keysBefore.APIKeys), "Expected one less API key after deletion")
}

//nolint:dupl // Test code is intentionally similar for consistency
func (s *E2ESuite) TestAPIKeys_Expire() {
	ctx := s.T().Context()

	resp, err := s.client.APIKeys().Create(ctx, apikeys.CreateAPIKeyRequest{
		Expiration: time.Now().Add(24 * time.Hour),
	})
	s.Require().NoError(err)
	s.NotEmpty(resp.APIKey, "Expected API key to be created")

	keys, err := s.client.APIKeys().List(ctx)
	s.Require().NoError(err)

	lastKey := keys.APIKeys[len(keys.APIKeys)-1]
	s.False(lastKey.Expiration.Before(time.Now()), "Expected key to not be expired")

	err = s.client.APIKeys().Expire(ctx, lastKey.Prefix)
	s.Require().NoError(err)

	err = s.client.APIKeys().Delete(ctx, lastKey.Prefix)
	s.Require().NoError(err)
}

//nolint:dupl // Test code is intentionally similar for consistency
func (s *E2ESuite) TestAPIKeys_ExpireByID() {
	ctx := s.T().Context()

	resp, err := s.client.APIKeys().Create(ctx, apikeys.CreateAPIKeyRequest{
		Expiration: time.Now().Add(24 * time.Hour),
	})
	s.Require().NoError(err)
	s.NotEmpty(resp.APIKey, "Expected API key to be created")

	keys, err := s.client.APIKeys().List(ctx)
	s.Require().NoError(err)

	lastKey := keys.APIKeys[len(keys.APIKeys)-1]
	s.False(lastKey.Expiration.Before(time.Now()), "Expected key to not be expired")

	err = s.client.APIKeys().ExpireByID(ctx, lastKey.ID)
	s.Require().NoError(err)

	err = s.client.APIKeys().Delete(ctx, lastKey.Prefix)
	s.Require().NoError(err)
}
