package service

import (
	"context"

	"github.com/urantiatech/microservices/auth/api"
)

func (Auth) Identify(_ context.Context, req api.IdentifyRequest) (api.IdentifyResponse, error) {
	var response api.IdentifyResponse

	user, err := ParseToken(req.AccessToken)
	if err != nil {
		response.Err = err.Error()
		return response, nil
	}

	// Check by looking for Blacklisted tokens in Redis Cache
	if _, found := BlacklistTokens.Get(req.AccessToken); found {
		response.Err = InvalidToken.Error()
		return response, nil
	}

	// Send the user details
	response.Uid = user.ID
	response.Fname = user.Fname
	response.Lname = user.Lname
	response.Email = user.Email
	//response.Roles = user.Roles

	return response, nil
}
