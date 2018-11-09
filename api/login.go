package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	ht "github.com/urantiatech/kit/transport/http"
)

// LoginRequest - Login request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Domain   string `json:"domain"`
}

// LoginResponse - Returns JWT token on successful login
type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Username     string `json:"username"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Err          string `json:"err,omitempty"`
}

// Login - logs a user in
func Login(req *LoginRequest, dns string) (*LoginResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + dns + "/login")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeLoginResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	response := resp.(LoginResponse)
	return &response, nil
}

// decodeLoginResponse decodes the response from the service
func decodeLoginResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var response LoginResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
