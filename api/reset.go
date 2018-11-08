package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	ht "github.com/urantiatech/kit/transport/http"
)

// ResetRequest - Reset the password
type ResetRequest struct {
	ForgotToken string `json:"forgot_token"`
	NewPassword string `json:"new_password"`
}

// ResetResponse - Reset password response
type ResetResponse struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Err       string `json:"err,omitempty"`
}

// Reset - resets the password
func Reset(req *ResetRequest, dns string) (*ResetResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + dns + "/reset")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeResetResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	response := resp.(ResetResponse)
	return &response, nil
}

// decodeResetResponse decodes the response from the service
func decodeResetResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var response ResetResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
