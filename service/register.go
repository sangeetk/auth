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

// Register - Register a new User
func (Auth) Register(ctx context.Context, req api.RegisterRequest) (api.RegisterResponse, error) {
	var response = api.RegisterResponse{}

	if req.Email == "" || req.Password == "" {
		response.Err = ErrorInvalidRequest.Error()
		return response, nil
	}

	// Use email as username if empty
	if req.Username == "" {
		req.Username = req.Email
	}

	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 11)
	if err != nil {
		log.Println("Bcrypt error:", err.Error())
	}

	var u = &user.User{
		Username:      req.Username,
		Name:          req.Name,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		Email:         req.Email,
		Password:      PasswordHash,
		Birthday:      req.Birthday,
		InitialDomain: req.Domain,
		Confirmed:     false,
	}

	u.Roles = make(map[string][]string)
	u.Roles[req.Domain] = req.Roles

	u.Address = req.Address
	u.Profile = req.Profile

	if err := u.Create(); err != nil {
		response.Err = ErrorAlreadyRegistered.Error()
		return response, nil
	}

	// Create the Confirmation token
	response.ConfirmToken, err = token.NewToken(u, req.Domain, token.ConfirmTokenValidity)
	if err != nil {
		response.Err = err.Error()
		return response, nil
	}

	// Create the Update token
	response.UpdateToken, err = token.NewToken(u, req.Domain, token.UpdateTokenValidity)
	if err != nil {
		response.Err = err.Error()
		return response, nil
	}

	response.Username = u.Username
	response.FirstName = u.FirstName
	response.LastName = u.LastName
	response.Email = u.Email
	response.Domain = req.Domain
	response.Roles = u.GetRoles(req.Domain)

	return response, nil
}

// MakeRegisterEndpoint -
func MakeRegisterEndpoint(svc Auth) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.RegisterRequest)
		return svc.Register(ctx, req)
	}
}

// DecodeRegisterRequest -
func DecodeRegisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
