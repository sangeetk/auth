package service

import (
	"context"
	"time"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/user"
	"golang.org/x/crypto/bcrypt"
)

// Delete - Deletes the user
func (Auth) Delete(_ context.Context, req api.DeleteRequest) (api.DeleteResponse, error) {
	var response = api.DeleteResponse{}

	// Validate the token and get user info
	u, err := ParseToken(req.AccessToken)
	if err == ErrorInvalidToken {
		response.Err = err.Error()
		return response, nil
	}

	u, err = user.Read(u.Username)
	if err != nil || u.Confirmed != true {
		response.Err = ErrorNotFound.Error()
		return response, nil
	}

	// Check password
	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)) != nil {
		response.Err = ErrorInvalidPassword.Error()
		return response, nil
	}

	u.Confirmed = false
	u.DeletedAt = time.Now()
	u.Save()

	return response, nil
}
