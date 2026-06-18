// Package preauthkeys provides a client for managing pre-auth keys in Headscale.
package preauthkeys

import (
	"context"
	"net/http"
	"time"

	"github.com/hibare/headscale-client-go/requests"
	"github.com/hibare/headscale-client-go/v1/users"
)

// PreAuthKeyResourceInterface is an interface for managing pre-auth keys in Headscale.
type PreAuthKeyResourceInterface interface {
	List(ctx context.Context) (PreAuthKeysResponse, error)
	Create(ctx context.Context, createPreAuthKeyRequest CreatePreAuthKeyRequest) (PreAuthKeyResponse, error)
	Expire(ctx context.Context, id string) error
	Delete(ctx context.Context, id string) error
}

// PreAuthKeyResource is a struct that implements the PreAuthKeyResourceInterface.
type PreAuthKeyResource struct {
	r requests.RequestInterface
}

// NewPreAuthKeyResource creates a new PreAuthKeyResource.
func NewPreAuthKeyResource(r requests.RequestInterface) *PreAuthKeyResource {
	return &PreAuthKeyResource{r: r}
}

// PreAuthKey represents a pre-auth key in Headscale.
type PreAuthKey struct {
	ID         string     `json:"id"`
	User       users.User `json:"user,omitempty"`
	Key        string     `json:"key,omitempty"`
	Reusable   bool       `json:"reusable"`
	Ephemeral  bool       `json:"ephemeral"`
	Used       bool       `json:"used"`
	Expiration time.Time  `json:"expiration"`
	CreatedAt  time.Time  `json:"createdAt"`
	ACLTags    []string   `json:"aclTags"`
}

// PreAuthKeysResponse represents a list of pre-auth keys response from the API.
type PreAuthKeysResponse struct {
	PreAuthKeys []PreAuthKey `json:"preAuthKeys"`
}

// List returns a list of pre-auth keys from the Headscale.
func (p *PreAuthKeyResource) List(ctx context.Context) (PreAuthKeysResponse, error) {
	var keys PreAuthKeysResponse

	url := p.r.BuildURL("preauthkey")
	req, err := p.r.BuildRequest(ctx, http.MethodGet, url, requests.RequestOptions{})
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
	ID string `json:"id"`
}

// Expire expires a pre-auth key in Headscale.
func (p *PreAuthKeyResource) Expire(ctx context.Context, id string) error {
	url := p.r.BuildURL("preauthkey", "expire")
	req, err := p.r.BuildRequest(ctx, http.MethodPost, url, requests.RequestOptions{
		Body: ExpirePreAuthKeyRequest{ID: id},
	})
	if err != nil {
		return err
	}

	return p.r.Do(ctx, req, nil)
}

// Delete removes a pre-auth key from the Headscale.
func (p *PreAuthKeyResource) Delete(ctx context.Context, id string) error {
	url := p.r.BuildURL("preauthkey")
	req, err := p.r.BuildRequest(ctx, http.MethodDelete, url, requests.RequestOptions{
		QueryParams: map[string]any{"id": id},
	})
	if err != nil {
		return err
	}

	return p.r.Do(ctx, req, nil)
}
