package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	ht "github.com/urantiatech/kit/transport/http"
)

// IdentifyRequest - Indentifies the user using JWT
type IdentifyRequest struct {
	AccessToken string `json:"access_token"`
}

// IdentifyResponse - Sends the user details as reponse
type IdentifyResponse struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Err       string `json:"err,omitempty"`
}

// Identify - identifies a user using its JWT token
func Identify(req *IdentifyRequest, dns string) (*IdentifyResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + dns + "/identify")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeIdentifyResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	response := resp.(IdentifyResponse)
	return &response, nil
}

// decodeIdentifyResponse decodes the response from the service
func decodeIdentifyResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var response IdentifyResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
