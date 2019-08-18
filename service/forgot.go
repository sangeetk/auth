package service

import (
	"context"
	"encoding/json"
	"net/http"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/token"
	"git.urantiatech.com/auth/auth/user"
	"github.com/go-kit/kit/endpoint"
)

// Forgot - Resets the password
func (Auth) Forgot(ctx context.Context, req *api.ForgotRequest) (*api.ForgotResponse, error) {
	var response = &api.ForgotResponse{}

	if req.Username == "" {
		response.Err = ErrorInvalidRequest.Error()
		return response, nil
	}

	u, err := user.Read(req.Username)
	if err != nil || !u.DeletedAt.IsZero() {
		response.Err = ErrorNotFound.Error()
		return response, nil
	}

	// Create the Forgot token
	response.ResetToken, err = token.NewToken(u, req.Domain, token.ResetTokenValidity)
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
func MakeForgotEndpoint(svc AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.ForgotRequest)
		return svc.Forgot(ctx, &req)
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
