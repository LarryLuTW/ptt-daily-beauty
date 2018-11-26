package jwt

import (
	"fmt"

	jwtGo "github.com/dgrijalva/jwt-go"
)

var secret = []byte("I am Larry Lu, a Gopher.")

// ParseToken parse a token string and get email
func ParseToken(tokenStr string) (email string, err error) {
	// ref: https://godoc.org/github.com/dgrijalva/jwt-go#example-Parse--Hmac
	keyFunc := func(token *jwtGo.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtGo.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	}

	token, err := jwtGo.Parse(tokenStr, keyFunc)
	if claims, ok := token.Claims.(jwtGo.MapClaims); ok && token.Valid {
		email := claims["email"].(string)
		return email, nil
	} else {
		return "", err
	}
}

// NewToken accept a email and generate a token
func NewToken(email string) (tokenStr string) {
	// ref: https://godoc.org/github.com/dgrijalva/jwt-go#example-New--Hmac
	payload := jwtGo.MapClaims{"email": email}
	token := jwtGo.NewWithClaims(jwtGo.SigningMethodHS256, payload)
	tokenStr, _ = token.SignedString(secret)
	return tokenStr
}
