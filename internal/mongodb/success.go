package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/infamous55/habit-tracker/internal/models"
)

func (db *DatabaseWrapper) GetSuccessByID(id primitive.ObjectID) (*models.Success, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := db.Collection("successes").FindOne(ctx, bson.M{"_id": id})

	var success *models.Success
	err := result.Decode(&success)
	if err != nil {
		return nil, err
	}
	return success, nil
}

func (db *DatabaseWrapper) DeleteSuccessByID(id primitive.ObjectID) (*models.Success, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := db.Collection("successes").FindOneAndDelete(ctx, filter)
	err := result.Err()
	if err != nil {
		return nil, err
	}

	var deletedSuccess models.Success
	err = result.Decode(&deletedSuccess)
	if err != nil {
		return nil, err
	}
	return &deletedSuccess, nil
}

func (db *DatabaseWrapper) GetSuccessesByHabitID(
	habitID primitive.ObjectID,
) ([]*models.Success, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := db.Collection("successes").Find(ctx, bson.M{"habit_id": habitID})
	if err != nil {
		return nil, err
	}

	var results []*models.Success
	err = cur.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (db *DatabaseWrapper) CreateSuccess(input models.SuccessCreate) (*models.Success, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := db.Collection("successes").InsertOne(ctx, input)
	if err != nil {
		return nil, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("invalid inserted ID: %v", result.InsertedID)
	}

	return &models.Success{
		ID:      insertedID,
		Date:    input.Date,
		HabitID: input.HabitID,
	}, nil
}
