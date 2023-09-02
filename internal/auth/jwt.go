package auth

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

func ParseWithCustomClaims(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return os.Getenv("JWT_SECRET"), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims type: %T", claims)
	}
	return claims, nil
}
