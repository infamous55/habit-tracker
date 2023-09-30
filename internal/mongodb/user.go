package mongodb

import (
	"context"
	"fmt"
	"net/mail"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"github.com/infamous55/habit-tracker/internal/models"
)

func hashPassword(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (db *DatabaseWrapper) GetUserByID(id primitive.ObjectID) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := db.Collection("users").FindOne(ctx, bson.M{"_id": id})

	var user *models.User
	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (db *DatabaseWrapper) CreateUser(credentials models.Credentials) (*models.User, error) {
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

	user.ID = oid
	return &user, nil
}

func (db *DatabaseWrapper) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := db.Collection("users").FindOne(ctx, bson.M{"email": email})

	var user *models.User
	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
