package service

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/token"
	"git.urantiatech.com/auth/auth/user"
	"github.com/urantiatech/kit/endpoint"
	"golang.org/x/crypto/bcrypt"
)

// Delete - Deletes the user
func (Auth) Delete(ctx context.Context, req *api.DeleteRequest) (*api.DeleteResponse, error) {
	var response = &api.DeleteResponse{}

	// Validate the token and get user info
	t, err := token.ParseToken(req.AccessToken)
	if err == token.ErrorInvalidToken {
		response.Err = err.Error()
		return response, nil
	}

	// Check against Blacklisted tokens
	if _, found := token.BlacklistAccessTokens.Get(req.AccessToken); found {
		response.Err = token.ErrorInvalidToken.Error()
		return response, nil
	}

	u, err := user.Read(t.Username)
	if err != nil || !u.Confirmed || !u.DeletedAt.IsZero() {
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

	response.Username = u.Username
	response.FirstName = u.FirstName
	response.LastName = u.LastName
	response.Email = u.Email
	response.Domain = t.Domain
	response.Roles = u.GetRoles(t.Domain)

	return response, nil
}

// MakeDeleteEndpoint -
func MakeDeleteEndpoint(svc AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.DeleteRequest)
		return svc.Delete(ctx, &req)
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
