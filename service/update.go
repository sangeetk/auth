package service

import (
	"context"
	"log"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/user"
	"golang.org/x/crypto/bcrypt"
)

// Update - Updates the user
func (Auth) Update(_ context.Context, req api.UpdateRequest) (api.UpdateResponse, error) {
	var response = api.UpdateResponse{}

	u, err := user.Read(req.Username)
	if err != nil || u.Confirmed != true {
		response.Err = ErrorNotFound.Error()
		return response, err
	}

	// Update user fields
	if req.FirstName != "" {
		u.FirstName = req.FirstName
	}
	if req.LastName != "" {
		u.LastName = req.LastName
	}
	if req.Email != "" {
		u.Email = req.Email
	}
	if !req.Birthday.IsZero() {
		u.Birthday = req.Birthday
	}
	if req.Password != "" {
		PasswordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 11)
		if err != nil {
			log.Println("Bcrypt error:", err.Error())
		}
		u.Password = PasswordHash
	}

	// Address information
	if req.Address1 != "" {
		u.Address.Address1 = req.Address1
	}
	if req.Address2 != "" {
		u.Address.Address2 = req.Address2
	}
	if req.City != "" {
		u.Address.City = req.City
	}
	if req.State != "" {
		u.Address.State = req.State
	}
	if req.Country != "" {
		u.Address.Country = req.Country
	}
	if req.Zip != "" {
		u.Address.Zip = req.Zip
	}

	// Profile information
	if req.Profession != "" {
		u.Profile.Profession = req.Profession
	}
	if req.Introduction != "" {
		u.Profile.Introduction = req.Introduction
	}

	u.Save()

	return response, nil
}
