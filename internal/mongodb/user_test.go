package mongodb

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"

	"github.com/infamous55/habit-tracker/internal/models"
)

func TestGetUserByID(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		dbw := DatabaseWrapper{mt.DB}

		expectedUser := models.User{
			ID:       primitive.NewObjectID(),
			Name:     "test",
			Email:    "test@test.com",
			Password: "test",
		}

		mt.AddMockResponses(
			mtest.CreateCursorResponse(
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

		response, err := dbw.GetUserByID(expectedUser.ID)
		assert.Nil(t, err)
		assert.Equal(t, &expectedUser, response)
	})
}

func TestCreateUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		dbw := DatabaseWrapper{mt.DB}

		credentials := models.Credentials{
			Email:    "test@test.com",
			Password: "test",
		}

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		response, err := dbw.CreateUser(credentials)
		assert.Nil(t, err)

		assert.Nil(t, response.ComparePassword(credentials.Password))
		assert.NotEqual(t, response.ID.Hex(), "")
	})

	mt.Run("invalid email", func(mt *mtest.T) {
		dbw := DatabaseWrapper{mt.DB}

		credentials := models.Credentials{
			Email:    "test",
			Password: "test",
		}

		_, err := dbw.CreateUser(credentials)
		assert.Error(t, err)
	})

	mt.Run("duplicate key", func(mt *mtest.T) {
		dbw := DatabaseWrapper{mt.DB}

		credentials := models.Credentials{
			Email:    "test@test.com",
			Password: "test",
		}

		mt.AddMockResponses(
			mtest.CreateWriteErrorsResponse(mtest.WriteError{
				Index:   1,
				Code:    11000,
				Message: "duplicate key error",
			}),
		)

		_, err := dbw.CreateUser(credentials)
		assert.Error(t, err)
	})
}

func TestGetUserByEmail(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		dbw := DatabaseWrapper{mt.DB}

		expectedUser := models.User{
			ID:       primitive.NewObjectID(),
			Name:     "test",
			Email:    "test@test.com",
			Password: "test",
		}

		mt.AddMockResponses(
			mtest.CreateCursorResponse(
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

		response, err := dbw.GetUserByEmail(expectedUser.Email)
		assert.Nil(t, err)
		assert.Equal(t, &expectedUser, response)
	})
}
