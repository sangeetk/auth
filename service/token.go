package service

import (
	"fmt"
	"math/rand"

	"git.urantiatech.com/auth/auth/user"
	"github.com/dgrijalva/jwt-go"
)

// ParseToken - Parses the access token and extract username, roles etc
func ParseToken(tokenString string) (*user.User, error) {
	var u = new(user.User)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrorInvalidToken
		}

		// signingKey is a []byte containing your secret, e.g. []byte("my_secret_key")
		return SigningKey, nil
	})

	if err != nil || !token.Valid {
		return u, ErrorInvalidToken
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		u.Username = claims["username"].(string)
		u.FirstName = claims["fname"].(string)
		u.LastName = claims["lname"].(string)
		u.Email = claims["email"].(string)
	}
	return u, nil
}

// RandomToken - Generates a random token
func RandomToken(size int) string {
	b := make([]byte, size)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
