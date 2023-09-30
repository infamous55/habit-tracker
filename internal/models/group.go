package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Group struct {
	ID          primitive.ObjectID `json:"id"                    bson:"_id"`
	Name        string             `json:"name"                  bson:"name"`
	Description *string            `json:"description,omitempty" bson:"description,omitempty"`
	UserID      primitive.ObjectID `json:"user_id"               bson:"user_id"`
}

type NewGroup struct {
	Name        string             `json:"name"                  bson:"name"`
	Description *string            `json:"description,omitempty" bson:"description,omitempty"`
	UserID      primitive.ObjectID `json:"user_id"               bson:"user_id"`
}
