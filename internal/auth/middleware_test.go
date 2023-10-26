package auth

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"

	"github.com/infamous55/habit-tracker/internal/models"
	"github.com/infamous55/habit-tracker/internal/mongodb"
)

func TestExtractUserMiddleware(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		userID := primitive.NewObjectID()
		expectedUser := models.User{
			ID:       userID,
			Name:     "test",
			Email:    "test@test.com",
			Password: "test",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			"collections.users",
			mtest.FirstBatch,
			bson.D{
				{Key: "_id", Value: expectedUser.ID},
				{Key: "name", Value: expectedUser.Name},
				{Key: "email", Value: expectedUser.Email},
				{Key: "password", Value: expectedUser.Password},
			},
		),
		)

		e := echo.New()
		req := &http.Request{
			Header: http.Header{},
		}
		res := &echo.Response{}
		ec := e.NewContext(req, res)

		jwt, err := NewJWTWithCustomClaims(userID)
		assert.Nil(t, err)

		authHeader := fmt.Sprintf("Bearer %s", jwt)
		ec.Request().Header.Set("Authorization", authHeader)

		dbw := mongodb.DatabaseWrapper{Database: mt.DB}

		mockHandlerWasCalled := false
		mockHandler := func(c echo.Context) error {
			mockHandlerWasCalled = true
			return nil
		}

		middleware := ExtractUserMiddleware(dbw)
		err = middleware(mockHandler)(ec)

		assert.Nil(t, err)
		assert.True(t, mockHandlerWasCalled)
	})
}
