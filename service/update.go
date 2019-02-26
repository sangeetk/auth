package service

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/token"
	"git.urantiatech.com/auth/auth/user"
	"github.com/urantiatech/kit/endpoint"
	"golang.org/x/crypto/bcrypt"
)

// Update - Updates the user
func (Auth) Update(ctx context.Context, req api.UpdateRequest) (api.UpdateResponse, error) {
	var response = api.UpdateResponse{}
	var t *token.Token
	var err error

	// UpdateToken takes priority and ignores Confirmed status
	if req.UpdateToken != "" {
		t, err = token.ParseToken(req.UpdateToken)
	} else {
		t, err = token.ParseToken(req.AccessToken)
	}
	if err == token.ErrorInvalidToken {
		response.Err = err.Error()
		return response, nil
	}

	if req.AccessToken != "" {
		// Check against Blacklisted tokens
		if _, found := token.BlacklistAccessTokens.Get(req.AccessToken); found {
			response.Err = token.ErrorInvalidToken.Error()
			return response, nil
		}

		// Blacklist the exising Access Token
		err = token.BlacklistAccessTokens.Add(req.AccessToken, nil, token.AccessTokenValidity)
		if err != nil {
			response.Err = token.ErrorInvalidToken.Error()
			return response, nil
		}
	}

	// Ignore u.Confirmed if UpdateToken is provided
	u, err := user.Read(t.Username)
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
	if len(req.Roles) > 0 {
		u.Roles[t.Domain] = req.Roles
	}

	// Profile information
	for k, v := range req.Profile {
		if u.Profile == nil {
			u.Profile = make(map[string]string)
		}
		if v == "" {
			delete(u.Profile, k)
		} else {
			u.Profile[k] = v
		}
	}

	u.Save()

	// Create new Access Token
	response.NewAccessToken, err = token.NewToken(u, t.Domain, token.AccessTokenValidity)
	if err != nil {
		response.Err = err.Error()
		return response, nil
	}

	response.UpdateToken = req.UpdateToken
	response.Username = u.Username
	response.FirstName = u.FirstName
	response.LastName = u.LastName
	response.Email = u.Email
	response.Domain = t.Domain
	response.Roles = u.GetRoles(t.Domain)

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
