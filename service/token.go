package service

import (
	"fmt"
	"math/rand"
	"strconv"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/dgrijalva/jwt-go"
	"github.com/urantiatech/microservices/auth/model"
)

func ParseToken(tokenString string) (model.User, error) {
	var user model.User

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, InvalidToken
		}

		// signingKey is a []byte containing your secret, e.g. []byte("my_secret_key")
		return SigningKey, nil
	})

	if err != nil || !token.Valid {
		return user, InvalidToken
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uid := fmt.Sprint(claims["uid"])
		user.ID, _ = strconv.ParseUint(uid, 10, 64)
		user.Fname = claims["fname"].(string)
		user.Lname = claims["lname"].(string)
		user.Email = claims["email"].(string)
	}
	return user, nil
}

func RandomToken(size int) string {
	b := make([]byte, size)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
