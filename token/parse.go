package token

import (
	"github.com/dgrijalva/jwt-go"
)

// ParseToken - Parses the access token and extract username etc
func ParseToken(tokenString string) (*Token, error) {
	var t = new(Token)

	if tokenString == "" {
		return nil, ErrorInvalidToken
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrorInvalidToken
		}

		// signingKey is a []byte containing your secret, e.g. []byte("my_secret_key")
		return SigningKey, nil
	})

	if err != nil || !token.Valid {
		return nil, ErrorInvalidToken
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		t.Username = claims["username"].(string)
		if _, ok := claims["fname"]; ok {
			t.FirstName = claims["fname"].(string)
		}
		if _, ok := claims["lname"]; ok {
			t.LastName = claims["lname"].(string)
		}
		if _, ok := claims["email"]; ok {
			t.Email = claims["email"].(string)
		}
		if _, ok := claims["domain"]; ok {
			t.Domain = claims["domain"].(string)
		}
		if roles, ok := claims["roles"]; ok {
			t.Roles = []string{}
			if roles != nil {
				for _, role := range roles.([]interface{}) {
					t.Roles = append(t.Roles, role.(string))
				}
			}
		}
	}
	return t, nil
}
