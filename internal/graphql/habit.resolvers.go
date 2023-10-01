package graphql

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/infamous55/habit-tracker/internal/auth"
	"github.com/infamous55/habit-tracker/internal/models"
)

func (r *queryResolver) GetHabits(
	ctx context.Context,
	groupID *string,
	startDate *string,
	endDate *string,
	succeeded *bool,
) ([]*models.Habit, error) {
	panic(fmt.Errorf("not implemented: GetHabits - getHabits"))
}

func (r *queryResolver) GetHabit(ctx context.Context, id string) (*models.Habit, error) {
	user, err := auth.ExtractUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	habitID, err := primitive.ObjectIDFromHex(id)
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

	return habit, nil
}

func (r *mutationResolver) CreateHabit(
	ctx context.Context,
	input models.NewHabit,
) (*models.Habit, error) {
	user, err := auth.ExtractUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	groupID, err := primitive.ObjectIDFromHex(input.GroupID)
	if err != nil {
		return nil, err
	}

	group, err := r.Database.GetGroupByID(groupID)
	if err != nil {
		return nil, err
	}

	if group.UserID != user.ID {
		return nil, fmt.Errorf("bad request")
	}

	data := models.HabitCreate{
		Name:        input.Name,
		Description: input.Description,
		Schedule:    input.Schedule,
		GroupID:     group.ID,
		UserID:      user.ID,
	}
	return r.Database.CreateHabit(data)
}

func (r *mutationResolver) UpdateHabit(
	ctx context.Context,
	input models.HabitData,
) (*models.Habit, error) {
	user, err := auth.ExtractUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	habitID, err := primitive.ObjectIDFromHex(input.ID)
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

	var groupID primitive.ObjectID
	if input.GroupID != nil {
		groupID, err = primitive.ObjectIDFromHex(*input.GroupID)
		if err != nil {
			return nil, err
		}

		group, err := r.Database.GetGroupByID(groupID)
		if err != nil {
			return nil, err
		}

		if group.UserID != user.ID {
			return nil, fmt.Errorf("bad request")
		}
	}

	data := models.HabitUpdate{
		ID:          habitID,
		Name:        input.Name,
		Description: input.Description,
		Schedule:    input.Schedule,
		GroupID:     &groupID,
	}
	return r.Database.UpdateHabit(data)
}

func (r *mutationResolver) DeleteHabit(ctx context.Context, id string) (*models.Habit, error) {
	panic(fmt.Errorf("not implemented: DeleteHabit - deleteHabit"))
}

type habitResolver struct{ *Resolver }

func (r *Resolver) Habit() HabitResolver { return &habitResolver{r} }

func (r *habitResolver) ID(ctx context.Context, obj *models.Habit) (string, error) {
	return obj.ID.Hex(), nil
}

func (r *habitResolver) Schedule(ctx context.Context, obj *models.Habit) (*models.Schedule, error) {
	panic(fmt.Errorf("not implemented: Schedule - schedule"))
}

func (r *habitResolver) Successes(
	ctx context.Context,
	obj *models.Habit,
) ([]*models.Success, error) {
	panic(fmt.Errorf("not implemented: Successes - successes"))
}

func (r *habitResolver) Group(ctx context.Context, obj *models.Habit) (*models.Group, error) {
	panic(fmt.Errorf("not implemented: Group - group"))
}

func (r *habitResolver) User(ctx context.Context, obj *models.Habit) (*models.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}
