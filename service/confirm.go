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
func (Auth) Confirm(_ context.Context, req api.ConfirmRequest) (api.ConfirmResponse, error) {
	var response api.ConfirmResponse

	user, err := user.Read(req.Username)
	if err != nil || user.ConfirmToken != req.ConfirmToken || user.Confirmed != false {
		response.Err = ErrorInvalidToken.Error()
		return response, nil
	}

	user.ConfirmToken = ""
	user.Confirmed = true

	user.Save()
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
