package service

import (
	"context"
	"encoding/json"
	"net/http"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/token"
	"github.com/urantiatech/kit/endpoint"
)

// Logout - Logouts the current user
func (Auth) Logout(ctx context.Context, req *api.LogoutRequest) (*api.LogoutResponse, error) {
	var response = &api.LogoutResponse{}
	var err error

	// Ignore if it is an invalid token
	_, err = token.ParseToken(req.AccessToken)
	if err != token.ErrorInvalidToken && err != token.ErrorExpiredToken {
		// Blacklist the access token
		err = token.BlacklistAccessTokens.Add(req.AccessToken, nil, token.AccessTokenValidity)
	}
	if err != nil {
		response.Err = token.ErrorInvalidToken.Error()
	}

	// Ignore if it is an invalid refresh token
	_, err = token.ParseToken(req.RefreshToken)
	if err != token.ErrorInvalidToken && err != token.ErrorExpiredToken {
		// Blacklist the refresh token
		err = token.BlacklistRefreshTokens.Add(req.RefreshToken, nil, token.RefreshTokenValidity)
	}
	if err != nil {
		response.Err = token.ErrorInvalidToken.Error()
	}

	return response, nil
}

// MakeLogoutEndpoint -
func MakeLogoutEndpoint(svc AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.LogoutRequest)
		return svc.Logout(ctx, &req)
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
