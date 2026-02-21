package e2e

import (
	"github.com/hibare/headscale-client-go/v1/users"
)

func (s *E2ESuite) TestUsers_List() {
	userList, err := s.client.Users().List(s.T().Context(), users.UserListFilter{})
	s.Require().NoError(err)
	s.NotEmpty(userList.Users, "Expected at least one user")
}

func (s *E2ESuite) TestUsers_CreateAndDelete() {
	ctx := s.T().Context()

	userName := "test-user-create"

	user, err := s.client.Users().Create(ctx, users.CreateUserRequest{
		Name: userName,
	})
	s.Require().NoError(err)
	s.Equal(userName, user.User.Name, "Expected user name to match")

	usersBefore, err := s.client.Users().List(ctx, users.UserListFilter{})
	s.Require().NoError(err)

	err = s.client.Users().Delete(ctx, user.User.ID)
	s.Require().NoError(err)

	usersAfter, err := s.client.Users().List(ctx, users.UserListFilter{})
	s.Require().NoError(err)
	s.Less(len(usersAfter.Users), len(usersBefore.Users), "Expected one less user after deletion")
}

func (s *E2ESuite) TestUsers_Rename() {
	ctx := s.T().Context()

	userName := "test-user-rename"
	newName := "test-user-renamed"

	user, err := s.client.Users().Create(ctx, users.CreateUserRequest{
		Name: userName,
	})
	s.Require().NoError(err)
	s.Equal(userName, user.User.Name, "Expected user name to match")

	renamedUser, err := s.client.Users().Rename(ctx, user.User.ID, newName)
	s.Require().NoError(err)
	s.Equal(newName, renamedUser.User.Name, "Expected renamed user name to match")

	err = s.client.Users().Delete(ctx, renamedUser.User.ID)
	s.Require().NoError(err)
}

func (s *E2ESuite) TestUsers_GetByID() {
	ctx := s.T().Context()

	userName := "test-user-get"

	user, err := s.client.Users().Create(ctx, users.CreateUserRequest{
		Name: userName,
	})
	s.Require().NoError(err)

	usersList, err := s.client.Users().List(ctx, users.UserListFilter{})
	s.Require().NoError(err)

	var found bool
	for _, u := range usersList.Users {
		if u.ID == user.User.ID {
			found = true
			s.Equal(user.User.Name, u.Name, "Expected user name to match")
			break
		}
	}
	s.True(found, "Expected to find created user in list")

	err = s.client.Users().Delete(ctx, user.User.ID)
	s.Require().NoError(err)
}
