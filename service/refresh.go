package service

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/urantiatech/kit/endpoint"
)

// Refresh - Generates a new token and extends existing session
func (Auth) Refresh(_ context.Context, req api.RefreshRequest) (api.RefreshResponse, error) {
	var response api.RefreshResponse
	var u *user.User

	// Validate the refresh token and get user info
	u, err := ParseToken(req.RefreshToken)
	if err == ErrorInvalidToken {
		response.Err = ErrorInvalidToken.Error()
		return response, nil
	}

	u, err = user.Read(u.Username)
	if err != nil || u.Confirmed != true {
		response.Err = ErrorNotFound.Error()
		return response, nil
	}

	// Create an Access JWT Token
	atoken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Username,
		"fname":    u.FirstName,
		"lname":    u.LastName,
		"email":    u.Email,
		"nbf":      time.Now().Unix(),
		"exp":      time.Now().Add(AccessTokenValidity).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	response.NewAccessToken, err = atoken.SignedString(SigningKey)
	if err != nil {
		response.Err = err.Error()
	}

	// Create an Access JWT Token
	rtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Username,
		"fname":    u.FirstName,
		"lname":    u.LastName,
		"email":    u.Email,
		"nbf":      time.Now().Unix(),
		"exp":      time.Now().Add(RefreshTokenValidity).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	response.NewRefreshToken, err = rtoken.SignedString(SigningKey)
	if err != nil {
		response.Err = err.Error()
	}

	return response, nil
}

// MakeRefreshEndpoint -
func MakeRefreshEndpoint(svc Auth) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.RefreshRequest)
		return svc.Refresh(ctx, req)
	}
}

// DecodeRefreshRequest -
func DecodeRefreshRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.RefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
