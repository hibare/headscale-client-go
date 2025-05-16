package headscale

import (
	"context"
	"net/http"
	"time"
)

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

// UserResource is a struct that provides methods to interact with the users API of Headscale.
type UserResource struct {
	Client ClientInterface
}

// UsersResponse represents a single user response from the API.
type UsersResponse struct {
	Users []User `json:"user"`
}

// List returns a list of users from the Headscale.
func (u *UserResource) List(ctx context.Context) (UsersResponse, error) {
	var users UsersResponse

	url := u.Client.buildURL("user")
	req, err := u.Client.buildRequest(ctx, http.MethodGet, url, requestOptions{})
	if err != nil {
		return users, err
	}

	err = u.Client.do(ctx, req, &users)
	return users, err
}

// UserResponse represents a single user response from the API.
type UserResponse struct {
	User User `json:"user"`
}

// Get retrieves a user by its name from the Headscale.
func (u *UserResource) Get(ctx context.Context, name string) (UserResponse, error) {
	var user UserResponse

	url := u.Client.buildURL("user", name)
	req, err := u.Client.buildRequest(ctx, http.MethodGet, url, requestOptions{})
	if err != nil {
		return user, err
	}

	err = u.Client.do(ctx, req, &user)
	return user, err
}

// CreateUserRequest represents a request to create a user.
type CreateUserRequest struct {
	Name string `json:"name"`
}

// Create creates a new user in Headscale.
func (u *UserResource) Create(ctx context.Context, name string) (User, error) {
	var user User

	url := u.Client.buildURL("user")
	req, err := u.Client.buildRequest(ctx, http.MethodPost, url, requestOptions{
		body: CreateUserRequest{
			Name: name,
		},
	})
	if err != nil {
		return user, err
	}

	err = u.Client.do(ctx, req, &user)
	return user, err
}

// Delete removes a user from the Headscale.
func (u *UserResource) Delete(ctx context.Context, name string) error {
	url := u.Client.buildURL("user", name)
	req, err := u.Client.buildRequest(ctx, http.MethodDelete, url, requestOptions{})
	if err != nil {
		return err
	}

	return u.Client.do(ctx, req, nil)
}

// Rename renames a user in the Headscale.
func (u *UserResource) Rename(ctx context.Context, name, newName string) error {
	url := u.Client.buildURL("user", name, "rename", newName)
	req, err := u.Client.buildRequest(ctx, http.MethodPost, url, requestOptions{})
	if err != nil {
		return err
	}

	return u.Client.do(ctx, req, nil)
}
