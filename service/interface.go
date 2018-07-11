package service

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"git.urantiatech.com/auth/auth/api"
	"github.com/patrickmn/go-cache"
)

// SigningKey - JWT Signing Key
var SigningKey []byte

// AccessTokenValidity - JWT Access Token Validity
var AccessTokenValidity time.Duration

// RefreshTokenValidity - JWT Refresh Token Validity
var RefreshTokenValidity time.Duration

// BlacklistTokens - Cache to store invalid tokens
var BlacklistTokens *cache.Cache

// AuthService - Authentication and Authorization Microservice
type AuthService interface {
	Register(context.Context, api.RegisterRequest) (api.RegisterResponse, error)
	Login(context.Context, api.LoginRequest) (api.LoginResponse, error)
	Logout(context.Context, api.LogoutRequest) (api.LogoutResponse, error)
	Identify(context.Context, api.IdentifyRequest) (api.IdentifyResponse, error)
	Profile(context.Context, api.ProfileRequest) (api.ProfileResponse, error)
	Refresh(context.Context, api.RefreshRequest) (api.RefreshResponse, error)
	Confirm(context.Context, api.ConfirmRequest) (api.ConfirmResponse, error)
	Recover(context.Context, api.RecoverRequest) (api.RecoverResponse, error)
	Update(context.Context, api.UpdateRequest) (api.UpdateResponse, error)
}

// Auth - Wrapper for AuthService Interface
type Auth struct{}

// ErrorNotFound - User Not Found
var ErrorNotFound = errors.New("Not Found")

// ErrorAlreadyRegistered - User Already Registered
var ErrorAlreadyRegistered = errors.New("Already Registered")

// ErrorInvalidToken - Invalid Token
var ErrorInvalidToken = errors.New("Invalid Token")

// ErrorExpiredToken - Expired Token
var ErrorExpiredToken = errors.New("Expired Token")

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
