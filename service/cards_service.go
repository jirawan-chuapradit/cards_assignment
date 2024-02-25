package service

import "github.com/jirawan-chuapradit/cards_assignment/models"

type CardsService interface {
	FindById(cardsId int) (models.CardsDetail, error)
}

type cardsService struct{}

func NewCardsService() CardsService {
	return &cardsService{}
}

func (s *cardsService) FindById(cardsId int) (models.CardsDetail, error) {
	// repository

	return models.CardsDetail{}, nil
}
