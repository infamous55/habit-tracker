package auth

import (
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/infamous55/habit-tracker/internal/mongodb"
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
			claims, err := ParseJWTWithCustomClaims(tokenString)
			if err != nil {
				return next(ctx)
			}

			user, err := db.GetUserByID(claims.UserID)
			if err != nil {
				return next(ctx)
			}

			ctx.Set(userKey, user)
			return next(ctx)
		}
	}
}
