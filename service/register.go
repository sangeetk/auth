package service

import (
	"context"
	"log"
	"time"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/user"
	"golang.org/x/crypto/bcrypt"
)

// Register - Register a new User
func (Auth) Register(_ context.Context, req api.RegisterRequest) (api.RegisterResponse, error) {
	var response = api.RegisterResponse{}

	if req.Username == "" || req.Email == "" || req.Password == "" {
		response.Err = ErrorInvalidRequest.Error()
		return response, nil
	}

	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 11)
	if err != nil {
		log.Println("Bcrypt error:", err.Error())
	}

	// Generate random confirm token
	ConfirmToken := RandomToken(16)

	// Use email as username if empty
	if req.Username == "" {
		req.Username = req.Email
	}

	var u = user.User{
		Username:      req.Username,
		Name:          req.Name,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		Email:         req.Email,
		Password:      PasswordHash,
		Birthday:      req.Birthday,
		InitialDomain: req.Domain,
		ConfirmToken:  ConfirmToken,
		Confirmed:     false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	u.Roles = make(map[string][]string)
	u.Roles[req.Domain] = req.Roles

	u.Address = user.Address{
		AddressType: "",
		Address1:    req.Address1,
		Address2:    req.Address2,
		City:        req.City,
		State:       req.State,
		Country:     req.Country,
		Zip:         req.Zip,
	}

	u.Profile = user.Profile{
		Profession:   req.Profession,
		Introduction: req.Introduction,
	}

	if err := u.Create(); err != nil {
		response.Err = ErrorAlreadyRegistered.Error()
		return response, nil
	}

	response.ConfirmToken = ConfirmToken
	return response, nil
}
