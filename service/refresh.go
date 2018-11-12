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

// Refresh - Generates a new token and extends existing session
func (Auth) Refresh(ctx context.Context, req api.RefreshRequest) (api.RefreshResponse, error) {
	var response api.RefreshResponse

	// Validate the refresh token and get user info
	t, err := token.ParseToken(req.RefreshToken)
	if err == token.ErrorInvalidToken {
		response.Err = token.ErrorInvalidToken.Error()
		return response, nil
	}

	// Check against Blacklisted refresh tokens
	if _, found := token.BlacklistRefreshTokens.Get(req.RefreshToken); found {
		response.Err = token.ErrorInvalidToken.Error()
		return response, nil
	}

	u, err := user.Read(t.Username)
	if err != nil || u.Confirmed != true {
		response.Err = ErrorNotFound.Error()
		return response, nil
	}

	// Create new Access Token
	response.NewAccessToken, err = token.NewToken(u, t.Domain, token.AccessTokenValidity)
	if err != nil {
		response.Err = err.Error()
		return response, nil
	}

	// Create new Refresh Token
	response.NewRefreshToken, err = token.NewToken(u, t.Domain, token.RefreshTokenValidity)
	if err != nil {
		response.Err = err.Error()
		return response, nil
	}

	response.Username = u.Username
	response.FirstName = u.FirstName
	response.LastName = u.LastName
	response.Email = u.Email
	response.Domain = t.Domain
	response.Roles = u.GetRoles(t.Domain)

	return response, nil
}

// MakeRefreshEndpoint -
func MakeRefreshEndpoint(svc Auth) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.RefreshRequest)
		return svc.Refresh(ctx, req)
	}
}

// DecodeRefreshRequest -
func DecodeRefreshRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.RefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
