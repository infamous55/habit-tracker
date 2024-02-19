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

func TestGetHabitByID(t *testing.T) {
	t.Parallel()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		dbw := DatabaseWrapper{mt.DB}

		habitID := primitive.NewObjectID()
		groupID := primitive.NewObjectID()
		userID := primitive.NewObjectID()

		description := "test"

		expectedHabit := models.Habit{
			ID:          habitID,
			Name:        "test",
			Description: &description,
			Schedule: models.Schedule{
				Type:  models.ScheduleTypeWeekly,
				Start: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			GroupID: groupID,
			UserID:  userID,
		}

		mt.AddMockResponses(
			mtest.CreateCursorResponse(
				1,
				"collections.habits",
				mtest.FirstBatch,
				bson.D{
					{Key: "_id", Value: habitID},
					{Key: "name", Value: expectedHabit.Name},
					{Key: "description", Value: expectedHabit.Description},
					{Key: "schedule", Value: expectedHabit.Schedule},
					{Key: "group_id", Value: expectedHabit.GroupID},
					{Key: "user_id", Value: expectedHabit.UserID},
				},
			),
		)

		response, err := dbw.GetHabitByID(habitID)
		assert.Nil(t, err)
		assert.Equal(t, &expectedHabit, response)
	})
}

func TestGetHabitsByUserID(t *testing.T) {
	t.Parallel()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		dbw := DatabaseWrapper{mt.DB}

		userID := primitive.NewObjectID()

		firstHabitID := primitive.NewObjectID()
		secondHabitID := primitive.NewObjectID()
		groupID := primitive.NewObjectID()

		description := "test"
		schedule := models.Schedule{
			Type:  models.ScheduleTypeWeekly,
			Start: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		}

		firstCursor := mtest.CreateCursorResponse(
			1,
			"collections.habits",
			mtest.FirstBatch,
			bson.D{
				{Key: "_id", Value: firstHabitID},
				{Key: "name", Value: "test"},
				{Key: "description", Value: description},
				{Key: "schedule", Value: schedule},
				{Key: "group_id", Value: groupID},
				{Key: "user_id", Value: userID},
			},
		)
		secondCursor := mtest.CreateCursorResponse(
			1,
			"collections.habits",
			mtest.NextBatch,
			bson.D{
				{Key: "_id", Value: secondHabitID},
				{Key: "name", Value: "test"},
				{Key: "description", Value: description},
				{Key: "schedule", Value: schedule},
				{Key: "group_id", Value: groupID},
				{Key: "user_id", Value: userID},
			},
		)
		killCursor := mtest.CreateCursorResponse(
			1,
			"collections.habits",
			mtest.NextBatch,
		)

		mt.AddMockResponses(
			firstCursor, secondCursor, killCursor,
		)

		response, err := dbw.GetHabitsByUserID(userID)
		assert.Nil(t, err)
		assert.Equal(t, []*models.Habit{
			{
				ID:          firstHabitID,
				Name:        "test",
				Description: &description,
				Schedule:    schedule,
				GroupID:     groupID,
				UserID:      userID,
			},
			{
				ID:          secondHabitID,
				Name:        "test",
				Description: &description,
				Schedule:    schedule,
				GroupID:     groupID,
				UserID:      userID,
			},
		}, response)
	})
}

