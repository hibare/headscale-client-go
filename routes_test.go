package headscale

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRoutesResource_List(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		routesResource := &RoutesResource{Client: client}
		client.(*MockClient).On("buildURL", "routes").Return(&url.URL{Path: "http://example.com/api/v1/routes"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodGet, &url.URL{Path: "http://example.com/api/v1/routes"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			resp := args.Get(1).(*RoutesResponse)
			resp.Routes = []Route{{ID: "1", Prefix: "192.168.1.0/24"}}
		})

		routes, err := routesResource.List(context.Background())
		assert.NoError(t, err)
		assert.Len(t, routes.Routes, 1)
		assert.Equal(t, "192.168.1.0/24", routes.Routes[0].Prefix)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		routesResource := &RoutesResource{Client: client}
		client.(*MockClient).On("buildURL", "routes").Return(&url.URL{Path: "http://example.com/api/v1/routes"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodGet, &url.URL{Path: "http://example.com/api/v1/routes"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		routes, err := routesResource.List(context.Background())
		assert.Error(t, err)
		assert.Empty(t, routes.Routes)
	})
}

func TestRoutesResource_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		routesResource := &RoutesResource{Client: client}
		client.(*MockClient).On("buildURL", "routes", "1").Return(&url.URL{Path: "http://example.com/api/v1/routes/1"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodDelete, &url.URL{Path: "http://example.com/api/v1/routes/1"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil)

		err := routesResource.Delete(context.Background(), "1")
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		routesResource := &RoutesResource{Client: client}
		client.(*MockClient).On("buildURL", "routes", "1").Return(&url.URL{Path: "http://example.com/api/v1/routes/1"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodDelete, &url.URL{Path: "http://example.com/api/v1/routes/1"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		err := routesResource.Delete(context.Background(), "1")
		assert.Error(t, err)
	})
}

func TestRoutesResource_Disable(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		routesResource := &RoutesResource{Client: client}
		client.(*MockClient).On("buildURL", "routes", "1", "disable").Return(&url.URL{Path: "http://example.com/api/v1/routes/1/disable"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPost, &url.URL{Path: "http://example.com/api/v1/routes/1/disable"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil)

		err := routesResource.Disable(context.Background(), "1")
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		routesResource := &RoutesResource{Client: client}
		client.(*MockClient).On("buildURL", "routes", "1", "disable").Return(&url.URL{Path: "http://example.com/api/v1/routes/1/disable"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPost, &url.URL{Path: "http://example.com/api/v1/routes/1/disable"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		err := routesResource.Disable(context.Background(), "1")
		assert.Error(t, err)
	})
}

func TestRoutesResource_Enable(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMockClient()
		routesResource := &RoutesResource{Client: client}
		client.(*MockClient).On("buildURL", "routes", "1", "enable").Return(&url.URL{Path: "http://example.com/api/v1/routes/1/enable"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPost, &url.URL{Path: "http://example.com/api/v1/routes/1/enable"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(nil)

		err := routesResource.Enable(context.Background(), "1")
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		client := NewMockClient()
		routesResource := &RoutesResource{Client: client}
		client.(*MockClient).On("buildURL", "routes", "1", "enable").Return(&url.URL{Path: "http://example.com/api/v1/routes/1/enable"})
		client.(*MockClient).On("buildRequest", mock.Anything, http.MethodPost, &url.URL{Path: "http://example.com/api/v1/routes/1/enable"}, requestOptions{}).Return(&http.Request{}, nil)
		client.(*MockClient).On("do", mock.Anything, mock.Anything).Return(errors.New("request failed"))

		err := routesResource.Enable(context.Background(), "1")
		assert.Error(t, err)
	})
}
