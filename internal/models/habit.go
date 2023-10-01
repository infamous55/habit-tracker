package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Habit struct {
	ID          primitive.ObjectID `json:"id"                    bson:"_id"`
	Name        string             `json:"name"                  bson:"name"`
	Description *string            `json:"description,omitempty" bson:"description,omitempty"`
	GroupID     primitive.ObjectID `json:"group_id"              bson:"group_id"`
	UserID      primitive.ObjectID `json:"user_id"               bson:"user_id"`
}
