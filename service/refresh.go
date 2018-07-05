package service

import (
	"context"
	"time"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/user"
	"github.com/dgrijalva/jwt-go"
)

// Refresh - Generates a new token and extends existing session
func (Auth) Refresh(_ context.Context, req api.RefreshRequest) (api.RefreshResponse, error) {
	var response api.RefreshResponse
	var u *user.User

	// Validate the access token and get user info
	u1, err := ParseToken(req.AccessToken)
	if err == ErrorInvalidToken {
		response.Err = err.Error()
		return response, nil
	}

	// Check by looking for Blacklisted tokens in Cache
	if _, found := BlacklistTokens.Get(req.AccessToken); found {
		response.Err = ErrorInvalidToken.Error()
		return response, nil
	}

	// Validate the refresh token and get user info
	u2, err := ParseToken(req.RefreshToken)
	if err == ErrorInvalidToken {
		response.Err = ErrorInvalidToken.Error()
		return response, nil
	}

	// Check if both these tokens belongs to same User
	if u1.Username != u2.Username {
		response.Err = ErrorInvalidToken.Error()
		return response, nil
	}

	u, err = user.Read(u1.Username)
	if err != nil || u.Confirmed != true {
		response.Err = ErrorNotFound.Error()
		return response, nil
	}

	// Create an Access JWT Token
	atoken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Username,
		"fname":    u.FirstName,
		"lname":    u.LastName,
		"email":    u.Email,
		"nbf":      time.Now().Unix(),
		"exp":      time.Now().Add(AccessTokenValidity).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	response.NewAccessToken, err = atoken.SignedString(SigningKey)
	if err != nil {
		response.Err = err.Error()
	}

	// Create an Access JWT Token
	rtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Username,
		"fname":    u.FirstName,
		"lname":    u.LastName,
		"email":    u.Email,
		"nbf":      time.Now().Unix(),
		"exp":      time.Now().Add(RefreshTokenValidity).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	response.NewRefreshToken, err = rtoken.SignedString(SigningKey)
	if err != nil {
		response.Err = err.Error()
	}

	return response, nil
}
