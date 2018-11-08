package service

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"git.urantiatech.com/auth/auth/api"
	"github.com/patrickmn/go-cache"
	"github.com/urantiatech/kit/endpoint"
)

// Logout - Logouts the current user
func (Auth) Logout(ctx context.Context, req api.LogoutRequest) (api.LogoutResponse, error) {
	var response api.LogoutResponse

	// Ignore if it is an invalid token
	_, err := ParseToken(req.AccessToken)
	if err == ErrorInvalidToken {
		response.Err = err.Error()
		return response, nil
	}

	if err != ErrorExpiredToken {
		// Blacklist the token
		err := BlacklistTokens.Add(req.AccessToken, nil, cache.DefaultExpiration)
		if err != nil {
			log.Println(err.Error())
		}
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
