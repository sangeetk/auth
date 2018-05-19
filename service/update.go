package service

import (
	"context"
	"log"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/model"
	"golang.org/x/crypto/bcrypt"
)

func (Auth) Update(_ context.Context, req api.UpdateRequest) (api.UpdateResponse, error) {
	var response = api.UpdateResponse{}
	var user model.User

	tx := DB.Begin()

	if req.Uid == 0 || tx.Where(&model.User{ID: req.Uid}).First(&user).RecordNotFound() {
		tx.Rollback()
		response.Err = NotFound.Error()
		return response, nil
	}

	// User table
	u := make(map[string]interface{})
	if req.Fname != "" {
		u["fname"] = req.Fname
	}
	if req.Lname != "" {
		u["lname"] = req.Lname
	}
	if req.Email != "" {
		u["email"] = req.Email
	}
	if !req.Birthday.IsZero() {
		u["birthday"] = req.Birthday
	}
	if req.Password != "" {
		PasswordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 11)
		if err != nil {
			log.Println("Bcrypt error:", err.Error())
		}
		u["password"] = PasswordHash
	}
	tx.Model(&user).Where(&model.User{ID: req.Uid}).Updates(u)

	// Address Table
	var address = model.Address{Uid: req.Uid}
	a := make(map[string]interface{})
	if req.Address1 != "" {
		a["address1"] = req.Address1
	}
	if req.Address2 != "" {
		a["address2"] = req.Address2
	}
	if req.City != "" {
		a["city"] = req.City
	}
	if req.State != "" {
		a["state"] = req.State
	}
	if req.Country != "" {
		a["country"] = req.Country
	}
	if req.Zip != "" {
		a["zip"] = req.Zip
	}

	tx.Model(&address).Where(&model.Address{Uid: req.Uid}).Updates(a)

	// Profile Table
	var profile = model.Profile{}
	p := make(map[string]interface{})
	if req.Profession != "" {
		p["profession"] = req.Profession
	}
	if req.Introduction != "" {
		p["introduction"] = req.Introduction
	}
	tx.Model(&profile).Where(&model.Profile{Uid: req.Uid}).Updates(p)

	// Role Table

	tx.Commit()

	return response, nil
}
