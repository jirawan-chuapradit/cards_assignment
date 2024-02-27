package request

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateCommentBody struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Description *string            `json:"description" bson:"description"`
	UpdatedAt   *time.Time         `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
