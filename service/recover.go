package service

import (
	"context"
	"time"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/user"
	"golang.org/x/crypto/bcrypt"
)

// Recover - Resets the password
func (Auth) Recover(_ context.Context, req api.RecoverRequest) (api.RecoverResponse, error) {
	var response api.RecoverResponse

	if req.Cmd == api.RecoveryToken {
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

	} else if req.Cmd == api.ResetPassword && req.RecoverToken != "" && req.Password != "" {
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

	} else {
		response.Err = ErrorInvalidRequest.Error()
	}
	return response, nil
}