func TestGetHabitsByGroupID(t *testing.T) {
	t.Parallel()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		dbw := DatabaseWrapper{mt.DB}

		groupID := primitive.NewObjectID()

		firstHabitID := primitive.NewObjectID()
		secondHabitID := primitive.NewObjectID()
		userID := primitive.NewObjectID()

		description := "test"
		schedule := models.Schedule{
			Type:  models.ScheduleTypeWeekly,
			Start: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		}

		firstCursor := mtest.CreateCursorResponse(
			1,
			"collections.habits",
			mtest.FirstBatch,
			bson.D{
				{Key: "_id", Value: firstHabitID},
				{Key: "name", Value: "test"},
				{Key: "description", Value: description},
				{Key: "schedule", Value: schedule},
				{Key: "group_id", Value: groupID},
				{Key: "user_id", Value: userID},
			},
		)
		secondCursor := mtest.CreateCursorResponse(
			1,
			"collections.habits",
			mtest.NextBatch,
			bson.D{
				{Key: "_id", Value: secondHabitID},
				{Key: "name", Value: "test"},
				{Key: "description", Value: description},
				{Key: "schedule", Value: schedule},
				{Key: "group_id", Value: groupID},
				{Key: "user_id", Value: userID},
			},
		)
		killCursor := mtest.CreateCursorResponse(
			1,
			"collections.habits",
			mtest.NextBatch,
		)

		mt.AddMockResponses(
			firstCursor, secondCursor, killCursor,
		)

		response, err := dbw.GetHabitsByGroupID(groupID)
		assert.Nil(t, err)
		assert.Equal(t, []*models.Habit{
			{
				ID:          firstHabitID,
				Name:        "test",
				Description: &description,
				Schedule:    schedule,
				GroupID:     groupID,
				UserID:      userID,
			},
			{
				ID:          secondHabitID,
				Name:        "test",
				Description: &description,
				Schedule:    schedule,
				GroupID:     groupID,
				UserID:      userID,
			},
		}, response)
	})
}

func TestCreateHabit(t *testing.T) {
	t.Parallel()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		dbw := DatabaseWrapper{mt.DB}

		description := "test"
		schedule := models.ScheduleInput{
			Type:  models.ScheduleTypeWeekly,
			Start: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		}
		data := models.HabitCreate{
			Name:        "test",
			Description: &description,
			Schedule:    schedule,
			GroupID:     primitive.NewObjectID(),
			UserID:      primitive.NewObjectID(),
		}

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		response, err := dbw.CreateHabit(data)
		assert.Nil(t, err)
		assert.NotEqual(t, "", response.ID.Hex())
	})
}

func TestUpdateHabit(t *testing.T) {
	t.Parallel()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		dbw := DatabaseWrapper{mt.DB}

		habitID := primitive.NewObjectID()
		name := "test"
		description := "test"
		scheduleInput := models.ScheduleInput{
			Type:  models.ScheduleTypeWeekly,
			Start: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		}
		groupID := primitive.NewObjectID()

		data := models.HabitUpdate{
			ID:          habitID,
			Name:        &name,
			Description: &description,
			Schedule:    &scheduleInput,
			GroupID:     &groupID,
		}

		schedule := models.Schedule{
			Type:  models.ScheduleTypeWeekly,
			Start: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		}
		expectedHabit := models.Habit{
			ID:          habitID,
			Name:        name,
			Description: &description,
			Schedule:    schedule,
			GroupID:     groupID,
		}

		updateResponse := bson.D{{Key: "ok", Value: 1}, {Key: "nModified", Value: 1}}
		getResponse := mtest.CreateCursorResponse(
			1,
			"collections.habits",
			mtest.FirstBatch,
			bson.D{
				{Key: "_id", Value: habitID},
				{Key: "name", Value: name},
				{Key: "description", Value: description},
				{Key: "schedule", Value: schedule},
				{Key: "group_id", Value: groupID},
			},
		)

		mt.AddMockResponses(
			updateResponse, getResponse,
		)

		response, err := dbw.UpdateHabit(data)
		assert.Nil(t, err)
		assert.Equal(t, &expectedHabit, response)
	})
}

func TestDeleteHabitByID(t *testing.T) {
	t.Parallel()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		dbw := DatabaseWrapper{mt.DB}

		habitID := primitive.NewObjectID()
		description := "test"
		schedule := models.Schedule{
			Type:  models.ScheduleTypeWeekly,
			Start: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		}
		groupID := primitive.NewObjectID()

		expectedHabit := models.Habit{
			ID:          habitID,
			Name:        "test",
			Description: &description,
			Schedule:    schedule,
			GroupID:     groupID,
		}

		mt.AddMockResponses(
			bson.D{
				{Key: "ok", Value: 1},
				{Key: "value", Value: bson.D{
					{Key: "_id", Value: habitID},
					{Key: "name", Value: "test"},
					{Key: "description", Value: description},
					{Key: "schedule", Value: schedule},
					{Key: "group_id", Value: groupID},
				}},
			})

		response, err := dbw.DeleteHabitByID(habitID)
		assert.Nil(t, err)
		assert.Equal(t, &expectedHabit, response)
	})
}
