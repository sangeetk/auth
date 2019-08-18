package service

import (
	"context"
	"encoding/json"
	"net/http"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/token"
	"github.com/go-kit/kit/endpoint"
)

// Identify - Identify the user based on the AccessToken
func (a Auth) Identify(ctx context.Context, req *api.IdentifyRequest) (*api.IdentifyResponse, error) {
	var response = &api.IdentifyResponse{}

	t, err := token.ParseToken(req.AccessToken)
	if err != nil {
		response.Err = err.Error()
		return response, nil
	}

	// Check against Blacklisted tokens
	if _, found := token.BlacklistAccessTokens.Get(req.AccessToken); found {
		response.Err = token.ErrorInvalidToken.Error()
		return response, nil
	}

	// Send the user details
	response.Username = t.Username
	response.FirstName = t.FirstName
	response.LastName = t.LastName
	response.Email = t.Email
	response.Domain = t.Domain
	response.Roles = t.Roles

	return response, nil
}

// MakeIdentifyEndpoint -
func MakeIdentifyEndpoint(svc AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.IdentifyRequest)
		return svc.Identify(ctx, &req)
	}
}

// DecodeIdentifyRequest -
func DecodeIdentifyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.IdentifyRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
