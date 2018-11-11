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
func (Auth) Authorize(ctx context.Context, req api.AuthorizeRequest) (api.AuthorizeResponse, error) {
	var response = api.AuthorizeResponse{Authorize: false}
	var u *user.User

	// Validate the token and get user info
	u, err := ParseToken(req.AccessToken)
	if err == ErrorInvalidToken {
		response.Err = err.Error()
		return response, nil
	}

	// Check by looking for Blacklisted access tokens in Cache
	if _, found := BlacklistAccessTokens.Get(req.AccessToken); found {
		response.Err = ErrorInvalidToken.Error()
		return response, nil
	}

	// Using InitialDomain as temp variable
	roles, ok := u.Roles[u.InitialDomain]
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
