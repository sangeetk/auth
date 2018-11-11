package service

import (
	"context"
	"encoding/json"
	"net/http"

	"git.urantiatech.com/auth/auth/api"
	"github.com/urantiatech/kit/endpoint"
)

// Logout - Logouts the current user
func (Auth) Logout(ctx context.Context, req api.LogoutRequest) (api.LogoutResponse, error) {
	var response api.LogoutResponse
	var err error

	// Ignore if it is an invalid token
	_, err = ParseToken(req.AccessToken)
	if err != ErrorInvalidToken && err != ErrorExpiredToken {
		// Blacklist the access token
		err = BlacklistAccessTokens.Add(req.AccessToken, nil, AccessTokenValidity)
	}
	if err != nil {
		response.Err = ErrorInvalidToken.Error()
	}

	// Ignore if it is an invalid refresh token
	_, err = ParseToken(req.RefreshToken)
	if err != ErrorInvalidToken && err != ErrorExpiredToken {
		// Blacklist the token
		err = BlacklistRefreshTokens.Add(req.RefreshToken, nil, RefreshTokenValidity)
	}
	if err != nil {
		response.Err = ErrorInvalidToken.Error()
	}

	return response, nil
}

// MakeLogoutEndpoint -
func MakeLogoutEndpoint(svc Auth) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.LogoutRequest)
		return svc.Logout(ctx, req)
	}
}

// DecodeLogoutRequest -
func DecodeLogoutRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.LogoutRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
