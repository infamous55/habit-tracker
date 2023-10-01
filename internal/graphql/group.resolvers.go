package graphql

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/infamous55/habit-tracker/internal/auth"
	"github.com/infamous55/habit-tracker/internal/models"
)

func (r *queryResolver) GetGroups(ctx context.Context) ([]*models.Group, error) {
	user, err := auth.ExtractUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.Database.GetGroupsByUserID(user.ID)
}

func (r *queryResolver) GetGroup(ctx context.Context, id string) (*models.Group, error) {
	user, err := auth.ExtractUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	groupID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	group, err := r.Database.GetGroupByID(groupID)
	if err != nil {
		return nil, err
	}

	if group.UserID != user.ID {
		return nil, fmt.Errorf("permission denied")
	}

	return group, nil
}

func (r *mutationResolver) CreateGroup(
	ctx context.Context,
	input models.NewGroup,
) (*models.Group, error) {
	user, err := auth.ExtractUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	input.UserID = user.ID
	return r.Database.CreateGroup(input)
}

func (r *mutationResolver) UpdateGroup(
	ctx context.Context,
	input models.GroupData,
) (*models.Group, error) {
	user, err := auth.ExtractUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	groupID, err := primitive.ObjectIDFromHex(input.ID)
	if err != nil {
		return nil, err
	}

	group, err := r.Database.GetGroupByID(groupID)
	if err != nil {
		return nil, err
	}

	if group.UserID != user.ID {
		return nil, fmt.Errorf("permission denied")
	}

	// this is the only place where the group id is of type string
	return r.Database.UpdateGroup(input)
}

func (r *mutationResolver) DeleteGroup(ctx context.Context, id string) (*models.Group, error) {
	user, err := auth.ExtractUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	groupID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	group, err := r.Database.GetGroupByID(groupID)
	if err != nil {
		return nil, err
	}

	if group.UserID != user.ID {
		return nil, fmt.Errorf("permission denied")
	}

	return r.Database.DeleteGroupByID(groupID)
}

type groupResolver struct{ *Resolver }

func (r *Resolver) Group() GroupResolver { return &groupResolver{r} }

func (r *groupResolver) ID(ctx context.Context, obj *models.Group) (string, error) {
	return obj.ID.Hex(), nil
}

func (r *groupResolver) Habits(ctx context.Context, obj *models.Group) ([]*models.Habit, error) {
	return r.Database.GetHabitsByGroupID(obj.ID)
}

func (r *groupResolver) User(ctx context.Context, obj *models.Group) (*models.User, error) {
	return r.Database.GetUserByID(obj.UserID)
}
