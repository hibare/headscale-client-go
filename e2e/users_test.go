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

	usersBefore, err := s.client.Users().List(ctx, users.UserListFilter{})
	s.Require().NoError(err)

	userName := "test-user-create"

	user, err := s.client.Users().Create(ctx, users.CreateUserRequest{
		Name: userName,
	})
	s.Require().NoError(err)
	s.Equal(userName, user.User.Name, "Expected user name to match")

	usersAfter, err := s.client.Users().List(ctx, users.UserListFilter{})
	s.Require().NoError(err)
	s.Greater(len(usersAfter.Users), len(usersBefore.Users), "Expected one more user after creation")

	err = s.client.Users().Delete(ctx, user.User.ID)
	s.Require().NoError(err)

	usersFinal, err := s.client.Users().List(ctx, users.UserListFilter{})
	s.Require().NoError(err)
	s.Len(usersFinal.Users, len(usersBefore.Users), "Expected same number of users as before creation")
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

	// Use filter to get the specific user by ID
	usersList, err := s.client.Users().List(ctx, users.UserListFilter{ID: user.User.ID})
	s.Require().NoError(err)
	s.Require().Len(usersList.Users, 1, "Expected exactly one user with the given ID")
	s.Equal(user.User.ID, usersList.Users[0].ID, "Expected user ID to match")
	s.Equal(userName, usersList.Users[0].Name, "Expected user name to match")

	err = s.client.Users().Delete(ctx, user.User.ID)
	s.Require().NoError(err)
}
