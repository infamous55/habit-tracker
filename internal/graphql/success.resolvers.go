package graphql

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/infamous55/habit-tracker/internal/auth"
	"github.com/infamous55/habit-tracker/internal/models"
)

func (r *mutationResolver) CreateSuccess(
	ctx context.Context,
	input models.NewSuccess,
) (*models.Success, error) {
	user, err := auth.ExtractUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	habitID, err := primitive.ObjectIDFromHex(input.HabitID)
	if err != nil {
		return nil, err
	}

	habit, err := r.Database.GetHabitByID(habitID)
	if err != nil {
		return nil, err
	}

	if habit.UserID != user.ID {
		return nil, fmt.Errorf("permission denied")
	}

	ok := habit.IsScheduled(input.Date)
	if !ok {
		return nil, fmt.Errorf("invalid date")
	}

	data := models.SuccessCreate{
		Date:    input.Date,
		HabitID: habitID,
	}
	return r.Database.CreateSuccess(data)
}

func (r *mutationResolver) DeleteSuccess(ctx context.Context, id string) (*models.Success, error) {
	user, err := auth.ExtractUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	successID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	success, err := r.Database.GetSuccessByID(successID)
	if err != nil {
		return nil, err
	}

	habit, err := r.Database.GetHabitByID(success.HabitID)
	if err != nil {
		return nil, err
	}

	if habit.UserID != user.ID {
		return nil, fmt.Errorf("permission denied")
	}

	return r.Database.DeleteSuccessByID(successID)
}

type successResolver struct{ *Resolver }

func (r *Resolver) Success() SuccessResolver { return &successResolver{r} }

func (r *successResolver) ID(ctx context.Context, obj *models.Success) (string, error) {
	return obj.ID.Hex(), nil
}

func (r *successResolver) Habit(ctx context.Context, obj *models.Success) (*models.Habit, error) {
	return r.Database.GetHabitByID(obj.HabitID)
}
