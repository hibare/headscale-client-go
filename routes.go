package headscale

import (
	"context"
	"net/http"
	"time"
)

type RoutesResource struct {
	Client HeadscaleClientInterface
}

type Route struct {
	ID         string       `json:"id"`
	Node       NodeResponse `json:"node"`
	Prefix     string       `json:"prefix"`
	Advertised bool         `json:"advertised"`
	Enabled    bool         `json:"enabled"`
	IsPrimary  bool         `json:"isPrimary"`
	CreatedAt  time.Time    `json:"createdAt"`
	UpdatedAt  time.Time    `json:"updatedAt"`
	DeletedAt  time.Time    `json:"deletedAt"`
}

type RoutesResponse struct {
	Routes []Route `json:"routes"`
}

func (r *RoutesResource) List(ctx context.Context) (RoutesResponse, error) {
	var routes RoutesResponse

	url := r.Client.buildURL("routes")
	req, err := r.Client.buildRequest(ctx, http.MethodGet, url, requestOptions{})
	if err != nil {
		return routes, err
	}

	err = r.Client.do(ctx, req, &routes)
	return routes, err
}

func (r *RoutesResource) Delete(ctx context.Context, id string) error {
	url := r.Client.buildURL("routes", id)
	req, err := r.Client.buildRequest(ctx, http.MethodDelete, url, requestOptions{})
	if err != nil {
		return err
	}

	return r.Client.do(ctx, req, nil)
}

func (r *RoutesResource) Disable(ctx context.Context, id string) error {
	url := r.Client.buildURL("routes", id, "disable")
	req, err := r.Client.buildRequest(ctx, http.MethodPost, url, requestOptions{})
	if err != nil {
		return err
	}

	return r.Client.do(ctx, req, nil)
}

func (r *RoutesResource) Enable(ctx context.Context, id string) error {
	url := r.Client.buildURL("routes", id, "enable")
	req, err := r.Client.buildRequest(ctx, http.MethodPost, url, requestOptions{})
	if err != nil {
		return err
	}

	return r.Client.do(ctx, req, nil)
}
