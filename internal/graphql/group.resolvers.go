package graphql

import (
	"context"
	"fmt"

	"github.com/infamous55/habit-tracker/internal/auth"
	"github.com/infamous55/habit-tracker/internal/ctxbridge"
	"github.com/infamous55/habit-tracker/internal/models"
)

func (r *queryResolver) GetGroups(ctx context.Context) ([]*models.Group, error) {
	ec, err := ctxbridge.EchoContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := auth.ExtractUserFromEchoContext(ec)
	if err != nil {
		return nil, err
	}

	return r.Database.GetGroupsByUserID(user.ID)
}

func (r *queryResolver) GetGroup(ctx context.Context, id string) (*models.Group, error) {
	ec, err := ctxbridge.EchoContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	_, err = auth.ExtractUserFromEchoContext(ec)
	if err != nil {
		return nil, err
	}

	return r.Database.GetGroupByID(id)
}

func (r *mutationResolver) CreateGroup(
	ctx context.Context,
	input models.NewGroup,
) (*models.Group, error) {
	ec, err := ctxbridge.EchoContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := auth.ExtractUserFromEchoContext(ec)
	if err != nil {
		return nil, err
	}

	data := models.Group{
		Name:        input.Name,
		Description: input.Description,
		UserID:      user.ID,
	}
	return r.Database.CreateGroup(data)
}

func (r *mutationResolver) UpdateGroup(
	ctx context.Context,
	input models.GroupData,
) (*models.Group, error) {
	ec, err := ctxbridge.EchoContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := auth.ExtractUserFromEchoContext(ec)
	if err != nil {
		return nil, err
	}

	group, err := r.Database.GetGroupByID(input.ID)
	if err != nil {
		return nil, err
	}

	if group.UserID != user.ID {
		return nil, fmt.Errorf("permission denied")
	}

	return r.Database.UpdateGroup(input)
}

func (r *mutationResolver) DeleteGroup(ctx context.Context, id string) (*models.Group, error) {
	ec, err := ctxbridge.EchoContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := auth.ExtractUserFromEchoContext(ec)
	if err != nil {
		return nil, err
	}

	group, err := r.Database.GetGroupByID(id)
	if err != nil {
		return nil, err
	}

	if group.UserID != user.ID {
		return nil, fmt.Errorf("permission denied")
	}

	return r.Database.DeleteGroupByID(id)
}

type groupResolver struct{ *Resolver }

func (r *Resolver) Group() GroupResolver { return &groupResolver{r} }

func (r *groupResolver) Habits(ctx context.Context, obj *models.Group) ([]*models.Habit, error) {
	panic(fmt.Errorf("not implemented: Habits - habits"))
}

func (r *groupResolver) User(ctx context.Context, obj *models.Group) (*models.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}
