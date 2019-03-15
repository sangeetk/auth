package token

import (
	"time"

	"git.urantiatech.com/auth/auth/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/patrickmn/go-cache"
)

// SigningKey - JWT Signing Key
var SigningKey []byte

// AccessTokenValidity - JWT Access Token Validity
var AccessTokenValidity time.Duration

// RefreshTokenValidity - JWT Refresh Token Validity
var RefreshTokenValidity time.Duration

// RememberMeAccessTokenValidity - JWT Access Token Validity
var RememberMeAccessTokenValidity time.Duration

// RememberMeRefreshTokenValidity - JWT Refresh Token Validity
var RememberMeRefreshTokenValidity time.Duration

// ResetTokenValidity - JWT Reset Token Validity
var ResetTokenValidity time.Duration

// ConfirmTokenValidity - JWT Confirm Token Validity
var ConfirmTokenValidity time.Duration

// UpdateTokenValidity - JWT Update Token Validity
var UpdateTokenValidity time.Duration

// BlacklistAccessTokens - Cache to store invalid access tokens
var BlacklistAccessTokens *cache.Cache

// BlacklistRefreshTokens - Cache to store invalid refresh tokens
var BlacklistRefreshTokens *cache.Cache

// Token - token fields
type Token struct {
	Username  string   `json:"username"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email"`
	Domain    string   `json:"domain"`
	Roles     []string `json:"roles"`
}

// NewToken creates a new Token
func NewToken(u *user.User, domain string, validity time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Username,
		"fname":    u.FirstName,
		"lname":    u.LastName,
		"email":    u.Email,
		"domain":   domain,
		"roles":    u.GetRoles(domain),
		"nbf":      time.Now().Unix(),
		"exp":      time.Now().Add(validity).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString(SigningKey)
}
