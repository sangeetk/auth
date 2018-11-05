package service

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/user"
	"github.com/urantiatech/kit/endpoint"
	"golang.org/x/crypto/bcrypt"
)

// Reset - Resets the password
func (Auth) Reset(_ context.Context, req api.ResetRequest) (api.ResetResponse, error) {
	var response api.ResetResponse

	if req.ForgotToken == "" || req.Password == "" {
		response.Err = ErrorInvalidRequest.Error()
		return response, nil
	}

	u, err := user.Read(req.Username)
	if err != nil || req.ForgotToken != u.ForgotToken || u.ForgotTokenExpiry.Unix() < time.Now().Unix() {
		response.Err = ErrorExpiredToken.Error()
		return response, nil
	}

	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 11)
	if err != nil {
		response.Err = ErrorUnknown.Error()
		return response, nil
	}
	u.ForgotToken = ""
	u.ForgotTokenExpiry = time.Unix(0, 0)
	u.Password = PasswordHash
	u.Save()

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
