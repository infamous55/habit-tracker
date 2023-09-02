package graphql

import (
	"context"
	"fmt"

	"github.com/infamous55/habit-tracker/internal/models"
)

func (r *mutationResolver) CreateGroup(
	ctx context.Context,
	input models.NewGroup,
) (*models.Group, error) {
	panic(fmt.Errorf("not implemented: CreateGroup - createGroup"))
}

func (r *mutationResolver) UpdateGroup(
	ctx context.Context,
	input models.GroupData,
) (*models.Group, error) {
	panic(fmt.Errorf("not implemented: UpdateGroup - updateGroup"))
}

func (r *mutationResolver) DeleteGroup(ctx context.Context, id string) (*models.Group, error) {
	panic(fmt.Errorf("not implemented: DeleteGroup - deleteGroup"))
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

func (r *queryResolver) GetCurrentUser(ctx context.Context) (*models.User, error) {
	panic(fmt.Errorf("not implemented: GetCurrentUser - getCurrentUser"))
}

func (r *queryResolver) GetGroups(ctx context.Context) ([]*models.Group, error) {
	panic(fmt.Errorf("not implemented: GetGroups - getGroups"))
}

func (r *queryResolver) GetGroup(ctx context.Context, id string) (*models.Group, error) {
	panic(fmt.Errorf("not implemented: GetGroup - getGroup"))
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

func (r *userResolver) Groups(ctx context.Context, obj *models.User) ([]*models.Group, error) {
	panic(fmt.Errorf("not implemented: Groups - groups"))
}

func (r *userResolver) Habits(ctx context.Context, obj *models.User) ([]*models.Habit, error) {
	panic(fmt.Errorf("not implemented: Habits - habits"))
}

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

func (r *Resolver) User() UserResolver { return &userResolver{r} }

type (
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
	userResolver     struct{ *Resolver }
)
