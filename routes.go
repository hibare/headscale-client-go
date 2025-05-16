package headscale

import (
	"context"
	"net/http"
	"time"
)

// RoutesResource is a struct that provides methods to interact with the routes API of Headscale.
type RoutesResource struct {
	Client ClientInterface
}

// Route represents a route in Headscale.
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

// RoutesResponse represents a single route response from the API.
type RoutesResponse struct {
	Routes []Route `json:"routes"`
}

// List returns a list of routes from the Headscale.
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

// Delete removes a route from the Headscale.
func (r *RoutesResource) Delete(ctx context.Context, id string) error {
	url := r.Client.buildURL("routes", id)
	req, err := r.Client.buildRequest(ctx, http.MethodDelete, url, requestOptions{})
	if err != nil {
		return err
	}

	return r.Client.do(ctx, req, nil)
}

// Disable disables a route in the Headscale.
func (r *RoutesResource) Disable(ctx context.Context, id string) error {
	url := r.Client.buildURL("routes", id, "disable")
	req, err := r.Client.buildRequest(ctx, http.MethodPost, url, requestOptions{})
	if err != nil {
		return err
	}

	return r.Client.do(ctx, req, nil)
}

// Enable enables a route in the Headscale.
func (r *RoutesResource) Enable(ctx context.Context, id string) error {
	url := r.Client.buildURL("routes", id, "enable")
	req, err := r.Client.buildRequest(ctx, http.MethodPost, url, requestOptions{})
	if err != nil {
		return err
	}

	return r.Client.do(ctx, req, nil)
}
