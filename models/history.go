package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CardsHistory struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Status      string             `json:"status" bson:"status"`
	CardID      primitive.ObjectID `json:"card_id" bson:"card_id"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
}
