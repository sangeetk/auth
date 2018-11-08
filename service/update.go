package service

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/user"
	"github.com/urantiatech/kit/endpoint"
	"golang.org/x/crypto/bcrypt"
)

// Update - Updates the user
func (Auth) Update(ctx context.Context, req api.UpdateRequest) (api.UpdateResponse, error) {
	var response = api.UpdateResponse{}
	var u *user.User
	var err error

	// UpdateToken takes priority and ignores Confirmed status
	if req.UpdateToken != "" {
		u, err = ParseToken(req.UpdateToken)
	} else {
		u, err = ParseToken(req.AccessToken)
	}
	if err == ErrorInvalidToken {
		response.Err = err.Error()
		return response, nil
	}

	// Ignore u.Confirmed if UpdateToken is provided
	u, err = user.Read(u.Username)
	if (err != nil) || (u.Confirmed == false && req.UpdateToken == "") {
		response.Err = ErrorNotFound.Error()
		return response, nil
	}

	// Don't check password if UpdateToken is provided
	if req.UpdateToken == "" && bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)) != nil {
		response.Err = ErrorInvalidPassword.Error()
		return response, nil
	}

	// Update user fields
	if req.Name != "" {
		u.Name = req.Name
	}
	if req.FirstName != "" {
		u.FirstName = req.FirstName
	}
	if req.LastName != "" {
		u.LastName = req.LastName
	}
	/*
		if req.Email != "" {
			u.Email = req.Email
		}
	*/
	if !req.Birthday.IsZero() {
		u.Birthday = req.Birthday
	}
	if req.NewPassword != "" {
		PasswordHash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 11)
		if err != nil {
			log.Println("Bcrypt error:", err.Error())
		}
		u.Password = PasswordHash
	}

	// Address information
	if req.Address.AddressType != "" {
		u.Address.AddressType = req.Address.AddressType
	}
	if req.Address.Address1 != "" {
		u.Address.Address1 = req.Address.Address1
	}
	if req.Address.Address2 != "" {
		u.Address.Address2 = req.Address.Address2
	}
	if req.Address.City != "" {
		u.Address.City = req.Address.City
	}
	if req.Address.State != "" {
		u.Address.State = req.Address.State
	}
	if req.Address.Country != "" {
		u.Address.Country = req.Address.Country
	}
	if req.Address.Zip != "" {
		u.Address.Zip = req.Address.Zip
	}

	// Update Roles
	if req.Domain != "" && len(req.Roles) > 0 {
		u.Roles[req.Domain] = req.Roles
	}

	// Profile information
	for k, v := range req.Profile {
		if v == "" {
			delete(u.Profile, k)
		} else {
			u.Profile[k] = v
		}
	}

	u.Save()
	response.UpdateToken = req.UpdateToken
	response.Username = u.Username
	response.FirstName = u.FirstName
	response.LastName = u.LastName
	response.Email = u.Email

	return response, nil
}

// MakeUpdateEndpoint -
func MakeUpdateEndpoint(svc Auth) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.UpdateRequest)
		return svc.Update(ctx, req)
	}
}

// DecodeUpdateRequest -
func DecodeUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
