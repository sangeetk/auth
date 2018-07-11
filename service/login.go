package service

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/urantiatech/kit/endpoint"
	"golang.org/x/crypto/bcrypt"
)

// Login - Log in the user after credentials are successfully authenticated
func (Auth) Login(_ context.Context, req api.LoginRequest) (api.LoginResponse, error) {
	var response api.LoginResponse

	// Get user details
	user, err := user.Read(req.Username)
	if err != nil || user.Confirmed != true {
		response.Err = ErrorInvalidLogin.Error()
		return response, nil
	}

	// Check password
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		response.Err = ErrorInvalidLogin.Error()
		return response, nil
	}

	// Create an Access JWT Token
	atoken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"fname":    user.FirstName,
		"lname":    user.LastName,
		"email":    user.Email,
		"nbf":      time.Now().Unix(),
		"exp":      time.Now().Add(AccessTokenValidity).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	response.AccessToken, err = atoken.SignedString(SigningKey)
	if err != nil {
		response.Err = err.Error()
	}

	// Create an Access JWT Token
	rtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"fname":    user.FirstName,
		"lname":    user.LastName,
		"email":    user.Email,
		"nbf":      time.Now().Unix(),
		"exp":      time.Now().Add(RefreshTokenValidity).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	response.RefreshToken, err = rtoken.SignedString(SigningKey)
	if err != nil {
		response.Err = err.Error()
	}

	return response, nil
}

// MakeLoginEndpoint -
func MakeLoginEndpoint(svc Auth) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.LoginRequest)
		return svc.Login(ctx, req)
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
