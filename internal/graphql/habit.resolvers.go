package graphql

import (
	"context"
	"fmt"

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
	panic(fmt.Errorf("not implemented: GetHabit - getHabit"))
}

func (r *mutationResolver) CreateHabit(
	ctx context.Context,
	input models.NewHabit,
) (*models.Habit, error) {
	panic(fmt.Errorf("not implemented: CreateHabit - createHabit"))
}

func (r *mutationResolver) UpdateHabit(
	ctx context.Context,
	input models.HabitData,
) (*models.Habit, error) {
	panic(fmt.Errorf("not implemented: UpdateHabit - updateHabit"))
}

func (r *mutationResolver) DeleteHabit(ctx context.Context, id string) (*models.Habit, error) {
	panic(fmt.Errorf("not implemented: DeleteHabit - deleteHabit"))
}

type habitResolver struct{ *Resolver }

func (r *Resolver) Habit() HabitResolver { return &habitResolver{r} }

func (r *habitResolver) ID(ctx context.Context, obj *models.Habit) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
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
