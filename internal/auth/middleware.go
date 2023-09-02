package auth

import (
	"strings"

	"github.com/infamous55/habit-tracker/internal/mongodb"
	"github.com/labstack/echo"
)

const (
	userKey string = "user"
)

func ExtractUserMiddleware(db mongodb.DatabaseWrapper) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			authTokenHeader := ctx.Request().Header.Get("Authorization")
			if authTokenHeader == "" || !strings.HasPrefix(authTokenHeader, "Bearer ") {
				return next(ctx)
			}

			tokenString := strings.TrimPrefix(authTokenHeader, "Bearer ")
			claims, err := ParseWithCustomClaims(tokenString)
			if err != nil {
				return next(ctx)
			}

			user, err := db.GetUserById(claims.UserId)
			if err != nil {
				return next(ctx)
			}

			ctx.Set(userKey, user)
			return next(ctx)
		}
	}
}
