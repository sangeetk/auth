package service

import (
	"context"
	"time"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/user"
)

// Recover - Resets the password
func (Auth) Recover(_ context.Context, req api.RecoverRequest) (api.RecoverResponse, error) {
	var response api.RecoverResponse

	if req.Username == "" {
		response.Err = ErrorInvalidRequest.Error()
		return response, nil
	}

	u, err := user.Read(req.Username)
	if err != nil || u.Confirmed != true {
		response.Err = ErrorNotFound.Error()
		return response, nil
	}

	if u.RecoverToken == "" || u.RecoverTokenExpiry.Unix() < time.Now().Unix() {
		u.RecoverToken = RandomToken(16)
		u.RecoverTokenExpiry = time.Now().Add(time.Hour * 24)
		u.Save()
	}
	response.RecoverToken = u.RecoverToken

	return response, nil
}
