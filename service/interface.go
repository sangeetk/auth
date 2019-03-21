package service

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"git.urantiatech.com/auth/auth/api"
)

// AuthService - Authentication and Authorization Microservice
type AuthService interface {
	Authorize(context.Context, *api.AuthorizeRequest) (*api.AuthorizeResponse, error)
	Confirm(context.Context, *api.ConfirmRequest) (*api.ConfirmResponse, error)
	Delete(context.Context, *api.DeleteRequest) (*api.DeleteResponse, error)
	Forgot(context.Context, *api.ForgotRequest) (*api.ForgotResponse, error)
	Identify(context.Context, *api.IdentifyRequest) (*api.IdentifyResponse, error)
	Login(context.Context, *api.LoginRequest) (*api.LoginResponse, error)
	Logout(context.Context, *api.LogoutRequest) (*api.LogoutResponse, error)
	Profile(context.Context, *api.ProfileRequest) (*api.ProfileResponse, error)
	Refresh(context.Context, *api.RefreshRequest) (*api.RefreshResponse, error)
	Register(context.Context, *api.RegisterRequest) (*api.RegisterResponse, error)
	Reset(context.Context, *api.ResetRequest) (*api.ResetResponse, error)
	Update(context.Context, *api.UpdateRequest) (*api.UpdateResponse, error)
}

// Auth - Wrapper for AuthService Interface
type Auth struct{}

// ErrorNotFound - User Not Found
var ErrorNotFound = errors.New("Not Found")

// ErrorAlreadyRegistered - User Already Registered
var ErrorAlreadyRegistered = errors.New("Already Registered")

// ErrorInvalidLogin - Invalid Login
var ErrorInvalidLogin = errors.New("Invalid Login")

// ErrorInvalidPassword - Invalid Password
var ErrorInvalidPassword = errors.New("Invalid Password")

// ErrorInvalidRequest - Invalid Request
var ErrorInvalidRequest = errors.New("Invalid Request")

// ErrorUnknown - Unknown Error
var ErrorUnknown = errors.New("Unknown Error")

// AuthMiddleware - Service Middleware is a chainable behavior modifier for AuthService.
type AuthMiddleware func(AuthService) AuthService

// EncodeResponse -
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
