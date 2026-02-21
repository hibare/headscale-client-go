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

	keysBefore, err := s.client.APIKeys().List(ctx)
	s.Require().NoError(err)

	resp, err := s.client.APIKeys().Create(ctx, apikeys.CreateAPIKeyRequest{
		Expiration: time.Now().Add(24 * time.Hour),
	})
	s.Require().NoError(err)
	s.NotEmpty(resp.APIKey, "Expected API key to be created")

	keysAfter, err := s.client.APIKeys().List(ctx)
	s.Require().NoError(err)
	s.Greater(len(keysAfter.APIKeys), len(keysBefore.APIKeys), "Expected one more API key after creation")

	newKey := s.findNewAPIKey(keysBefore.APIKeys, keysAfter.APIKeys)
	s.Require().NotNil(newKey, "Expected to find newly created key")

	err = s.client.APIKeys().Delete(ctx, newKey.Prefix)
	s.Require().NoError(err)

	keysFinal, err := s.client.APIKeys().List(ctx)
	s.Require().NoError(err)
	s.Len(keysFinal.APIKeys, len(keysBefore.APIKeys), "Expected same number of keys as before creation")
}

//nolint:dupl // Test code is intentionally similar for consistency
func (s *E2ESuite) TestAPIKeys_Expire() {
	ctx := s.T().Context()

	keysBefore, err := s.client.APIKeys().List(ctx)
	s.Require().NoError(err)

	resp, err := s.client.APIKeys().Create(ctx, apikeys.CreateAPIKeyRequest{
		Expiration: time.Now().Add(24 * time.Hour),
	})
	s.Require().NoError(err)
	s.NotEmpty(resp.APIKey, "Expected API key to be created")

	keysAfter, err := s.client.APIKeys().List(ctx)
	s.Require().NoError(err)

	newKey := s.findNewAPIKey(keysBefore.APIKeys, keysAfter.APIKeys)
	s.Require().NotNil(newKey, "Expected to find newly created key")
	s.False(newKey.Expiration.Before(time.Now()), "Expected key to not be expired")

	err = s.client.APIKeys().Expire(ctx, newKey.Prefix)
	s.Require().NoError(err)

	// Verify the key is expired
	keysExpired, err := s.client.APIKeys().List(ctx)
	s.Require().NoError(err)
	expiredKey := s.findAPIKeyByID(keysExpired.APIKeys, newKey.ID)
	s.Require().NotNil(expiredKey, "Expected to find the expired key")
	s.True(expiredKey.Expiration.Before(time.Now()), "Expected key to be expired")

	err = s.client.APIKeys().Delete(ctx, newKey.Prefix)
	s.Require().NoError(err)
}

//nolint:dupl // Test code is intentionally similar for consistency
func (s *E2ESuite) TestAPIKeys_ExpireByID() {
	ctx := s.T().Context()

	keysBefore, err := s.client.APIKeys().List(ctx)
	s.Require().NoError(err)

	resp, err := s.client.APIKeys().Create(ctx, apikeys.CreateAPIKeyRequest{
		Expiration: time.Now().Add(24 * time.Hour),
	})
	s.Require().NoError(err)
	s.NotEmpty(resp.APIKey, "Expected API key to be created")

	keysAfter, err := s.client.APIKeys().List(ctx)
	s.Require().NoError(err)

	newKey := s.findNewAPIKey(keysBefore.APIKeys, keysAfter.APIKeys)
	s.Require().NotNil(newKey, "Expected to find newly created key")
	s.False(newKey.Expiration.Before(time.Now()), "Expected key to not be expired")

	err = s.client.APIKeys().ExpireByID(ctx, newKey.ID)
	s.Require().NoError(err)

	// Verify the key is expired
	keysExpired, err := s.client.APIKeys().List(ctx)
	s.Require().NoError(err)
	expiredKey := s.findAPIKeyByID(keysExpired.APIKeys, newKey.ID)
	s.Require().NotNil(expiredKey, "Expected to find the expired key")
	s.True(expiredKey.Expiration.Before(time.Now()), "Expected key to be expired")

	err = s.client.APIKeys().Delete(ctx, newKey.Prefix)
	s.Require().NoError(err)
}

// findNewAPIKey finds the key that exists in after but not in before by comparing IDs.
func (s *E2ESuite) findNewAPIKey(before, after []apikeys.APIKey) *apikeys.APIKey {
	beforeIDs := make(map[string]bool)
	for _, k := range before {
		beforeIDs[k.ID] = true
	}

	for i := range after {
		if !beforeIDs[after[i].ID] {
			return &after[i]
		}
	}
	return nil
}

// findAPIKeyByID finds a key by ID in the given slice.
func (s *E2ESuite) findAPIKeyByID(keys []apikeys.APIKey, id string) *apikeys.APIKey {
	for i := range keys {
		if keys[i].ID == id {
			return &keys[i]
		}
	}
	return nil
}
