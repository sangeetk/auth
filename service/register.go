package service

import (
	"context"
	"log"

	"github.com/urantiatech/microservices/auth/api"
	"github.com/urantiatech/microservices/auth/model"
	"golang.org/x/crypto/bcrypt"
)

func (Auth) Register(_ context.Context, req api.RegisterRequest) (api.RegisterResponse, error) {
	var response = api.RegisterResponse{}

	if req.Email == "" || req.Password == "" {
		response.Err = InvalidRequest.Error()
		return response, nil
	}

	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 11)
	if err != nil {
		log.Println("Bcrypt error:", err.Error())
	}

	// Generate random confirm token
	ConfirmToken := RandomToken(16)

	user := model.User{Email: req.Email}
	tx := DB.Begin()

	tx.Where("email = ?", user.Email).First(&user)
	if user.ID != 0 {
		response.Err = AlreadyRegistered.Error()
		tx.Rollback()
		return response, nil
	}

	// Use email as username if empty
	if req.Username == "" {
		req.Username = req.Email
	}

	user = model.User{
		Domain:       req.Domain,
		Username:     req.Username,
		Fname:        req.Fname,
		Lname:        req.Lname,
		Email:        req.Email,
		Password:     PasswordHash,
		Birthday:     req.Birthday,
		ConfirmToken: ConfirmToken,
		Confirmed:    false,
	}
	DB.Create(&user)

	address := model.Address{
		Uid:         user.ID,
		AddressType: "",
		Address1:    req.Address1,
		Address2:    req.Address2,
		City:        req.City,
		State:       req.State,
		Country:     req.Country,
		Zip:         req.Zip,
	}
	DB.Create(&address)

	profile := model.Profile{Uid: user.ID, Profession: req.Profession, Introduction: req.Introduction}
	if err != nil {
		log.Println("Error: ", err.Error())
	}
	DB.Create(&profile)

	tx.Commit()

	response.ConfirmToken = ConfirmToken
	return response, nil
}
