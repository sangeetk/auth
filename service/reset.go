package service

import (
	"context"
	"encoding/json"
	"net/http"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/user"
	"github.com/urantiatech/kit/endpoint"
	"golang.org/x/crypto/bcrypt"
)

// Reset - Resets the password
func (a Auth) Reset(ctx context.Context, req api.ResetRequest) (api.ResetResponse, error) {
	var response api.ResetResponse

	if req.ResetToken == "" || req.NewPassword == "" {
		response.Err = ErrorInvalidRequest.Error()
		return response, nil
	}

	identify := api.IdentifyRequest{AccessToken: req.ResetToken}
	id, err := a.Identify(ctx, identify)
	if err == nil || id.Err != "" {
		response.Err = ErrorInvalidToken.Error()
		return response, nil
	}

	// Read user details
	u, err := user.Read(id.Username)
	if err != nil || u.Confirmed != false {
		response.Err = ErrorInvalidToken.Error()
		return response, nil
	}

	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 11)
	if err != nil {
		response.Err = ErrorUnknown.Error()
		return response, nil
	}

	u.Password = PasswordHash
	u.Save()

	response.Username = u.Username
	response.FirstName = u.FirstName
	response.LastName = u.LastName
	response.Email = u.Email

	return response, nil
}

// MakeResetEndpoint -
func MakeResetEndpoint(svc Auth) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.ResetRequest)
		return svc.Reset(ctx, req)
	}
}

// DecodeResetRequest -
func DecodeResetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.ResetRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
