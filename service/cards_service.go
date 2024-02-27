package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jirawan-chuapradit/cards_assignment/config"
	"github.com/jirawan-chuapradit/cards_assignment/models"
	"github.com/jirawan-chuapradit/cards_assignment/models/request"
	"github.com/jirawan-chuapradit/cards_assignment/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CardsService interface {
	FindById(ctx context.Context, cardsId primitive.ObjectID) (models.CardsDetail, error)
	FindAll(ctx context.Context) ([]models.CardsDetail, error)
	Create(ctx context.Context, cardReq request.CreateCardRequestBody) (models.CardsDetail, error)
	Update(ctx context.Context, cardReq request.UpdateCardRequestBody) error
	Store(ctx context.Context, cardsId primitive.ObjectID) error
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
	now := time.Now().In(config.Location)
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

// update card
func (s *cardsService) Update(ctx context.Context, cardReq request.UpdateCardRequestBody) error {
	// TODO: validate created  equal user session

	// repository
	now := time.Now().In(config.Location)
	cardReq.UpdatedAt = &now
	return s.cardsRepository.Update(ctx, cardReq)
}

// store card in an archive
func (s *cardsService) Store(ctx context.Context, cardsId primitive.ObjectID) error {
	// TODO: validate archive  equal user session

	return s.cardsRepository.Store(ctx, cardsId)
}
