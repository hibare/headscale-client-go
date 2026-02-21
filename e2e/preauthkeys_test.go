package e2e

import (
	"time"

	"github.com/hibare/headscale-client-go/v1/preauthkeys"
)

func (s *E2ESuite) TestPreAuthKeys_List() {
	keys, err := s.client.PreAuthKeys().List(s.T().Context())
	s.Require().NoError(err)
	s.NotEmpty(keys.PreAuthKeys, "Expected at least one pre-auth key")
}

func (s *E2ESuite) TestPreAuthKeys_CreateAndDelete() {
	ctx := s.T().Context()

	keysBefore, err := s.client.PreAuthKeys().List(ctx)
	s.Require().NoError(err)

	key, err := s.client.PreAuthKeys().Create(ctx, preauthkeys.CreatePreAuthKeyRequest{
		User:       s.testUser.ID,
		Reusable:   true,
		Ephemeral:  false,
		Expiration: time.Now().Add(24 * time.Hour),
	})
	s.Require().NoError(err)
	s.NotEmpty(key.PreAuthKey.Key, "Expected pre-auth key to be created")

	keysAfter, err := s.client.PreAuthKeys().List(ctx)
	s.Require().NoError(err)
	s.Greater(len(keysAfter.PreAuthKeys), len(keysBefore.PreAuthKeys), "Expected one more pre-auth key after creation")

	err = s.client.PreAuthKeys().Delete(ctx, key.PreAuthKey.ID)
	s.Require().NoError(err)

	keysFinal, err := s.client.PreAuthKeys().List(ctx)
	s.Require().NoError(err)
	s.Len(keysFinal.PreAuthKeys, len(keysBefore.PreAuthKeys), "Expected same number of keys as before creation")
}

func (s *E2ESuite) TestPreAuthKeys_Expire() {
	ctx := s.T().Context()

	key, err := s.client.PreAuthKeys().Create(ctx, preauthkeys.CreatePreAuthKeyRequest{
		User:       s.testUser.ID,
		Reusable:   true,
		Ephemeral:  false,
		Expiration: time.Now().Add(24 * time.Hour),
	})
	s.Require().NoError(err)
	s.NotEmpty(key.PreAuthKey.Key, "Expected pre-auth key to be created")
	s.False(key.PreAuthKey.Expiration.Before(time.Now()), "Expected key to not be expired initially")

	err = s.client.PreAuthKeys().Expire(ctx, key.PreAuthKey.ID)
	s.Require().NoError(err)

	// Verify the key is expired
	keys, err := s.client.PreAuthKeys().List(ctx)
	s.Require().NoError(err)
	expiredKey := s.findPreAuthKeyByID(keys.PreAuthKeys, key.PreAuthKey.ID)
	s.Require().NotNil(expiredKey, "Expected to find the expired key")
	s.True(expiredKey.Expiration.Before(time.Now()), "Expected key to be expired after expire call")

	err = s.client.PreAuthKeys().Delete(ctx, key.PreAuthKey.ID)
	s.Require().NoError(err)
}

func (s *E2ESuite) TestPreAuthKeys_Reusable() {
	ctx := s.T().Context()

	key, err := s.client.PreAuthKeys().Create(ctx, preauthkeys.CreatePreAuthKeyRequest{
		User:       s.testUser.ID,
		Reusable:   true,
		Ephemeral:  false,
		Expiration: time.Now().Add(24 * time.Hour),
	})
	s.Require().NoError(err)
	s.True(key.PreAuthKey.Reusable, "Expected pre-auth key to be reusable")

	err = s.client.PreAuthKeys().Delete(ctx, key.PreAuthKey.ID)
	s.Require().NoError(err)
}

func (s *E2ESuite) TestPreAuthKeys_Ephemeral() {
	ctx := s.T().Context()

	key, err := s.client.PreAuthKeys().Create(ctx, preauthkeys.CreatePreAuthKeyRequest{
		User:       s.testUser.ID,
		Reusable:   false,
		Ephemeral:  true,
		Expiration: time.Now().Add(24 * time.Hour),
	})
	s.Require().NoError(err)
	s.True(key.PreAuthKey.Ephemeral, "Expected pre-auth key to be ephemeral")

	err = s.client.PreAuthKeys().Delete(ctx, key.PreAuthKey.ID)
	s.Require().NoError(err)
}

func (s *E2ESuite) TestPreAuthKeys_WithACLTags() {
	ctx := s.T().Context()

	key, err := s.client.PreAuthKeys().Create(ctx, preauthkeys.CreatePreAuthKeyRequest{
		User:       s.testUser.ID,
		Reusable:   true,
		Ephemeral:  false,
		Expiration: time.Now().Add(24 * time.Hour),
		ACLTags:    []string{"tag:test"},
	})
	s.Require().NoError(err)
	s.Require().Contains(key.PreAuthKey.ACLTags, "tag:test", "Expected pre-auth key to have the specified ACL tag")

	err = s.client.PreAuthKeys().Delete(ctx, key.PreAuthKey.ID)
	s.Require().NoError(err)
}

// findPreAuthKeyByID finds a pre-auth key by ID in the given slice.
func (s *E2ESuite) findPreAuthKeyByID(keys []preauthkeys.PreAuthKey, id string) *preauthkeys.PreAuthKey {
	for i := range keys {
		if keys[i].ID == id {
			return &keys[i]
		}
	}
	return nil
}
