package mongodb

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"

	"github.com/infamous55/habit-tracker/internal/models"
)

func TestGetGroupsByUserID(t *testing.T) {
	t.Parallel()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		dbw := DatabaseWrapper{mt.DB}

		userID := primitive.NewObjectID()

		firstGroupID := primitive.NewObjectID()
		secondGroupID := primitive.NewObjectID()

		description := "test"

		firstCursor := mtest.CreateCursorResponse(
			1,
			"collections.groups",
			mtest.FirstBatch,
			bson.D{
				{Key: "_id", Value: firstGroupID},
				{Key: "name", Value: "test"},
				{Key: "description", Value: description},
				{Key: "user_id", Value: userID},
			},
		)
		secondCursor := mtest.CreateCursorResponse(
			1,
			"collections.groups",
			mtest.NextBatch,
			bson.D{
				{Key: "_id", Value: secondGroupID},
				{Key: "name", Value: "test"},
				{Key: "description", Value: description},
				{Key: "user_id", Value: userID},
			},
		)
		killCursor := mtest.CreateCursorResponse(
			1,
			"collections.groups",
			mtest.NextBatch,
		)
		mt.AddMockResponses(
			firstCursor, secondCursor, killCursor,
		)

		response, err := dbw.GetGroupsByUserID(userID)
		assert.Nil(t, err)
		assert.Equal(t, []*models.Group{
			{
				ID:          firstGroupID,
				Name:        "test",
				Description: &description,
				UserID:      userID,
			},
			{
				ID:          secondGroupID,
				Name:        "test",
				Description: &description,
				UserID:      userID,
			},
		}, response)
	})
}

func TestGetGroupByID(t *testing.T) {
	t.Parallel()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		dbw := DatabaseWrapper{mt.DB}

		groupID := primitive.NewObjectID()
		userID := primitive.NewObjectID()

		description := "test"

		expectedGroup := models.Group{
			ID:          groupID,
			Name:        "test",
			Description: &description,
			UserID:      userID,
		}

		mt.AddMockResponses(
			mtest.CreateCursorResponse(
				1,
				"collections.groups",
				mtest.FirstBatch,
				bson.D{
					{Key: "_id", Value: groupID},
					{Key: "name", Value: "test"},
					{Key: "description", Value: description},
					{Key: "user_id", Value: userID},
				},
			),
		)

		response, err := dbw.GetGroupByID(groupID)
		assert.Nil(t, err)
		assert.Equal(t, &expectedGroup, response)
	})
}

func TestCreateGroup(t *testing.T) {
	t.Parallel()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		dbw := DatabaseWrapper{mt.DB}

		description := "test"
		userID := primitive.NewObjectID()
		data := models.NewGroup{
			Name:        "test",
			Description: &description,
			UserID:      userID,
		}

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		response, err := dbw.CreateGroup(data)
		assert.Nil(t, err)
		assert.NotEqual(t, "", response.ID.Hex())
	})
}

func TestUpdateGroup(t *testing.T) {
	t.Parallel()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		dbw := DatabaseWrapper{mt.DB}

		groupID := primitive.NewObjectID()
		userID := primitive.NewObjectID()
		name := "test"
		description := "test"

		data := models.GroupUpdate{
			ID:          groupID,
			Name:        &name,
			Description: &description,
		}

		expectedGroup := models.Group{
			ID:          groupID,
			Name:        name,
			Description: &description,
			UserID:      userID,
		}

		updateResponse := bson.D{{Key: "ok", Value: 1}, {Key: "nModified", Value: 1}}
		getResponse := mtest.CreateCursorResponse(
			1,
			"collections.groups",
			mtest.FirstBatch,
			bson.D{
				{Key: "_id", Value: groupID},
				{Key: "name", Value: "test"},
				{Key: "description", Value: description},
				{Key: "user_id", Value: userID},
			},
		)
		mt.AddMockResponses(updateResponse, getResponse)

		response, err := dbw.UpdateGroup(data)
		assert.Nil(t, err)
		assert.Equal(t, &expectedGroup, response)
	})
}

func TestDeleteGroupByID(t *testing.T) {
	t.Parallel()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		dbw := DatabaseWrapper{mt.DB}

		groupID := primitive.NewObjectID()
		userID := primitive.NewObjectID()
		description := "test"

		expectedGroup := models.Group{
			ID:          groupID,
			Name:        "test",
			Description: &description,
			UserID:      userID,
		}

		mt.AddMockResponses(
			bson.D{
				{Key: "ok", Value: 1},
				{Key: "value", Value: bson.D{
					{Key: "_id", Value: groupID},
					{Key: "name", Value: "test"},
					{Key: "description", Value: description},
					{Key: "user_id", Value: userID},
				}},
			},
		)

		response, err := dbw.DeleteGroupByID(groupID)
		assert.Nil(t, err)
		assert.Equal(t, &expectedGroup, response)
	})
}
