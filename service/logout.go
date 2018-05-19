package service

import (
	"context"
	"log"

	"git.urantiatech.com/auth/auth/api"
	"github.com/patrickmn/go-cache"
)

func (Auth) Logout(_ context.Context, req api.LogoutRequest) (api.LogoutResponse, error) {
	var response api.LogoutResponse

	// Ignore if it is an invalid token
	_, err := ParseToken(req.AccessToken)
	if err == InvalidToken {
		response.Err = err.Error()
		return response, nil
	}

	if err != ExpiredToken {
		// Blacklist the token
		err := BlacklistTokens.Add(req.AccessToken, nil, cache.DefaultExpiration)
		if err != nil {
			log.Println(err.Error())
		}
	}

	return response, nil
}
