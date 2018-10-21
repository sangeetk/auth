package api

import (
	"context"
	"log"
	"net/url"

	ht "github.com/urantiatech/kit/transport/http"
)

// LogoutRequest - Closes the session
type LogoutRequest struct {
	AccessToken string `json:"access_token"`
}

// LogoutResponse - Returns error for invalid sessions
type LogoutResponse struct {
	Err string `json:"err,omitempty"`
}

// Logout - logs the user out of the system
func (a *AuthClient) Logout(req *LogoutRequest) (*LogoutResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + a.DNS + "/logout")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	response := resp.(LogoutResponse)
	return &response, nil
}
