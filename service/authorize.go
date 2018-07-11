package service

import (
	"context"
	"encoding/json"
	"net/http"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/user"
	"github.com/urantiatech/kit/endpoint"
)

// Authorize - Sends Auth Token if user has "role"
func (Auth) Authorize(_ context.Context, req api.AuthorizationRequest) (api.AuthorizationResponse, error) {
	var response = api.AuthorizationResponse{Authorize: false}
	var u *user.User

	// Validate the token and get user info
	u, err := ParseToken(req.AccessToken)
	if err == ErrorInvalidToken {
		response.Err = err.Error()
		return response, nil
	}

	// Check by looking for Blacklisted tokens in Cache
	if _, found := BlacklistTokens.Get(req.AccessToken); found {
		response.Err = ErrorInvalidToken.Error()
		return response, nil
	}

	u, err = user.Read(u.Username)
	if err != nil || u.Confirmed != true {
		response.Err = ErrorNotFound.Error()
		return response, nil
	}

	roles, ok := u.Roles[req.Domain]
	if !ok {
		return response, nil
	}
	for _, r := range roles {
		if r == req.Role {
			response.Authorize = true
		}
	}

	return response, nil
}

// MakeAuthorizationEndpoint -
func MakeAuthorizationEndpoint(svc Auth) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.AuthorizationRequest)
		return svc.Authorize(ctx, req)
	}
}

// DecodeAuthorizationRequest -
func DecodeAuthorizationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.AuthorizationRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
