package service

import (
	"context"
	"encoding/json"
	"net/http"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/token"
	"git.urantiatech.com/auth/auth/user"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/crypto/bcrypt"
)

// Login - Log in the user after credentials are successfully authenticated
func (Auth) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {
	var response = &api.LoginResponse{}

	// Get user details
	u, err := user.Read(req.Username)
	if err != nil || !u.Confirmed || !u.DeletedAt.IsZero() {
		response.Err = ErrorInvalidLogin.Error()
		return response, nil
	}

	// Check password
	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)) != nil {
		response.Err = ErrorInvalidLogin.Error()
		return response, nil
	}

	// Create new Access Token
	if req.RememberMe {
		response.AccessToken, err = token.NewToken(u, req.Domain, token.RememberMeAccessTokenValidity)
	} else {
		response.AccessToken, err = token.NewToken(u, req.Domain, token.AccessTokenValidity)
	}
	if err != nil {
		response.Err = err.Error()
		return response, nil
	}

	// Create new Refresh Token
	if req.RememberMe {
		response.RefreshToken, err = token.NewToken(u, req.Domain, token.RememberMeRefreshTokenValidity)
	} else {
		response.RefreshToken, err = token.NewToken(u, req.Domain, token.RefreshTokenValidity)
	}
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

// MakeLoginEndpoint -
func MakeLoginEndpoint(svc AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.LoginRequest)
		return svc.Login(ctx, &req)
	}
}

// DecodeLoginRequest -
func DecodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
