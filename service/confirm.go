package service

import (
	"context"
	"encoding/json"
	"net/http"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/user"
	"github.com/urantiatech/kit/endpoint"
)

// Confirm - Activates the user account after confirmation
func (a Auth) Confirm(ctx context.Context, req api.ConfirmRequest) (api.ConfirmResponse, error) {
	var response api.ConfirmResponse
	var u *user.User
	var err error

	u, err = ParseToken(req.ConfirmToken)
	if err != nil {
		response.Err = ErrorInvalidToken.Error()
		return response, nil
	}

	// Read user details
	u, err = user.Read(u.Username)
	if err != nil || u.Confirmed != false {
		response.Err = ErrorInvalidToken.Error()
		return response, nil
	}

	u.Confirmed = true
	u.Save()

	response.Username = u.Username
	response.FirstName = u.FirstName
	response.LastName = u.LastName
	response.Email = u.Email

	return response, nil
}

// MakeConfirmEndpoint -
func MakeConfirmEndpoint(svc Auth) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.ConfirmRequest)
		return svc.Confirm(ctx, req)
	}
}

// DecodeConfirmRequest -
func DecodeConfirmRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.ConfirmRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
