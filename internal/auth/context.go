package auth

import (
	"context"
	"fmt"

	"github.com/labstack/echo"

	"github.com/infamous55/habit-tracker/internal/ctxbridge"
	"github.com/infamous55/habit-tracker/internal/models"
)

func ExtractUserFromEchoContext(ec echo.Context) (*models.User, error) {
	err := fmt.Errorf("no user in context")

	if ec.Get(userKey) == nil {
		return nil, err
	}

	user, ok := ec.Get(userKey).(*models.User)
	if !ok {
		return nil, err
	}

	return user, nil
}

func ExtractUserFromContext(ctx context.Context) (*models.User, error) {
	ec, err := ctxbridge.EchoContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return ExtractUserFromEchoContext(ec)
}
