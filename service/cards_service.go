package service

type CardsService interface{}

type cardsService struct{}

func NewCardsService() CardsService {
	return &cardsService{}
}

func (s *cardsService) FindById(cardsId int) error {
	return nil
}
