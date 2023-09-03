package auth

import (
	"context"
	"errors"

	"github.com/infamous55/habit-tracker/internal/models"
)

func ExtractUserFromContext(ctx context.Context) (*models.User, error) {
	err := errors.New("no user in context")

	if ctx.Value(userKey) == nil {
		return nil, err
	}

	user, ok := ctx.Value(userKey).(*models.User)
	if !ok {
		return nil, err
	}

	return user, nil
}
