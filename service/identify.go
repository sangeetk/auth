package service

import (
	"context"

	"git.urantiatech.com/auth/auth/api"
)

// Identify - Identify the user based on the AccessToken
func (Auth) Identify(_ context.Context, req api.IdentifyRequest) (api.IdentifyResponse, error) {
	var response api.IdentifyResponse

	user, err := ParseToken(req.AccessToken)
	if err != nil {
		response.Err = err.Error()
		return response, nil
	}

	// Check by looking for Blacklisted tokens in Redis Cache
	if _, found := BlacklistTokens.Get(req.AccessToken); found {
		response.Err = ErrorInvalidToken.Error()
		return response, nil
	}

	// Send the user details
	response.FirstName = user.FirstName
	response.LastName = user.LastName
	response.Email = user.Email
	//response.Roles = user.Roles

	return response, nil
}
