package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	ht "github.com/urantiatech/kit/transport/http"
)

// ConfirmRequest - Confirms and logs the user in
type ConfirmRequest struct {
	ConfirmToken string `json:"confirm_token"`
}

// ConfirmResponse - Returns error if registration confirmation fails
type ConfirmResponse struct {
	AccessToken string   `json:"access_token"`
	Username    string   `json:"username"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	Email       string   `json:"email"`
	Domain      string   `json:"domain"`
	Roles       []string `json:"roles"`
	Err         string   `json:"err,omitempty"`
}

// Confirm - confirms the new user registration
func Confirm(req *ConfirmRequest, dns string) (*ConfirmResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + dns + "/confirm")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeConfirmResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	response := resp.(ConfirmResponse)
	return &response, nil
}

// decodeConfirmResponse decodes the response from the service
func decodeConfirmResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var response ConfirmResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
