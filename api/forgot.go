package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	ht "github.com/urantiatech/kit/transport/http"
)

// ForgotRequest - Forgot the account
type ForgotRequest struct {
	Username string `json:"username"`
}

// ForgotResponse - Reset the new password
type ForgotResponse struct {
	ForgotToken string `json:"forgot_token,omitempty"`
	Err         string `json:"err,omitempty"`
}

// Forgot - sends a recovery mail to reset the password
func Forgot(req *ForgotRequest, dns string) (*ForgotResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + dns + "/forgot")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeForgotResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	response := resp.(ForgotResponse)
	return &response, nil
}

// decodeForgotResponse decodes the response from the service
func decodeForgotResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var response ForgotResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
