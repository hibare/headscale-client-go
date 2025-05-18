// Package users provides a client for managing users in Headscale.
package users

import (
	"context"
	"net/http"
	"time"

	"github.com/hibare/headscale-client-go/requests"
)

// UserResourceInterface is an interface for managing users in Headscale.
type UserResourceInterface interface {
	List(ctx context.Context, filter UserListFilter) (UsersResponse, error)
	Create(ctx context.Context, request CreateUserRequest) (UserResponse, error)
	Delete(ctx context.Context, id string) error
	Rename(ctx context.Context, id, newName string) (UserResponse, error)
}

// User represents a user in Headscale.
type User struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	CreatedAt     time.Time `json:"createdAt"`
	DisplayName   string    `json:"displayName"`
	Email         string    `json:"email"`
	ProviderID    string    `json:"providerId"`
	Provider      string    `json:"provider"`
	ProfilePicURL string    `json:"profilePicUrl"`
}

// UsersResponse represents a single user response from the API.
//
//nolint:revive // This is a struct for a response from the API.
type UsersResponse struct {
	Users []User `json:"users"`
}

// UserListFilter represents a filter for a list of users.
type UserListFilter struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// List returns a list of users from the Headscale.
func (u *UserResource) List(ctx context.Context, filter UserListFilter) (UsersResponse, error) {
	var users UsersResponse

	queryParams := map[string]any{}

	if filter.ID != "" {
		queryParams["id"] = filter.ID
	}
	if filter.Name != "" {
		queryParams["name"] = filter.Name
	}
	if filter.Email != "" {
		queryParams["email"] = filter.Email
	}

	url := u.r.BuildURL("user")
	req, err := u.r.BuildRequest(ctx, http.MethodGet, url, requests.RequestOptions{
		QueryParams: queryParams,
	})
	if err != nil {
		return users, err
	}

	err = u.r.Do(ctx, req, &users)
	return users, err
}

// UserResponse represents a single user response from the API.
type UserResponse struct {
	User User `json:"user"`
}

// CreateUserRequest represents a request to create a user.
type CreateUserRequest struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	PictureURL  string `json:"pictureUrl"`
}

// Create creates a new user in Headscale.
func (u *UserResource) Create(ctx context.Context, createUserRequest CreateUserRequest) (UserResponse, error) {
	var user UserResponse

	url := u.r.BuildURL("user")
	req, err := u.r.BuildRequest(ctx, http.MethodPost, url, requests.RequestOptions{
		Body: createUserRequest,
	})
	if err != nil {
		return user, err
	}

	err = u.r.Do(ctx, req, &user)
	return user, err
}

// Delete removes a user from the Headscale.
func (u *UserResource) Delete(ctx context.Context, id string) error {
	url := u.r.BuildURL("user", id)
	req, err := u.r.BuildRequest(ctx, http.MethodDelete, url, requests.RequestOptions{})
	if err != nil {
		return err
	}

	return u.r.Do(ctx, req, nil)
}

// Rename renames a user in the Headscale.
func (u *UserResource) Rename(ctx context.Context, id, newName string) (UserResponse, error) {
	var user UserResponse

	url := u.r.BuildURL("user", id, "rename", newName)
	req, err := u.r.BuildRequest(ctx, http.MethodPost, url, requests.RequestOptions{})
	if err != nil {
		return user, err
	}

	err = u.r.Do(ctx, req, &user)
	return user, err
}

// UserResource is a struct that implements the UserResourceInterface.
type UserResource struct {
	r requests.RequestInterface
}

// NewUserResource creates a new UserResource.
func NewUserResource(r requests.RequestInterface) *UserResource {
	return &UserResource{
		r: r,
	}
}
