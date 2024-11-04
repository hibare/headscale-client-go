package headscale

import (
	"context"
	"net/http"
	"time"
)

type PreAuthKeyResource struct {
	Client HeadscaleClientInterface
}

type PreAuthKey struct {
	User       string    `json:"user"`
	ID         string    `json:"id"`
	Key        string    `json:"key"`
	Reusable   bool      `json:"reusable"`
	Ephemeral  bool      `json:"ephemeral"`
	Used       bool      `json:"used"`
	Expiration time.Time `json:"expiration"`
	CreatedAt  time.Time `json:"createdAt"`
	AclTags    []string  `json:"aclTags"`
}

type PreAuthKeysResponse struct {
	PreAuthKeys []PreAuthKey `json:"preAuthKeys"`
}

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

type CreatePreAuthKeyRequest struct {
	User       string    `json:"user"`
	Reusable   bool      `json:"reusable"`
	Ephemeral  bool      `json:"ephemeral"`
	Expiration time.Time `json:"expiration"`
	AclTags    []string  `json:"aclTags"`
}

type PreAuthKeyResponse struct {
	PreAuthKey []PreAuthKey `json:"preAuthKey"`
}

func (p *PreAuthKeyResource) Create(ctx context.Context, user string, reusable bool, ephemeral bool, expiration time.Time, aclTags []string) (PreAuthKeyResponse, error) {
	var key PreAuthKeyResponse

	url := p.Client.buildURL("preauthkey")
	req, err := p.Client.buildRequest(ctx, http.MethodPost, url, requestOptions{
		body: CreatePreAuthKeyRequest{
			User:       user,
			Reusable:   reusable,
			Ephemeral:  ephemeral,
			Expiration: expiration,
			AclTags:    aclTags,
		},
	})
	if err != nil {
		return key, err
	}

	err = p.Client.do(ctx, req, &key)
	return key, err
}

type ExpirePreAuthKeyRequest struct {
	User string `json:"user"`
	Key  string `json:"key"`
}

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
