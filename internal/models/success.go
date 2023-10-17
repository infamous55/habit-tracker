package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Success struct {
	ID      primitive.ObjectID `json:"id"       bson:"_id"`
	Date    time.Time          `json:"date"     bson:"date"`
	HabitID primitive.ObjectID `json:"habit_id" bson:"habit_id"`
}

type NewSuccess struct {
	Date    time.Time `json:"date"     bson:"date"`
	HabitID string    `json:"habit_id" bson:"habit_id"`
}
