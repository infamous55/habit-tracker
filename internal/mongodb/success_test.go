package mongodb

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"

	"github.com/infamous55/habit-tracker/internal/models"
)

func TestGetSuccessByID(t *testing.T) {
	t.Parallel()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		dbw := DatabaseWrapper{mt.DB}

		successID := primitive.NewObjectID()
		habitID := primitive.NewObjectID()

		date := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

		expectedSuccess := models.Success{
			ID:      successID,
			Date:    date,
			HabitID: habitID,
		}

		mt.AddMockResponses(
			mtest.CreateCursorResponse(
				1,
				"collections.successes",
				mtest.FirstBatch,
				bson.D{
					{Key: "_id", Value: successID},
					{Key: "date", Value: date},
					{Key: "habit_id", Value: habitID},
				},
			),
		)

		response, err := dbw.GetSuccessByID(successID)
		assert.Nil(t, err)
		assert.Equal(t, &expectedSuccess, response)
	})
}

func TestGetSuccessByHabitID(t *testing.T) {
	t.Parallel()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		dbw := DatabaseWrapper{mt.DB}

		habitID := primitive.NewObjectID()

		firstSuccessID := primitive.NewObjectID()
		secondSuccessID := primitive.NewObjectID()

		date := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

		firstCursor := mtest.CreateCursorResponse(
			1,
			"collections.successes",
			mtest.FirstBatch,
			bson.D{
				{Key: "_id", Value: firstSuccessID},
				{Key: "date", Value: date},
				{Key: "habit_id", Value: habitID},
			},
		)
		secondCursor := mtest.CreateCursorResponse(
			1,
			"collections.successes",
			mtest.NextBatch,
			bson.D{
				{Key: "_id", Value: secondSuccessID},
				{Key: "date", Value: date},
				{Key: "habit_id", Value: habitID},
			},
		)
		killCursor := mtest.CreateCursorResponse(
			1,
			"collections.successes",
			mtest.NextBatch,
		)
		mt.AddMockResponses(
			firstCursor, secondCursor, killCursor,
		)

		response, err := dbw.GetSuccessesByHabitID(habitID)
		assert.Nil(t, err)
		assert.Equal(t, []*models.Success{
			{
				ID:      firstSuccessID,
				Date:    date,
				HabitID: habitID,
			},
			{
				ID:      secondSuccessID,
				Date:    date,
				HabitID: habitID,
			},
		}, response)
	})
}

func TestCreateSuccess(t *testing.T) {
	t.Parallel()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		dbw := DatabaseWrapper{mt.DB}

		date := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		habitID := primitive.NewObjectID()
		data := models.SuccessCreate{
			Date:    date,
			HabitID: habitID,
		}

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		response, err := dbw.CreateSuccess(data)
		assert.Nil(t, err)
		assert.NotEqual(t, "", response.ID.Hex())
	})
}

func TestDeleteSuccessByID(t *testing.T) {
	t.Parallel()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		dbw := DatabaseWrapper{mt.DB}

		successID := primitive.NewObjectID()
		habitID := primitive.NewObjectID()
		date := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		expectedSuccess := models.Success{
			ID:      successID,
			Date:    date,
			HabitID: habitID,
		}

		mt.AddMockResponses(
			bson.D{
				{Key: "ok", Value: 1},
				{Key: "value", Value: bson.D{
					{Key: "_id", Value: successID},
					{Key: "date", Value: date},
					{Key: "habit_id", Value: habitID},
				}},
			},
		)

		response, err := dbw.DeleteSuccessByID(successID)
		assert.Nil(t, err)
		assert.Equal(t, &expectedSuccess, response)
	})
}
