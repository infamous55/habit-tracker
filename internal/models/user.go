package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       primitive.ObjectID `json:"id"       bson:"_id,omitempty"`
	Name     string             `json:"name"     bson:"name,omitempty"`
	Email    string             `json:"email"    bson:"email,omitempty"`
	Password string             `json:"password" bson:"password,omitempty"`
}

func (u *User) ComparePassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
