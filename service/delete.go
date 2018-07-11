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

// Delete - Deletes the user
func (Auth) Delete(_ context.Context, req api.DeleteRequest) (api.DeleteResponse, error) {
	var response = api.DeleteResponse{}

	// Validate the token and get user info
	u, err := ParseToken(req.AccessToken)
	if err == ErrorInvalidToken {
		response.Err = err.Error()
		return response, nil
	}

	u, err = user.Read(u.Username)
	if err != nil || u.Confirmed != true {
		response.Err = ErrorNotFound.Error()
		return response, nil
	}

	// Check password
	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)) != nil {
		response.Err = ErrorInvalidPassword.Error()
		return response, nil
	}

	u.Confirmed = false
	u.DeletedAt = time.Now()
	u.Save()

	return response, nil
}

// MakeDeleteEndpoint -
func MakeDeleteEndpoint(svc Auth) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.DeleteRequest)
		return svc.Delete(ctx, req)
	}
}

// DecodeDeleteRequest -
func DecodeDeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.DeleteRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
