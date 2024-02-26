package service

import (
	"context"

	"github.com/jirawan-chuapradit/cards_assignment/models"
	"github.com/jirawan-chuapradit/cards_assignment/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CardsService interface {
	FindById(ctx context.Context, cardsId primitive.ObjectID) (models.CardsDetail, error)
}

type cardsService struct {
	cardsRepository repository.CardsRepository
}

func NewCardsService(cardsRepository repository.CardsRepository) CardsService {
	return &cardsService{
		cardsRepository: cardsRepository,
	}
}

func (s *cardsService) FindById(ctx context.Context, cardsId primitive.ObjectID) (models.CardsDetail, error) {
	// repository
	return s.cardsRepository.FindById(ctx, cardsId)
}
