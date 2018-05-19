package service

import (
	"context"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/model"
)

func (Auth) Confirm(_ context.Context, req api.ConfirmRequest) (api.ConfirmResponse, error) {
	var response api.ConfirmResponse

	user := model.User{ConfirmToken: req.ConfirmToken}

	tx := DB.Begin()
	tx.Where("confirm_token = ?", user.ConfirmToken).Where("confirmed = ?", false).First(&user)

	if user.ID == 0 {
		response.Err = InvalidToken.Error()
		tx.Rollback()
		return response, nil
	}

	user.ConfirmToken = ""
	user.Confirmed = true

	tx.Save(&user)
	tx.Commit()
	return response, nil
}
