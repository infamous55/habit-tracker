package mongodb

import (
	"context"
	"time"

	"github.com/infamous55/habit-tracker/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db DatabaseWrapper) GetUserById(id string) (*models.User, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	queryContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := db.Collection("user").FindOne(queryContext, bson.M{"_id": ID})

	var user *models.User
	err = result.Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
