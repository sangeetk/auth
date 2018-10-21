package api

import (
	"context"
	"log"
	"net/url"

	ht "github.com/urantiatech/kit/transport/http"
)

// IdentifyRequest - Indentifies the user using JWT
type IdentifyRequest struct {
	AccessToken string `json:"access_token"`
}

// IdentifyResponse - Sends the user details as reponse
type IdentifyResponse struct {
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Err       string `json:"err,omitempty"`
}

// Identify - identifies a user using its JWT token
func (a *AuthClient) Identify(req *IdentifyRequest) (*IdentifyResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + a.DNS + "/identify")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	response := resp.(IdentifyResponse)
	return &response, nil
}
