package graphql

import (
	"context"

	"github.com/infamous55/habit-tracker/internal/auth"
	"github.com/infamous55/habit-tracker/internal/models"
	"github.com/infamous55/habit-tracker/internal/validator"
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
	err := validator.Validate.Struct(input)
	if err != nil {
		return nil, err
	}

	user, err := r.Database.GetUserByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	err = user.ComparePassword(input.Password)
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

func (r *mutationResolver) RefreshToken(ctx context.Context) (*models.AuthData, error) {
	user, err := auth.ExtractUserFromContext(ctx)
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
