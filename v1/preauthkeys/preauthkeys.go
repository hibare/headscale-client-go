// Package preauthkeys provides a client for managing pre-auth keys in Headscale.
package preauthkeys

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/hibare/headscale-client-go/requests"
	"github.com/hibare/headscale-client-go/v1/users"
)

var (
	// ErrNoUser is returned when a user is required but not provided.
	ErrNoUser = errors.New("user is required")
)

// PreAuthKeyResourceInterface is an interface for managing pre-auth keys in Headscale.
type PreAuthKeyResourceInterface interface {
	List(ctx context.Context, filter PreAuthKeyListFilter) (PreAuthKeysResponse, error)
	Create(ctx context.Context, createPreAuthKeyRequest CreatePreAuthKeyRequest) (PreAuthKeyResponse, error)
	Expire(ctx context.Context, user string, key string) error
}

// PreAuthKey represents a pre-auth key in Headscale.
type PreAuthKey struct {
	User       users.User `json:"user"`
	ID         string     `json:"id"`
	Key        string     `json:"key"`
	Reusable   bool       `json:"reusable"`
	Ephemeral  bool       `json:"ephemeral"`
	Used       bool       `json:"used"`
	Expiration time.Time  `json:"expiration"`
	CreatedAt  time.Time  `json:"createdAt"`
	ACLTags    []string   `json:"aclTags"`
}

// PreAuthKeysResponse represents a list of pre-auth keys response from the API.
//
//nolint:revive // This is a struct for a response from the API.
type PreAuthKeysResponse struct {
	PreAuthKeys []PreAuthKey `json:"preAuthKeys"`
}

// PreAuthKeyListFilter represents a filter for listing pre-auth keys.
type PreAuthKeyListFilter struct {
	User int `json:"user"`
}

// List returns a list of pre-auth keys from the Headscale.
func (p *PreAuthKeyResource) List(ctx context.Context, filter PreAuthKeyListFilter) (PreAuthKeysResponse, error) {
	var keys PreAuthKeysResponse

	queryParams := map[string]any{}
	if filter.User == 0 {
		return keys, ErrNoUser
	}

	queryParams["user"] = filter.User
	url := p.r.BuildURL("preauthkey")
	req, err := p.r.BuildRequest(ctx, http.MethodGet, url, requests.RequestOptions{
		QueryParams: queryParams,
	})
	if err != nil {
		return keys, err
	}

	err = p.r.Do(ctx, req, &keys)
	return keys, err
}

// CreatePreAuthKeyRequest represents a request to create a pre-auth key.
type CreatePreAuthKeyRequest struct {
	User       string    `json:"user"`
	Reusable   bool      `json:"reusable"`
	Ephemeral  bool      `json:"ephemeral"`
	Expiration time.Time `json:"expiration"`
	ACLTags    []string  `json:"aclTags"`
}

// PreAuthKeyResponse represents a single pre-auth key response from the API.
type PreAuthKeyResponse struct {
	PreAuthKey PreAuthKey `json:"preAuthKey"`
}

// Create creates a new pre-auth key in Headscale.
func (p *PreAuthKeyResource) Create(ctx context.Context, createPreAuthKeyRequest CreatePreAuthKeyRequest) (PreAuthKeyResponse, error) {
	var key PreAuthKeyResponse

	url := p.r.BuildURL("preauthkey")
	req, err := p.r.BuildRequest(ctx, http.MethodPost, url, requests.RequestOptions{
		Body: createPreAuthKeyRequest,
	})
	if err != nil {
		return key, err
	}

	err = p.r.Do(ctx, req, &key)
	return key, err
}

// ExpirePreAuthKeyRequest represents a request to expire a pre-auth key.
type ExpirePreAuthKeyRequest struct {
	User string `json:"user"`
	Key  string `json:"key"`
}

// Expire expires a pre-auth key in Headscale.
func (p *PreAuthKeyResource) Expire(ctx context.Context, user string, key string) error {
	url := p.r.BuildURL("preauthkey", "expire")
	req, err := p.r.BuildRequest(ctx, http.MethodPost, url, requests.RequestOptions{
		Body: ExpirePreAuthKeyRequest{
			User: user,
			Key:  key,
		},
	})
	if err != nil {
		return err
	}

	return p.r.Do(ctx, req, nil)
}

// PreAuthKeyResource is a struct that implements the PreAuthKeyResourceInterface.
type PreAuthKeyResource struct {
	r requests.RequestInterface
}

// NewPreAuthKeyResource creates a new PreAuthKeyResource.
func NewPreAuthKeyResource(r requests.RequestInterface) *PreAuthKeyResource {
	return &PreAuthKeyResource{
		r: r,
	}
}
