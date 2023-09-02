// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
)

type AuthData struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Group struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description *string  `json:"description,omitempty"`
	Habits      []*Habit `json:"habits"`
	User        *User    `json:"user"`
}

type GroupData struct {
	ID          string  `json:"id"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type Habit struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	Schedule    *Schedule  `json:"schedule"`
	Successes   []*Success `json:"successes"`
	Group       *Group     `json:"group"`
	User        *User      `json:"user"`
}

type HabitData struct {
	ID          string         `json:"id"`
	Name        *string        `json:"name,omitempty"`
	Description *string        `json:"description,omitempty"`
	Schedule    *ScheduleInput `json:"schedule,omitempty"`
	GroupID     *string        `json:"groupId,omitempty"`
}

type NewGroup struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type NewHabit struct {
	Name        string         `json:"name"`
	Description *string        `json:"description,omitempty"`
	Schedule    *ScheduleInput `json:"schedule"`
	GroupID     string         `json:"groupId"`
}

type NewSuccess struct {
	Date    string `json:"date"`
	HabitID string `json:"habitId"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type Schedule struct {
	Type           ScheduleType `json:"type"`
	Weekdays       []Weekday    `json:"weekdays,omitempty"`
	Monthdays      []int        `json:"monthdays,omitempty"`
	RepeatInterval *int         `json:"repeatInterval,omitempty"`
	RepeatUnit     *RepeatUnit  `json:"repeatUnit,omitempty"`
	Start          string       `json:"start"`
}

type ScheduleInput struct {
	Type           ScheduleType `json:"type"`
	Weekdays       []Weekday    `json:"weekdays,omitempty"`
	Monthdays      []int        `json:"monthdays,omitempty"`
	RepeatInterval *int         `json:"repeatInterval,omitempty"`
	RepeatUnit     *RepeatUnit  `json:"repeatUnit,omitempty"`
	Start          string       `json:"start"`
}

type Success struct {
	ID    string `json:"id"`
	Date  string `json:"date"`
	Habit *Habit `json:"habit"`
}

type RepeatUnit string

const (
	RepeatUnitDay   RepeatUnit = "DAY"
	RepeatUnitWeek  RepeatUnit = "WEEK"
	RepeatUnitMonth RepeatUnit = "MONTH"
)

var AllRepeatUnit = []RepeatUnit{
	RepeatUnitDay,
	RepeatUnitWeek,
	RepeatUnitMonth,
}

func (e RepeatUnit) IsValid() bool {
	switch e {
	case RepeatUnitDay, RepeatUnitWeek, RepeatUnitMonth:
		return true
	}
	return false
}

func (e RepeatUnit) String() string {
	return string(e)
}

func (e *RepeatUnit) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RepeatUnit(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RepeatUnit", str)
	}
	return nil
}

func (e RepeatUnit) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ScheduleType string

const (
	ScheduleTypeWeekly   ScheduleType = "WEEKLY"
	ScheduleTypeMonthly  ScheduleType = "MONTHLY"
	ScheduleTypePeriodic ScheduleType = "PERIODIC"
)

var AllScheduleType = []ScheduleType{
	ScheduleTypeWeekly,
	ScheduleTypeMonthly,
	ScheduleTypePeriodic,
}

func (e ScheduleType) IsValid() bool {
	switch e {
	case ScheduleTypeWeekly, ScheduleTypeMonthly, ScheduleTypePeriodic:
		return true
	}
	return false
}

func (e ScheduleType) String() string {
	return string(e)
}

func (e *ScheduleType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ScheduleType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ScheduleType", str)
	}
	return nil
}

func (e ScheduleType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Weekday string

const (
	WeekdayMonday    Weekday = "MONDAY"
	WeekdayTuesday   Weekday = "TUESDAY"
	WeekdayWednesday Weekday = "WEDNESDAY"
	WeekdayThursday  Weekday = "THURSDAY"
	WeekdayFriday    Weekday = "FRIDAY"
	WeekdaySaturday  Weekday = "SATURDAY"
	WeekdaySunday    Weekday = "SUNDAY"
)

var AllWeekday = []Weekday{
	WeekdayMonday,
	WeekdayTuesday,
	WeekdayWednesday,
	WeekdayThursday,
	WeekdayFriday,
	WeekdaySaturday,
	WeekdaySunday,
}

func (e Weekday) IsValid() bool {
	switch e {
	case WeekdayMonday, WeekdayTuesday, WeekdayWednesday, WeekdayThursday, WeekdayFriday, WeekdaySaturday, WeekdaySunday:
		return true
	}
	return false
}

func (e Weekday) String() string {
	return string(e)
}

func (e *Weekday) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Weekday(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Weekday", str)
	}
	return nil
}

func (e Weekday) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
