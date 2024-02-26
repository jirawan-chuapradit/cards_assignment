package repository

import (
	"context"
	"log"

	"github.com/jirawan-chuapradit/cards_assignment/models"
	"github.com/jirawan-chuapradit/cards_assignment/models/request"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CardsRepository interface {
	FindById(ctx context.Context, cardsId primitive.ObjectID) (cardsDetails models.CardsDetail, err error)
	FindAll(ctx context.Context) ([]models.CardsDetail, error)
	Create(ctx context.Context, card models.CardsDetail) (models.CardsDetail, error)
	Update(ctx context.Context, cardReq request.UpdateCardRequestBody) error
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

// find cards
func (r *cardsRepository) FindAll(ctx context.Context) ([]models.CardsDetail, error) {
	var cards []models.CardsDetail
	cursor, err := r.CardsAssignmentDatabase.Collection("cards").Find(ctx, map[string]interface{}{})
	if err != nil {
		return cards, err
	}

	if err := cursor.All(ctx, &cards); err != nil {
		return cards, err
	}
	return cards, nil
}

// create card
func (r *cardsRepository) Create(ctx context.Context, card models.CardsDetail) (models.CardsDetail, error) {
	result, err := r.CardsAssignmentDatabase.Collection("cards").InsertOne(ctx, card)
	if err != nil {
		return card, err
	}

	card.ID = result.InsertedID.(primitive.ObjectID)
	return card, nil
}

// update card
func (r *cardsRepository) Update(ctx context.Context, cardReq request.UpdateCardRequestBody) error {
	filter := map[string]interface{}{
		"_id": cardReq.ID,
	}
	update := map[string]interface{}{
		"$set": cardReq,
	}

	return r.CardsAssignmentDatabase.Collection("cards").FindOneAndUpdate(ctx, filter, update).Err()
}
