package auth

import (
	"context"
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"github.com/infamous55/habit-tracker/internal/ctxbridge"
	"github.com/infamous55/habit-tracker/internal/models"
)

func TestExtractUserFromEchoContext(t *testing.T) {
	e := echo.New()
	req := &http.Request{}
	res := &echo.Response{}
	ec := e.NewContext(req, res)

	user := &models.User{}
	ec.Set(userKey, user)

	extractedUser, err := ExtractUserFromEchoContext(ec)
	assert.Nil(t, err)
	assert.Equal(t, user, extractedUser)
}

func TestExtractUserFromContext(t *testing.T) {
	e := echo.New()
	req := &http.Request{}
	rec := &echo.Response{}
	ec := e.NewContext(req, rec)

	user := &models.User{}
	ec.Set(userKey, user)

	ctx := context.WithValue(ec.Request().Context(), ctxbridge.EchoContextKey, ec)

	extractedUser, err := ExtractUserFromContext(ctx)
	assert.Nil(t, err)
	assert.Equal(t, user, extractedUser)
}
