package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
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
func Logout(req *LogoutRequest, dns string) (*LogoutResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + dns + "/logout")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeLogoutResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	response := resp.(LogoutResponse)
	return &response, nil
}

// decodeLogoutResponse decodes the response from the service
func decodeLogoutResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var response LogoutResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
