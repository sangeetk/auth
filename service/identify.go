package service

import (
	"context"
	"encoding/json"
	"net/http"

	"git.urantiatech.com/auth/auth/api"
	"github.com/urantiatech/kit/endpoint"
)

// Identify - Identify the user based on the AccessToken
func (Auth) Identify(_ context.Context, req api.IdentifyRequest) (api.IdentifyResponse, error) {
	var response api.IdentifyResponse

	user, err := ParseToken(req.AccessToken)
	if err != nil {
		response.Err = err.Error()
		return response, nil
	}

	// Check by looking for Blacklisted tokens in Cache
	if _, found := BlacklistTokens.Get(req.AccessToken); found {
		response.Err = ErrorInvalidToken.Error()
		return response, nil
	}

	// Send the user details
	response.Username = user.Username
	response.FirstName = user.FirstName
	response.LastName = user.LastName
	response.Email = user.Email

	return response, nil
}

// MakeIdentifyEndpoint -
func MakeIdentifyEndpoint(svc Auth) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.IdentifyRequest)
		return svc.Identify(ctx, req)
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
