package models

import (
	"slices"
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

type HabitData struct {
	ID          string         `json:"id"`
	Name        *string        `json:"name,omitempty"`
	Description *string        `json:"description,omitempty"`
	Schedule    *ScheduleInput `json:"schedule,omitempty"`
	GroupID     *string        `json:"group_id,omitempty"`
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
	Succeeded *bool
}

func (h *Habit) IsScheduled(date time.Time) bool {
	startDate := h.Schedule.Start.Truncate(24 * time.Hour)
	currentDate := date.Truncate(24 * time.Hour)

	if currentDate.Before(startDate) {
		return false
	}

	if currentDate.Equal(startDate) {
		return true
	}

	switch h.Schedule.Type {
	case ScheduleTypeWeekly:
		return slices.Contains(h.Schedule.Weekdays, TimeToWeekday(currentDate))
	case ScheduleTypeMonthly:
		return slices.Contains(h.Schedule.Monthdays, currentDate.Day())
	case ScheduleTypePeriodic:
		return int(currentDate.Sub(startDate).Hours()/24)%*h.Schedule.PeriodInDays == 0
	default:
		return false
	}
}

func TimeToWeekday(date time.Time) Weekday {
	switch date.Weekday() {
	case time.Monday:
		return WeekdayMonday
	case time.Tuesday:
		return WeekdayTuesday
	case time.Wednesday:
		return WeekdayWednesday
	case time.Thursday:
		return WeekdayThursday
	case time.Friday:
		return WeekdayFriday
	case time.Saturday:
		return WeekdaySaturday
	case time.Sunday:
		return WeekdaySunday
	}
	return ""
}
