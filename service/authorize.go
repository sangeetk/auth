package service

import (
	"context"
	"encoding/json"
	"net/http"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/token"
	"git.urantiatech.com/auth/auth/user"
	"github.com/urantiatech/kit/endpoint"
)

// Authorize - Sends Auth Token if user has "role"
func (Auth) Authorize(ctx context.Context, req api.AuthorizeRequest) (api.AuthorizeResponse, error) {
	var response = api.AuthorizeResponse{Authorize: false}

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

	// Read user details
	u, err := user.Read(t.Username)
	if err != nil || !u.Confirmed || !u.DeletedAt.IsZero() {
		response.Err = token.ErrorInvalidToken.Error()
		return response, nil
	}

	for _, r := range u.GetRoles(t.Domain) {
		if r == req.Role {
			response.Authorize = true
		}
	}

	return response, nil
}

// MakeAuthorizeEndpoint -
func MakeAuthorizeEndpoint(svc Auth) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.AuthorizeRequest)
		return svc.Authorize(ctx, req)
	}
}

// DecodeAuthorizeRequest -
func DecodeAuthorizeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.AuthorizeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
