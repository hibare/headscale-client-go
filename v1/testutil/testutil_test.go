package testutil

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/hibare/headscale-client-go/requests"
)

func TestRunResourceTest(t *testing.T) {
	type dummyResponse struct {
		Val string
	}

	fixture := TestFixture[dummyResponse]{
		Endpoint:    "dummy",
		Method:      http.MethodGet,
		SuccessResp: dummyResponse{Val: "ok"},
		BuildErr:    errors.New("custom build error"),
		DoErr:       errors.New("custom do error"),
	}

	action := func(ctx context.Context, mockReq *requests.MockRequest) (dummyResponse, error) {
		var resp dummyResponse
		url := mockReq.BuildURL("dummy")
		req, err := mockReq.BuildRequest(ctx, http.MethodGet, url, requests.RequestOptions{})
		if err != nil {
			return resp, err
		}
		err = mockReq.Do(ctx, req, &resp)
		return resp, err
	}

	RunResourceTest(t, fixture, action)
}
