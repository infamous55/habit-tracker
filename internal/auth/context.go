package auth

import (
	"fmt"

	"github.com/infamous55/habit-tracker/internal/models"
	"github.com/labstack/echo"
)

func ExtractUserFromEchoContext(ctx echo.Context) (*models.User, error) {
	err := fmt.Errorf("no user in context")

	if ctx.Get(userKey) == nil {
		return nil, err
	}

	user, ok := ctx.Get(userKey).(*models.User)
	if !ok {
		return nil, err
	}

	return user, nil
}
