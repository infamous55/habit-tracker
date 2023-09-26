package graphql

import (
	"context"
	"fmt"

	"github.com/infamous55/habit-tracker/internal/models"
)

func (r *queryResolver) GetGroups(ctx context.Context) ([]*models.Group, error) {
	panic(fmt.Errorf("not implemented: GetGroups - getGroups"))
}

func (r *queryResolver) GetGroup(ctx context.Context, id string) (*models.Group, error) {
	panic(fmt.Errorf("not implemented: GetGroup - getGroup"))
}

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

type groupResolver struct{ *Resolver }

func (r *Resolver) Group() GroupResolver { return &groupResolver{r} }

func (r *groupResolver) Habits(ctx context.Context, obj *models.Group) ([]*models.Habit, error) {
	panic(fmt.Errorf("not implemented: Habits - habits"))
}

func (r *groupResolver) User(ctx context.Context, obj *models.Group) (*models.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}
