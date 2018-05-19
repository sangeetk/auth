package service

import (
	"context"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/patrickmn/go-cache"
	"github.com/urantiatech/microservices/auth/api"
)

var DB *gorm.DB

var SigningKey []byte
var TokenValidity time.Duration
var BlacklistTokens *cache.Cache

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

type Auth struct{}

var NotFound = errors.New("Not Found")
var AlreadyRegistered = errors.New("Already Registered")
var InvalidToken = errors.New("Invalid Token")
var ExpiredToken = errors.New("Expired Token")
var InvalidLogin = errors.New("Invalid Login")
var InvalidRequest = errors.New("Invalid Request")
var UnknownError = errors.New("Unknown Error")

// ServiceMiddleware is a chainable behavior modifier for AuthService.
type ServiceMiddleware func(AuthService) AuthService
