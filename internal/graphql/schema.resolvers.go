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

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type (
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
)
