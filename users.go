package headscale

import (
	"context"
	"net/http"
	"time"
)

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

type UserResource struct {
	Client HeadscaleClientInterface
}

type UsersResponse struct {
	Users []User `json:"user"`
}

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

type UserResponse struct {
	User User `json:"user"`
}

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

type CreateUserRequest struct {
	Name string `json:"name"`
}

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

func (u *UserResource) Delete(ctx context.Context, name string) error {
	url := u.Client.buildURL("user", name)
	req, err := u.Client.buildRequest(ctx, http.MethodDelete, url, requestOptions{})
	if err != nil {
		return err
	}

	return u.Client.do(ctx, req, nil)
}

func (u *UserResource) Rename(ctx context.Context, name, newName string) error {
	url := u.Client.buildURL("user", name, "rename", newName)
	req, err := u.Client.buildRequest(ctx, http.MethodPost, url, requestOptions{})
	if err != nil {
		return err
	}

	return u.Client.do(ctx, req, nil)
}
