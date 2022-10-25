package lib

import (
	"burlyeducation/log"

	"github.com/golang-jwt/jwt/v4"
)

type Tokener struct {
}

type MyCustomClaims struct {
	jwt.StandardClaims
	UserId int    `json:"user_id"`
	Mobile string `json:"verified_mobile"`
}

func (t Tokener) Generate(secret string, claim jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	ss, err := token.SignedString([]byte(secret))
	return ss, err
}

func (t Tokener) Validate(tokenString, secret string, claim jwt.Claims) (*jwt.Token, error) {

	token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err == nil {
		return token, nil
	} else {
		log.Info(1104, map[string]interface{}{"error_details": err})
		return nil, err
	}
}
