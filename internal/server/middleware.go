package server

import (
	"context"

	gqlgen "github.com/99designs/gqlgen/graphql"
	"github.com/infamous55/habit-tracker/internal/config"
	"github.com/labstack/echo"
)

const (
	playgroundPasswordKey string = "playgroundPassword"
)

func extractHeaders(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		req := ctx.Request()

		playgroundPasswordHeader := req.Header.Get("Playground-Password")

		ctx.Set(playgroundPasswordKey, playgroundPasswordHeader)

		return next(ctx)
	}
}

func operationMiddleware(ctx context.Context, next gqlgen.OperationHandler) gqlgen.ResponseHandler {
	playgroundPassword, ok := ctx.Value(playgroundPasswordKey).(string)
	if !ok || playgroundPassword != config.GetSecret("GRAPHQL_PLAYGROUND_PASSWORD") {
		gqlgen.GetOperationContext(ctx).DisableIntrospection = true
	}
	return next(ctx)
}
