package repository

import (
	"context"
	"log"

	"github.com/jirawan-chuapradit/cards_assignment/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CardsRepository interface {
	FindById(ctx context.Context, cardsId primitive.ObjectID) (cardsDetails models.CardsDetail, err error)
}

type cardsRepository struct {
	CardsAssignmentDatabase mongo.Database
}

func NewCardsRepository(conn *mongo.Client) CardsRepository {
	return &cardsRepository{
		CardsAssignmentDatabase: *conn.Database("cards_assignment"),
	}
}

// find card by cards Id
func (r *cardsRepository) FindById(ctx context.Context, cardsId primitive.ObjectID) (cardsDetails models.CardsDetail, err error) {
	result := r.CardsAssignmentDatabase.Collection("cards").FindOne(ctx, map[string]interface{}{"_id": cardsId})
	if err = result.Decode(&cardsDetails); err != nil {
		return
	}
	log.Println(cardsDetails)
	return
}
