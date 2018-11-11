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

// Forgot - Resets the password
func (Auth) Forgot(ctx context.Context, req api.ForgotRequest) (api.ForgotResponse, error) {
	var response api.ForgotResponse

	if req.Username == "" {
		response.Err = ErrorInvalidRequest.Error()
		return response, nil
	}

	u, err := user.Read(req.Username)
	if err != nil || u.Confirmed != true {
		response.Err = ErrorNotFound.Error()
		return response, nil
	}

	// Create the Forgot token
	resetToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Username,
		"nbf":      time.Now().Unix(),
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	response.ResetToken, err = resetToken.SignedString(SigningKey)
	if err != nil {
		response.Err = err.Error()
		return response, nil
	}
	response.Username = u.Username
	response.FirstName = u.FirstName
	response.LastName = u.LastName
	response.Email = u.Email

	return response, nil
}

// MakeForgotEndpoint -
func MakeForgotEndpoint(svc Auth) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.ForgotRequest)
		return svc.Forgot(ctx, req)
	}
}

// DecodeForgotRequest -
func DecodeForgotRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.ForgotRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
