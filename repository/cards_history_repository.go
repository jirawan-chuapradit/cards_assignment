package repository

import (
	"context"

	"github.com/jirawan-chuapradit/cards_assignment/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CardsHistoryRepository interface {
	FindHistoryByCardsId(ctx context.Context, cardsId primitive.ObjectID) ([]models.CardsHistory, error)
}

type cardsHistoryRepository struct {
	CardsAssignmentDatabase mongo.Database
}

func NewCardsHistoryRepository(conn *mongo.Client) CardsHistoryRepository {
	return &cardsHistoryRepository{
		CardsAssignmentDatabase: *conn.Database("cards_assignment"),
	}
}

/// find card history by cards Id
func (r *cardsHistoryRepository) FindHistoryByCardsId(ctx context.Context, cardsId primitive.ObjectID) ([]models.CardsHistory, error) {
	cursor, err := r.CardsAssignmentDatabase.Collection("cards_history").Find(ctx, map[string]interface{}{"card_id": cardsId})
	if err != nil {
		return []models.CardsHistory{}, err
	}
	history := []models.CardsHistory{}
	if err := cursor.All(ctx, &history); err != nil {
		return []models.CardsHistory{}, err
	}
	return history, nil
}
