package models

type AuthData struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

type Credentials struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
