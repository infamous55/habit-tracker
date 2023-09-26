package graphql

import (
	"context"
	"fmt"

	"github.com/infamous55/habit-tracker/internal/models"
)

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

func (r *mutationResolver) CreateSuccess(
	ctx context.Context,
	input models.NewSuccess,
) (*models.Success, error) {
	panic(fmt.Errorf("not implemented: CreateSuccess - createSuccess"))
}

func (r *mutationResolver) DeleteSuccess(ctx context.Context, id string) (*models.Success, error) {
	panic(fmt.Errorf("not implemented: DeleteSuccess - deleteSuccess"))
}

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

func (r *Resolver) Habit() HabitResolver { return &habitResolver{r} }

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type (
	habitResolver    struct{ *Resolver }
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
)
