package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	ht "github.com/urantiatech/kit/transport/http"
)

// RefreshRequest - Request to refresh the JWT token
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// RefreshResponse - Returns new JWT token on success
type RefreshResponse struct {
	NewAccessToken  string `json:"new_access_token,omitempty"`
	NewRefreshToken string `json:"new_refresh_token,omitempty"`
	Err             string `json:"err,omitempty"`
}

// Refresh - extends the session by refreshing the JWT token
func Refresh(req *RefreshRequest, dns string) (*RefreshResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + dns + "/refresh")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeRefreshResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	response := resp.(RefreshResponse)
	return &response, nil
}

// decodeRefreshResponse decodes the response from the service
func decodeRefreshResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var response RefreshResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
