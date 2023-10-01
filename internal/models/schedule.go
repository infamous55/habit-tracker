package models

import (
	"fmt"
	"io"
	"strconv"
)

type Schedule struct {
	Type           ScheduleType `json:"type"                      bson:"type"`
	Weekdays       []Weekday    `json:"weekdays,omitempty"        bson:"weekdays,omitempty"`
	Monthdays      []int        `json:"monthdays,omitempty"       bson:"monthdays,omitempty"`
	RepeatInterval *int         `json:"repeat_interval,omitempty" bson:"repeat_interval,omitempty"`
	RepeatUnit     *RepeatUnit  `json:"repeat_unit,omitempty"     bson:"repeat_unit,omitempty"`
	Start          string       `json:"start"                     bson:"start"`
}

// gqlgen does not allow using objects as inputs; only scalars, enums, and input_objects work
type ScheduleInput struct {
	Type           ScheduleType `json:"type"                      bson:"type"`
	Weekdays       []Weekday    `json:"weekdays,omitempty"        bson:"weekdays,omitempty"`
	Monthdays      []int        `json:"monthdays,omitempty"       bson:"monthdays,omitempty"`
	RepeatInterval *int         `json:"repeat_interval,omitempty" bson:"repeat_interval,omitempty"`
	RepeatUnit     *RepeatUnit  `json:"repeat_unit,omitempty"     bson:"repeat_unit,omitempty"`
	Start          string       `json:"start"                     bson:"start"`
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
	case WeekdayMonday,
		WeekdayTuesday,
		WeekdayWednesday,
		WeekdayThursday,
		WeekdayFriday,
		WeekdaySaturday,
		WeekdaySunday:
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
