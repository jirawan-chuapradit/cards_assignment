package request

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateCardRequestBody struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

type UpdateCardRequestBody struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty" binding:"required"`
	Title       string             `json:"title" bson:"title" binding:"required"`
	Description string             `json:"description" bson:"description" binding:"required"`
	Status      string             `json:"status" bson:"status" binding:"required"`
	UpdatedBy   string             `json:"updated_by,omitempty" bson:"updated_by,omitempty"`
	UpdatedAt   *time.Time         `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
