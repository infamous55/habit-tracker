package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Habit struct {
	ID          primitive.ObjectID `json:"id"                    bson:"_id"`
	Name        string             `json:"name"                  bson:"name"`
	Description *string            `json:"description,omitempty" bson:"description,omitempty"`
	Schedule    Schedule           `json:"schedule"              bson:"schedule"`
	GroupID     primitive.ObjectID `json:"group_id"              bson:"group_id"`
	UserID      primitive.ObjectID `json:"user_id"               bson:"user_id"`
}

type NewHabit struct {
	Name        string        `json:"name"                  bson:"name"`
	Description *string       `json:"description,omitempty" bson:"description,omitempty"`
	Schedule    ScheduleInput `json:"schedule"              bson:"schedule"`
	GroupID     string        `json:"group_id"              bson:"group_id"`
}

type HabitCreate struct {
	Name        string             `json:"name"                  bson:"name"`
	Description *string            `json:"description,omitempty" bson:"description,omitempty"`
	Schedule    ScheduleInput      `json:"schedule"              bson:"schedule"`
	GroupID     primitive.ObjectID `json:"group_id"              bson:"group_id"`
	UserID      primitive.ObjectID `json:"user_id"               bson:"user_id"`
}

type HabitUpdate struct {
	ID          primitive.ObjectID  `json:"id"                    bson:"_id"`
	Name        *string             `json:"name,omitempty"        bson:"name,omitempty"`
	Description *string             `json:"description,omitempty" bson:"description,omitempty"`
	Schedule    *ScheduleInput      `json:"schedule,omitempty"    bson:"schedule,omitempty"`
	GroupID     *primitive.ObjectID `json:"group_id,omitempty"    bson:"group_id,omitempty"`
}

type HabitFilterOptions struct {
	UserID    primitive.ObjectID
	GroupID   *primitive.ObjectID
	StartDate *time.Time
	EndDate   *time.Time
	Succeded  *bool
}
