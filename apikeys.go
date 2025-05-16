package headscale

import (
	"context"
	"net/http"
	"time"
)

// APIKeyResource is a resource for managing API keys in Headscale.
type APIKeyResource struct {
	Client ClientInterface
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
type APIKeysResponse struct {
	APIKeys []APIKey `json:"apiKeys"`
}

// List returns a list of API keys from the Headscale.
func (a *APIKeyResource) List(ctx context.Context) (APIKeysResponse, error) {
	var keys APIKeysResponse

	url := a.Client.buildURL("apikey")
	req, err := a.Client.buildRequest(ctx, http.MethodGet, url, requestOptions{})
	if err != nil {
		return keys, err
	}

	err = a.Client.do(ctx, req, &keys)
	return keys, err
}

// AddAPIKeyRequest represents a request to create a new API key.
type AddAPIKeyRequest struct {
	Expiration time.Time `json:"expiration"`
}

// Create creates a new API key in Headscale.
func (a *APIKeyResource) Create(ctx context.Context, expiration time.Time) (APIKey, error) {
	var key APIKey

	url := a.Client.buildURL("apikey")
	req, err := a.Client.buildRequest(ctx, http.MethodPost, url, requestOptions{
		body: AddAPIKeyRequest{
			Expiration: expiration,
		},
	})
	if err != nil {
		return key, err
	}

	err = a.Client.do(ctx, req, &key)
	return key, err
}

// ExpireAPIKeyRequest represents a request to expire an API key.
type ExpireAPIKeyRequest struct {
	Prefix string `json:"prefix"`
}

// Expire expires an API key in Headscale.
func (a *APIKeyResource) Expire(ctx context.Context, prefix string) error {
	url := a.Client.buildURL("apikey", "expire")
	req, err := a.Client.buildRequest(ctx, http.MethodPost, url, requestOptions{
		body: ExpireAPIKeyRequest{
			Prefix: prefix,
		},
	})
	if err != nil {
		return err
	}

	return a.Client.do(ctx, req, nil)
}

// Delete removes an API key from the Headscale.
func (a *APIKeyResource) Delete(ctx context.Context, prefix string) error {
	url := a.Client.buildURL("apikey", prefix)
	req, err := a.Client.buildRequest(ctx, http.MethodDelete, url, requestOptions{})
	if err != nil {
		return err
	}

	return a.Client.do(ctx, req, nil)
}
