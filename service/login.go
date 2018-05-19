package service

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/urantiatech/microservices/auth/api"
	"github.com/urantiatech/microservices/auth/model"
	"golang.org/x/crypto/bcrypt"
)

func (Auth) Login(_ context.Context, req api.LoginRequest) (api.LoginResponse, error) {
	var response api.LoginResponse
	var user model.User

	DB.Where("email = ?", req.Email).Where("confirmed = ?", true).First(&user)

	if user.ID == 0 || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		response.Err = InvalidLogin.Error()
		return response, nil
	}

	// Create an Access JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":   user.ID,
		"fname": user.Fname,
		"lname": user.Lname,
		"email": user.Email,
		"nbf":   time.Now().Unix(),
		"exp":   time.Now().Add(TokenValidity).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(SigningKey)
	response.AccessToken = tokenString
	if err != nil {
		response.Err = err.Error()
	}
	return response, nil
}
