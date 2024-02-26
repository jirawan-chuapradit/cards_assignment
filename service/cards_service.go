package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jirawan-chuapradit/cards_assignment/models"
	"github.com/jirawan-chuapradit/cards_assignment/models/request"
	"github.com/jirawan-chuapradit/cards_assignment/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CardsService interface {
	FindById(ctx context.Context, cardsId primitive.ObjectID) (models.CardsDetail, error)
	FindAll(ctx context.Context) ([]models.CardsDetail, error)
	Create(ctx context.Context, cardReq request.CreateCardRequestBody) (models.CardsDetail, error)
}

type cardsService struct {
	cardsRepository repository.CardsRepository
}

func NewCardsService(cardsRepository repository.CardsRepository) CardsService {
	return &cardsService{
		cardsRepository: cardsRepository,
	}
}

// find card by id
func (s *cardsService) FindById(ctx context.Context, cardsId primitive.ObjectID) (models.CardsDetail, error) {
	// repository
	return s.cardsRepository.FindById(ctx, cardsId)
}

// find cards
func (s *cardsService) FindAll(ctx context.Context) ([]models.CardsDetail, error) {
	// repository
	return s.cardsRepository.FindAll(ctx)
}

// create card
func (s *cardsService) Create(ctx context.Context, cardReq request.CreateCardRequestBody) (models.CardsDetail, error) {
	var card models.CardsDetail
	now := time.Now()
	crb, err := json.Marshal(cardReq)
	if err != nil {
		return card, err
	}
	if err := json.Unmarshal(crb, &card); err != nil {
		return card, err
	}
	card.CreatedAt = &now
	// TODO: assign created at by get user from session
	card.CreatedBy = "mock user"
	// repository
	return s.cardsRepository.Create(ctx, card)
}
