package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/infamous55/habit-tracker/internal/models"
)

func (db *DatabaseWrapper) GetGroupsByUserID(userID primitive.ObjectID) ([]*models.Group, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := db.Collection("groups").Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []*models.Group
	err = cur.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (db *DatabaseWrapper) GetGroupByID(id primitive.ObjectID) (*models.Group, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := db.Collection("groups").FindOne(ctx, bson.M{"_id": id})

	var group *models.Group
	err := result.Decode(&group)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (db *DatabaseWrapper) CreateGroup(data models.NewGroup) (*models.Group, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := db.Collection("groups").InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("invalid inserted ID: %v", result.InsertedID)
	}

	return &models.Group{
		ID:          insertedID,
		Name:        data.Name,
		Description: data.Description,
		UserID:      data.UserID,
	}, nil
}

func (db *DatabaseWrapper) UpdateGroup(data models.GroupUpdate) (*models.Group, error) {
	filter := bson.D{{Key: "_id", Value: data.ID}}

	update := bson.D{}
	if data.Name != nil {
		update = append(
			update,
			bson.E{Key: "$set", Value: bson.D{{Key: "name", Value: *data.Name}}},
		)
	}
	if data.Description != nil {
		update = append(
			update,
			bson.E{Key: "$set", Value: bson.D{{Key: "description", Value: *data.Description}}},
		)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.Collection("groups").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return db.GetGroupByID(data.ID)
}

func (db *DatabaseWrapper) DeleteGroupByID(id primitive.ObjectID) (*models.Group, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := db.Collection("groups").FindOneAndDelete(ctx, filter)
	err := result.Err()
	if err != nil {
		return nil, err
	}

	var deletedGroup models.Group
	err = result.Decode(&deletedGroup)
	if err != nil {
		return nil, err
	}
	return &deletedGroup, nil
}
