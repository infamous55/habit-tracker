package server

import (
	"context"
	"os"

	gqlgen "github.com/99designs/gqlgen/graphql"
	"github.com/labstack/echo"
)

const (
	playgroundPasswordKey string = "playgroundPassword"
)

func extractPlaygroundPassword(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		playgroundPasswordHeader := ctx.Request().Header.Get("Playground-Password")

		ctx.Set(playgroundPasswordKey, playgroundPasswordHeader)

		return next(ctx)
	}
}

func verifyPlaygroundPassword(
	ctx context.Context,
	next gqlgen.OperationHandler,
) gqlgen.ResponseHandler {
	playgroundPassword, ok := ctx.Value(playgroundPasswordKey).(string)
	if !ok || playgroundPassword != os.Getenv("GRAPHQL_PLAYGROUND_PASSWORD") {
		gqlgen.GetOperationContext(ctx).DisableIntrospection = true
	}
	return next(ctx)
}
