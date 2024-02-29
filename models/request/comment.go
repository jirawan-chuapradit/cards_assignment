package request

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateCommentBody struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" binding:"required"`
	Description *string            `json:"description" bson:"description" binding:"required"`
	UpdatedBy   string             `json:"updated_by,omitempty" bson:"updated_by,omitempty"`
	UpdatedAt   *time.Time         `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type CreateCommentBody struct {
	CardID      primitive.ObjectID `json:"card_id,omitempty" binding:"required"`
	Description *string            `json:"description" binding:"required"`
}
