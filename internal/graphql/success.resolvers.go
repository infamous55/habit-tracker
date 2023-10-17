package graphql

import (
	"context"
	"fmt"

	"github.com/infamous55/habit-tracker/internal/models"
)

func (r *mutationResolver) CreateSuccess(
	ctx context.Context,
	input models.NewSuccess,
) (*models.Success, error) {
	panic(fmt.Errorf("not implemented: CreateSuccess - createSuccess"))
}

func (r *mutationResolver) DeleteSuccess(ctx context.Context, id string) (*models.Success, error) {
	panic(fmt.Errorf("not implemented: DeleteSuccess - deleteSuccess"))
}

type successResolver struct{ *Resolver }

func (r *Resolver) Success() SuccessResolver { return &successResolver{r} }

func (r *successResolver) ID(ctx context.Context, obj *models.Success) (string, error) {
	return obj.ID.Hex(), nil
}

func (r *successResolver) Habit(ctx context.Context, obj *models.Success) (*models.Habit, error) {
	return r.Database.GetHabitByID(obj.HabitID)
}
