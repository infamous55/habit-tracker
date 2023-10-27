package auth

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestParseJWTWithCustomClaims(t *testing.T) {
	t.Parallel()

	userID := primitive.NewObjectID()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	})

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	assert.NoError(t, err)

	claims, err := ParseJWTWithCustomClaims(signedToken)
	assert.Nil(t, err)
	assert.Equal(t, userID, claims.UserID)
}

func TestNewJWTWithCustomClaims(t *testing.T) {
	t.Parallel()

	userID := primitive.NewObjectID()

	tokenString, err := NewJWTWithCustomClaims(userID)
	assert.Nil(t, err)

	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		},
	)
	assert.Nil(t, err)

	claims, ok := token.Claims.(*CustomClaims)
	assert.True(t, ok)

	assert.Equal(t, userID, claims.UserID)
}
