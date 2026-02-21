// Package apikeys provides a client for managing API keys in Headscale.
package apikeys

import (
	"context"
	"net/http"
	"time"

	"github.com/hibare/headscale-client-go/requests"
)

// APIKeyResourceInterface is an interface for managing API keys in Headscale.
type APIKeyResourceInterface interface {
	List(ctx context.Context) (APIKeysResponse, error)
	Create(ctx context.Context, createAPIKeyRequest CreateAPIKeyRequest) (CreateAPIKeyResponse, error)
	Expire(ctx context.Context, prefix string) error
	ExpireByID(ctx context.Context, id string) error
	Delete(ctx context.Context, prefix string) error
	DeleteByID(ctx context.Context, id string) error
}

// APIKey represents an API key in Headscale.
type APIKey struct {
	ID         string    `json:"id"`
	Prefix     string    `json:"prefix"`
	Expiration time.Time `json:"expiration"`
	CreatedAt  time.Time `json:"createdAt"`
	LastSeen   time.Time `json:"lastSeen"`
}

// APIKeysResponse represents a single API key response from the API.
//
//nolint:revive // This is a struct for a response from the API.
type APIKeysResponse struct {
	APIKeys []APIKey `json:"apiKeys"`
}

// List returns a list of API keys from the Headscale.
func (a *APIKeyResource) List(ctx context.Context) (APIKeysResponse, error) {
	var keys APIKeysResponse

	url := a.R.BuildURL("apikey")
	req, err := a.R.BuildRequest(ctx, http.MethodGet, url, requests.RequestOptions{})
	if err != nil {
		return keys, err
	}

	err = a.R.Do(ctx, req, &keys)
	return keys, err
}

// CreateAPIKeyRequest represents a request to create a new API key.
type CreateAPIKeyRequest struct {
	Expiration time.Time `json:"expiration"`
}

// CreateAPIKeyResponse represents a response from the API when creating a new API key.
type CreateAPIKeyResponse struct {
	APIKey string `json:"apiKey"`
}

// Create creates a new API key in Headscale.
func (a *APIKeyResource) Create(ctx context.Context, createAPIKeyRequest CreateAPIKeyRequest) (CreateAPIKeyResponse, error) {
	var key CreateAPIKeyResponse

	url := a.R.BuildURL("apikey")
	req, err := a.R.BuildRequest(ctx, http.MethodPost, url, requests.RequestOptions{
		Body: createAPIKeyRequest,
	})
	if err != nil {
		return key, err
	}

	err = a.R.Do(ctx, req, &key)
	return key, err
}

// ExpireAPIKeyRequest represents a request to expire an API key.
type ExpireAPIKeyRequest struct {
	Prefix string `json:"prefix,omitempty"`
	ID     string `json:"id,omitempty"`
}

// Expire expires an API key in Headscale.
func (a *APIKeyResource) Expire(ctx context.Context, prefix string) error {
	url := a.R.BuildURL("apikey", "expire")
	req, err := a.R.BuildRequest(ctx, http.MethodPost, url, requests.RequestOptions{
		Body: ExpireAPIKeyRequest{
			Prefix: prefix,
		},
	})
	if err != nil {
		return err
	}

	return a.R.Do(ctx, req, nil)
}

// ExpireByID expires an API key by ID in Headscale.
func (a *APIKeyResource) ExpireByID(ctx context.Context, id string) error {
	url := a.R.BuildURL("apikey", "expire")
	req, err := a.R.BuildRequest(ctx, http.MethodPost, url, requests.RequestOptions{
		Body: ExpireAPIKeyRequest{
			ID: id,
		},
	})
	if err != nil {
		return err
	}

	return a.R.Do(ctx, req, nil)
}

// Delete removes an API key from the Headscale.
func (a *APIKeyResource) Delete(ctx context.Context, prefix string) error {
	url := a.R.BuildURL("apikey", prefix)
	req, err := a.R.BuildRequest(ctx, http.MethodDelete, url, requests.RequestOptions{})
	if err != nil {
		return err
	}

	return a.R.Do(ctx, req, nil)
}

// DeleteByID removes an API key by ID from the Headscale.
func (a *APIKeyResource) DeleteByID(ctx context.Context, id string) error {
	url := a.R.BuildURL("apikey", "-")
	req, err := a.R.BuildRequest(ctx, http.MethodDelete, url, requests.RequestOptions{
		QueryParams: map[string]any{"id": id},
	})
	if err != nil {
		return err
	}

	return a.R.Do(ctx, req, nil)
}

// APIKeyResource is a struct that implements the APIKeyResourceInterface.
type APIKeyResource struct {
	R requests.RequestInterface
}
