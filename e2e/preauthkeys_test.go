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

	key, err := s.client.PreAuthKeys().Create(ctx, preauthkeys.CreatePreAuthKeyRequest{
		User:       s.testUser.ID,
		Reusable:   true,
		Ephemeral:  false,
		Expiration: time.Now().Add(24 * time.Hour),
	})
	s.Require().NoError(err)
	s.NotEmpty(key.PreAuthKey.Key, "Expected pre-auth key to be created")

	keysBefore, err := s.client.PreAuthKeys().List(ctx)
	s.Require().NoError(err)

	err = s.client.PreAuthKeys().Delete(ctx, key.PreAuthKey.ID)
	s.Require().NoError(err)

	keysAfter, err := s.client.PreAuthKeys().List(ctx)
	s.Require().NoError(err)
	s.Less(len(keysAfter.PreAuthKeys), len(keysBefore.PreAuthKeys), "Expected one less pre-auth key after deletion")
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

	err = s.client.PreAuthKeys().Expire(ctx, key.PreAuthKey.ID)
	s.Require().NoError(err)

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
	s.NotEmpty(key.PreAuthKey.ACLTags, "Expected pre-auth key to have ACL tags")

	err = s.client.PreAuthKeys().Delete(ctx, key.PreAuthKey.ID)
	s.Require().NoError(err)
}
