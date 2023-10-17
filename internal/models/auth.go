package models

type AuthData struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
