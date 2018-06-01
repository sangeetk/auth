package service

import (
	"context"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/user"
)

// Confirm - Activates the user account after confirmation
func (Auth) Confirm(_ context.Context, req api.ConfirmRequest) (api.ConfirmResponse, error) {
	var response api.ConfirmResponse

	user, err := user.Read(req.Username)
	if err != nil || user.ConfirmToken != req.ConfirmToken || user.Confirmed != false {
		response.Err = ErrorInvalidToken.Error()
		return response, nil
	}

	user.ConfirmToken = ""
	user.Confirmed = true

	user.Save()
	return response, nil
}
