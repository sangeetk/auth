package service

import (
	"context"
	"time"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/user"
	"golang.org/x/crypto/bcrypt"
)

// Reset - Resets the password
func (Auth) Reset(_ context.Context, req api.ResetRequest) (api.ResetResponse, error) {
	var response api.ResetResponse

	if req.RecoverToken == "" || req.Password == "" {
		response.Err = ErrorInvalidRequest.Error()
		return response, nil
	}

	u, err := user.Read(req.Username)
	if err != nil || req.RecoverToken != u.RecoverToken || u.RecoverTokenExpiry.Unix() < time.Now().Unix() {
		response.Err = ErrorExpiredToken.Error()
		return response, nil
	}

	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 11)
	if err != nil {
		response.Err = ErrorUnknown.Error()
		return response, nil
	}
	u.RecoverToken = ""
	u.RecoverTokenExpiry = time.Unix(0, 0)
	u.Password = PasswordHash
	u.Save()

	return response, nil
}
