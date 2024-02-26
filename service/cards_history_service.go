package service

import (
	"context"

	"github.com/jirawan-chuapradit/cards_assignment/models"
	"github.com/jirawan-chuapradit/cards_assignment/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CardsHistoryService interface {
	FindHistoryById(ctx context.Context, cardsId primitive.ObjectID) ([]models.CardsHistory, error)
}

type cardsHistoryService struct {
	cardsHistory repository.CardsHistoryRepository
}

func NewCardsHistoryService(cardsHistory repository.CardsHistoryRepository) CardsHistoryService {
	return &cardsHistoryService{
		cardsHistory: cardsHistory,
	}
}

func (s *cardsHistoryService) FindHistoryById(ctx context.Context, cardsId primitive.ObjectID) ([]models.CardsHistory, error) {
	// repository
	return s.cardsHistory.FindHistoryByCardsId(ctx, cardsId)
}
