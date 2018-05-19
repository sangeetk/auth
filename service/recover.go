package service

import (
	"context"
	"time"

	"github.com/urantiatech/microservices/auth/api"
	"github.com/urantiatech/microservices/auth/model"
	"golang.org/x/crypto/bcrypt"
)

func (Auth) Recover(_ context.Context, req api.RecoverRequest) (api.RecoverResponse, error) {
	var response api.RecoverResponse
	var user model.User

	if req.Cmd == api.RecoveryToken {
		DB.Where("email = ?", req.Email).Where("confirmed = ?", true).First(&user)
		if user.ID == 0 {
			response.Err = NotFound.Error()
			return response, nil
		}

		if user.RecoverToken == "" || user.RecoverTokenExpiry.Unix() < time.Now().Unix() {
			user.RecoverToken = RandomToken(16)
			user.RecoverTokenExpiry = time.Now().Add(time.Hour * 24)
			DB.Save(&user)
		}
		response.RecoverToken = user.RecoverToken

	} else if req.Cmd == api.ResetPassword && req.RecoverToken != "" && req.Password != "" {
		DB.Where("recover_token = ?", req.RecoverToken).Where("confirmed = ?", true).First(&user)
		if user.ID == 0 || user.RecoverTokenExpiry.Unix() < time.Now().Unix() {
			response.Err = ExpiredToken.Error()
			return response, nil
		}

		PasswordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 11)
		if err != nil {
			response.Err = UnknownError.Error()
			return response, nil
		}
		user.RecoverToken = ""
		user.RecoverTokenExpiry = time.Unix(0, 0)
		user.Password = PasswordHash
		DB.Save(&user)

	} else {
		response.Err = InvalidRequest.Error()
	}
	return response, nil
}
