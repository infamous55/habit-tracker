package models

type Group struct {
	ID          string  `json:"id"                    bson:"_id"`
	Name        string  `json:"name"                  bson:"name"`
	Description *string `json:"description,omitempty" bson:"description,omitempty"`
	UserID      string  `json:"user_id"               bson:"user_id"`
}
