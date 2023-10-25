package ctxbridge

import (
	"context"
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestEchoContextToContext(t *testing.T) {
	e := echo.New()
	req := &http.Request{}
	res := &echo.Response{}
	ec := e.NewContext(req, res)

	mockHandlerWasCalled := false
	mockHandler := func(c echo.Context) error {
		mockHandlerWasCalled = true
		return nil
	}

	err := EchoContextToContext(mockHandler)(ec)
	assert.Nil(t, err)

	assert.Equal(t, ec.Request().Context().Value(echoContextKey), ec)
	assert.True(t, mockHandlerWasCalled)
}

func TestEchoContextFromContext(t *testing.T) {
	e := echo.New()
	req := &http.Request{}
	rec := &echo.Response{}
	ec := e.NewContext(req, rec)

	ctx := context.WithValue(ec.Request().Context(), echoContextKey, ec)

	echoContext, err := EchoContextFromContext(ctx)
	assert.Nil(t, err)
	assert.Equal(t, ec, echoContext)
}
