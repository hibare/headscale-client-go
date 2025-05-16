package headscale

import (
	"context"
	"net/http"
	"time"
)

// PreAuthKeyResource is a resource for managing pre-auth keys in Headscale.
type PreAuthKeyResource struct {
	Client ClientInterface
}

// PreAuthKey represents a pre-auth key in Headscale.
type PreAuthKey struct {
	User       string    `json:"user"`
	ID         string    `json:"id"`
	Key        string    `json:"key"`
	Reusable   bool      `json:"reusable"`
	Ephemeral  bool      `json:"ephemeral"`
	Used       bool      `json:"used"`
	Expiration time.Time `json:"expiration"`
	CreatedAt  time.Time `json:"createdAt"`
	ACLTags    []string  `json:"aclTags"`
}

// PreAuthKeysResponse represents a list of pre-auth keys response from the API.
type PreAuthKeysResponse struct {
	PreAuthKeys []PreAuthKey `json:"preAuthKeys"`
}

// List returns a list of pre-auth keys from the Headscale.
func (p *PreAuthKeyResource) List(ctx context.Context) (PreAuthKeysResponse, error) {
	var keys PreAuthKeysResponse

	url := p.Client.buildURL("preauthkey")
	req, err := p.Client.buildRequest(ctx, http.MethodGet, url, requestOptions{})
	if err != nil {
		return keys, err
	}

	err = p.Client.do(ctx, req, &keys)
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
	PreAuthKey []PreAuthKey `json:"preAuthKey"`
}

// Create creates a new pre-auth key in Headscale.
func (p *PreAuthKeyResource) Create(
	ctx context.Context,
	user string,
	reusable bool,
	ephemeral bool,
	expiration time.Time,
	aclTags []string,
) (PreAuthKeyResponse, error) {
	var key PreAuthKeyResponse

	url := p.Client.buildURL("preauthkey")
	req, err := p.Client.buildRequest(ctx, http.MethodPost, url, requestOptions{
		body: CreatePreAuthKeyRequest{
			User:       user,
			Reusable:   reusable,
			Ephemeral:  ephemeral,
			Expiration: expiration,
			ACLTags:    aclTags,
		},
	})
	if err != nil {
		return key, err
	}

	err = p.Client.do(ctx, req, &key)
	return key, err
}

// ExpirePreAuthKeyRequest represents a request to expire a pre-auth key.
type ExpirePreAuthKeyRequest struct {
	User string `json:"user"`
	Key  string `json:"key"`
}

// Expire expires a pre-auth key in Headscale.
func (p *PreAuthKeyResource) Expire(ctx context.Context, user string, key string) error {
	url := p.Client.buildURL("preauthkey", "expire")
	req, err := p.Client.buildRequest(ctx, http.MethodPost, url, requestOptions{
		body: ExpirePreAuthKeyRequest{
			User: user,
			Key:  key,
		},
	})
	if err != nil {
		return err
	}

	return p.Client.do(ctx, req, nil)
}
