package api

import (
	"context"
	"log"
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
func (a *AuthClient) Refresh(req *RefreshRequest) (*RefreshResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + a.DNS + "/refresh")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	response := resp.(RefreshResponse)
	return &response, nil
}
