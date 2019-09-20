package service

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/token"
	"git.urantiatech.com/auth/auth/user"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/crypto/bcrypt"
)

// Reset - Resets the password
func (a Auth) Reset(ctx context.Context, req *api.ResetRequest) (*api.ResetResponse, error) {
	var response = &api.ResetResponse{}

	if req.ResetToken == "" || req.NewPassword == "" {
		response.Err = ErrorInvalidRequest.Error()
		return response, nil
	}

	t, err := token.ParseToken(req.ResetToken)
	if err != nil {
		response.Err = token.ErrorInvalidToken.Error()
		return response, nil
	}

	// Check against Blacklisted tokens
	if _, found := token.BlacklistAccessTokens.Get(req.ResetToken); found {
		response.Err = token.ErrorInvalidToken.Error()
		return response, nil
	}

	// Read user details
	u, err := user.Read(t.Username)
	if err != nil {
		response.Err = token.ErrorInvalidToken.Error()
		return response, nil
	}

	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 11)
	if err != nil {
		response.Err = ErrorUnknown.Error()
		return response, nil
	}

	// Return error if password is changed in last 24 hours
	if !u.PasswordUpdatedAt.Add(24 * time.Hour).Before(time.Now()) {
		response.Err = token.ErrorInvalidToken.Error()
		return response, nil
	}

	u.Password = PasswordHash
	u.PasswordUpdatedAt = time.Now()
	u.Save()

	response.Username = u.Username
	response.FirstName = u.FirstName
	response.LastName = u.LastName
	response.Email = u.Email
	response.Domain = t.Domain
	response.Roles = u.GetRoles(t.Domain)

	return response, nil
}

// MakeResetEndpoint -
func MakeResetEndpoint(svc AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.ResetRequest)
		return svc.Reset(ctx, &req)
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
