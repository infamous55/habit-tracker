package ctxbridge

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	ctx context.Context
}

type ContextKey string

const (
	EchoContextKey ContextKey = "echoContextKey"
)

func EchoContextToContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), EchoContextKey, c)
		c.SetRequest(c.Request().WithContext(ctx))

		cc := &CustomContext{c, ctx}
		return next(cc)
	}
}

func EchoContextFromContext(ctx context.Context) (echo.Context, error) {
	echoContext := ctx.Value(EchoContextKey)
	if echoContext == nil {
		return nil, fmt.Errorf("could not retrieve echo.Context")
	}

	ec, ok := echoContext.(echo.Context)
	if !ok {
		return nil, fmt.Errorf("echo.Context has the wrong type")
	}
	return ec, nil
}
