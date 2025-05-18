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
	Create(ctx context.Context, expiration time.Time) (CreateAPIKeyResponse, error)
	Expire(ctx context.Context, prefix string) error
	Delete(ctx context.Context, prefix string) error
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

	url := a.r.BuildURL("apikey")
	req, err := a.r.BuildRequest(ctx, http.MethodGet, url, requests.RequestOptions{})
	if err != nil {
		return keys, err
	}

	err = a.r.Do(ctx, req, &keys)
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
func (a *APIKeyResource) Create(ctx context.Context, expiration time.Time) (CreateAPIKeyResponse, error) {
	var key CreateAPIKeyResponse

	url := a.r.BuildURL("apikey")
	req, err := a.r.BuildRequest(ctx, http.MethodPost, url, requests.RequestOptions{
		Body: CreateAPIKeyRequest{
			Expiration: expiration,
		},
	})
	if err != nil {
		return key, err
	}

	err = a.r.Do(ctx, req, &key)
	return key, err
}

// ExpireAPIKeyRequest represents a request to expire an API key.
type ExpireAPIKeyRequest struct {
	Prefix string `json:"prefix"`
}

// Expire expires an API key in Headscale.
func (a *APIKeyResource) Expire(ctx context.Context, prefix string) error {
	url := a.r.BuildURL("apikey", "expire")
	req, err := a.r.BuildRequest(ctx, http.MethodPost, url, requests.RequestOptions{
		Body: ExpireAPIKeyRequest{
			Prefix: prefix,
		},
	})
	if err != nil {
		return err
	}

	return a.r.Do(ctx, req, nil)
}

// Delete removes an API key from the Headscale.
func (a *APIKeyResource) Delete(ctx context.Context, prefix string) error {
	url := a.r.BuildURL("apikey", prefix)
	req, err := a.r.BuildRequest(ctx, http.MethodDelete, url, requests.RequestOptions{})
	if err != nil {
		return err
	}

	return a.r.Do(ctx, req, nil)
}

// APIKeyResource is a struct that implements the APIKeyResourceInterface.
type APIKeyResource struct {
	r requests.RequestInterface
}

// NewAPIKeyResource creates a new APIKeyResource.
func NewAPIKeyResource(r requests.RequestInterface) *APIKeyResource {
	return &APIKeyResource{
		r: r,
	}
}
