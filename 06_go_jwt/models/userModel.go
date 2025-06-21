package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// *string  can be nil or point to a string value

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	Name          *string            `json:"name" validate:"required, min=2, max=100"`
	Password      *string            `json:"password" validate:"required"`
	Email         *string            `json:"email" validate:"email, required"`
	Token         *string            `json:"token"`
	User_type     *string            `json:"user_type" validate:"reqired, eq=ADMIN|eq=USER"`
	Refresh_token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
}
