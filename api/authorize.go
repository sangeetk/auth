package api

import (
	"context"
	"log"
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
func (a *AuthClient) Authorize(req *AuthorizeRequest) (*AuthorizeResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + a.DNS + "/authorize")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	response := resp.(AuthorizeResponse)
	return &response, nil
}
