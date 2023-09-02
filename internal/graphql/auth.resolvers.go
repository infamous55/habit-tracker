package graphql

import (
	"context"
	"fmt"

	"github.com/infamous55/habit-tracker/internal/auth"
	"github.com/infamous55/habit-tracker/internal/models"
)

func (r *mutationResolver) Register(
	ctx context.Context,
	input models.Credentials,
) (*models.AuthData, error) {
	user, err := r.Database.CreateUser(input)
	if err != nil {
		return nil, err
	}

	token, err := auth.NewJWTWithCustomClaims(user.ID)
	if err != nil {
		return nil, err
	}

	return &models.AuthData{
		Token: token,
		User:  user,
	}, nil
}

func (r *mutationResolver) Login(
	ctx context.Context,
	input models.Credentials,
) (*models.AuthData, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

func (r *mutationResolver) RefreshToken(
	ctx context.Context,
	input models.RefreshTokenInput,
) (*models.AuthData, error) {
	panic(fmt.Errorf("not implemented: RefreshToken - refreshToken"))
}
