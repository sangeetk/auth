package api

import (
	"context"
	"log"
	"net/url"

	ht "github.com/urantiatech/kit/transport/http"
)

// LoginRequest - Login request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse - Returns JWT token on successful login
type LoginResponse struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Err          string `json:"err,omitempty"`
}

// Login - logs a user in
func (a *AuthClient) Login(req *LoginRequest) (*LoginResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + a.DNS + "/login")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	response := resp.(LoginResponse)
	return &response, nil
}
