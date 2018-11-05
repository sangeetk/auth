package service

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/user"
	"github.com/urantiatech/kit/endpoint"
)

// Forgot - Resets the password
func (Auth) Forgot(_ context.Context, req api.ForgotRequest) (api.ForgotResponse, error) {
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

	if u.ForgotToken == "" || u.ForgotTokenExpiry.Unix() < time.Now().Unix() {
		u.ForgotToken = RandomToken(16)
		u.ForgotTokenExpiry = time.Now().Add(time.Hour * 24)
		u.Save()
	}
	response.ForgotToken = u.ForgotToken

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
