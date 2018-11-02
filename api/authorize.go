package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	ht "github.com/urantiatech/kit/transport/http"
)

// AuthorizeRequest - Authorize against the role
type AuthorizeRequest struct {
	AccessToken string `json:"access_token"`
	Domain      string `json:"domain"`
	Role        string `json:"role"`
}

// AuthorizeResponse - Returns Token or  error
type AuthorizeResponse struct {
	Authorize bool   `json:"authorize"`
	Err       string `json:"err,omitempty"`
}

// Authorize - authorize a user for the given role
func Authorize(req *AuthorizeRequest, dns string) (*AuthorizeResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + dns + "/authorize")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeAuthorizeResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	response := resp.(AuthorizeResponse)
	return &response, nil
}

// decodeAuthorizeResponse decodes the response from the service
func decodeAuthorizeResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var response AuthorizeResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
