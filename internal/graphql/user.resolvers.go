package graphql

import (
	"context"

	"github.com/infamous55/habit-tracker/internal/auth"
	"github.com/infamous55/habit-tracker/internal/models"
)

func (r *queryResolver) GetCurrentUser(ctx context.Context) (*models.User, error) {
	return auth.ExtractUserFromContext(ctx)
}

type userResolver struct{ *Resolver }

func (r *Resolver) User() UserResolver { return &userResolver{r} }

func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	return obj.ID.Hex(), nil
}

func (r *userResolver) Groups(ctx context.Context, obj *models.User) ([]*models.Group, error) {
	return r.Database.GetGroupsByUserID(obj.ID)
}

func (r *userResolver) Habits(ctx context.Context, obj *models.User) ([]*models.Habit, error) {
	return r.Database.GetHabitsByUserID(obj.ID)
}
