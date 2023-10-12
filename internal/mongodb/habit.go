package mongodb

import (
	"context"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/infamous55/habit-tracker/internal/models"
)

func (db *DatabaseWrapper) GetHabitsByUserID(userID primitive.ObjectID) ([]*models.Habit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := db.Collection("habits").Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []*models.Habit
	err = cur.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (db *DatabaseWrapper) GetHabitsByGroupID(groupID primitive.ObjectID) ([]*models.Habit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := db.Collection("habits").Find(ctx, bson.M{"group_id": groupID})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []*models.Habit
	err = cur.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (db *DatabaseWrapper) GetHabitByID(id primitive.ObjectID) (*models.Habit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := db.Collection("habits").FindOne(ctx, bson.M{"_id": id})

	var habit *models.Habit
	err := result.Decode(&habit)
	if err != nil {
		return nil, err
	}
	return habit, nil
}

func (db *DatabaseWrapper) CreateHabit(data models.HabitCreate) (*models.Habit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := db.Collection("habits").InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}

	insertedID := result.InsertedID.(primitive.ObjectID)
	return &models.Habit{
		ID:          insertedID,
		Name:        data.Name,
		Description: data.Description,
		Schedule:    models.Schedule(data.Schedule),
		GroupID:     data.GroupID,
		UserID:      data.UserID,
	}, nil
}

func (db *DatabaseWrapper) UpdateHabit(data models.HabitUpdate) (*models.Habit, error) {
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
	if data.Schedule != nil {
		update = append(
			update,
			bson.E{Key: "$set", Value: bson.D{{Key: "schedule", Value: *data.Schedule}}},
		)
	}
	if data.GroupID != nil {
		update = append(
			update,
			bson.E{Key: "$set", Value: bson.D{{Key: "group_id", Value: *data.GroupID}}},
		)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.Collection("habits").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return db.GetHabitByID(data.ID)
}

func (db *DatabaseWrapper) DeleteHabitByID(id primitive.ObjectID) (*models.Habit, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := db.Collection("habits").FindOneAndDelete(ctx, filter)
	err := result.Err()
	if err != nil {
		return nil, err
	}

	var deletedHabit models.Habit
	err = result.Decode(&deletedHabit)
	if err != nil {
		return nil, err
	}
	return &deletedHabit, nil
}

func (db *DatabaseWrapper) GetHabitsWithFilter(
	options models.HabitFilterOptions,
) ([]*models.Habit, error) {
	filter := bson.M{
		"user_id": options.UserID,
	}

	if options.GroupID != nil {
		filter["group_id"] = *options.GroupID
	}

	if options.StartDate != nil && options.EndDate != nil {
		filter["$and"] = []bson.M{
			{
				"schedule.type": "WEEKLY",
				"schedule.weekdays": bson.M{
					"$elemMatch": bson.M{
						"$in": weekdaysWithinInterval(*options.StartDate, *options.EndDate),
					},
				},
			},
			{
				"schedule.type": "MONTHLY",
				"schedule.monthdays": bson.M{
					"$elemMatch": bson.M{
						"$gte": options.StartDate.Day(),
						"$lte": options.EndDate.Day(),
					},
				},
			},
			{
				"schedule.type":  "PERIODIC",
				"schedule.start": bson.M{"$lte": *options.EndDate},
				"$or": []bson.M{
					{
						"schedule.repeat_unit": "DAY",
						"schedule.repeat_interval": bson.M{
							"$gte": int(options.StartDate.Sub(*options.EndDate).Hours() / 24),
						},
					},
					{
						"schedule.repeat_unit": "WEEK",
						"schedule.repeat_interval": bson.M{
							"$gte": int(options.StartDate.Sub(*options.EndDate).Hours() / 24 / 7),
						},
					},
					{
						"schedule.repeat_unit": "MONTH",
						"schedule.repeat_interval": bson.M{
							"$gte": int(options.StartDate.Sub(*options.EndDate).Hours() / 24 / 30),
						},
					},
				},
			},
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := db.Collection("habits").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []*models.Habit
	err = cur.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func weekdaysWithinInterval(start, end time.Time) []string {
	weekdays := make(map[string]struct{})
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		if d.Weekday() != time.Saturday && d.Weekday() != time.Sunday {
			weekday := strings.ToUpper(d.Weekday().String())
			weekdays[weekday] = struct{}{}
		}
	}

	return mapKeys(weekdays)
}

func mapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}
