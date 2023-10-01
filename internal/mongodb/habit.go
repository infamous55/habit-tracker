package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/infamous55/habit-tracker/internal/models"
)

func (db *DatabaseWrapper) GetHabitsByUserID(userID primitive.ObjectID) ([]*models.Habit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := db.Collection("habits").Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []*models.Habit
	err = cur.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (db *DatabaseWrapper) GetHabitsByGroupID(groupID primitive.ObjectID) ([]*models.Habit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := db.Collection("habits").Find(ctx, bson.M{"group_id": groupID})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []*models.Habit
	err = cur.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
