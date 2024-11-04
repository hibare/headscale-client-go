package headscale

import (
	"context"
	"net/http"
	"time"
)

type APIKeyResource struct {
	Client HeadscaleClientInterface
}

type APIKey struct {
	ID         string    `json:"id"`
	Prefix     string    `json:"prefix"`
	Expiration time.Time `json:"expiration"`
	CreatedAt  time.Time `json:"createdAt"`
	LastSeen   time.Time `json:"lastSeen"`
}

type APIKeysResponse struct {
	APIKeys []APIKey `json:"apiKeys"`
}

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

type AddAPIKeyRequest struct {
	Expiration time.Time `json:"expiration"`
}

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

type ExpireAPIKeyRequest struct {
	Prefix string `json:"prefix"`
}

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

func (a *APIKeyResource) Delete(ctx context.Context, prefix string) error {
	url := a.Client.buildURL("apikey", prefix)
	req, err := a.Client.buildRequest(ctx, http.MethodDelete, url, requestOptions{})
	if err != nil {
		return err
	}

	return a.Client.do(ctx, req, nil)
}
