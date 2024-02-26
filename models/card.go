package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CardsDetail struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Status      string             `json:"status" bson:"status"`
	CreatedBy   string             `json:"created_by" bson:"created_by"`
	CreatedAt   *time.Time         `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   *time.Time         `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Comments    []*Comment         `json:"comments,omitempty" bson:"comments,omitempty"`
}

type Comment struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Img         string             `json:"img" bson:"img"`
	Description string             `json:"description" bson:"description"`
	CreatedBy   string             `json:"created_by" bson:"created_by"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}
