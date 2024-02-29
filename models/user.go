package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Login struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}
type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserName  string             `json:"username" bson:"username"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
