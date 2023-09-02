package mongodb

import (
	"context"
	"fmt"
	"net/mail"
	"strings"
	"time"

	"github.com/infamous55/habit-tracker/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (db DatabaseWrapper) GetUserById(id string) (*models.User, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := db.Collection("users").FindOne(ctx, bson.M{"_id": ID})

	var user *models.User
	err = result.Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (db DatabaseWrapper) CreateUser(credentials models.Credentials) (*models.User, error) {
	_, err := mail.ParseAddress(credentials.Email)
	if err != nil {
		return nil, err
	}

	name := strings.Split(credentials.Email, "@")[0]
	hashedPassword, err := hashPassword(credentials.Password)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user := models.User{
		Name:     name,
		Email:    credentials.Email,
		Password: hashedPassword,
	}
	result, err := db.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("inserted id error")
	}

	user.ID = oid.Hex()
	return &user, nil
}
