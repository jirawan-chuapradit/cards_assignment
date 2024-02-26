package service

import (
	"context"

	"github.com/jirawan-chuapradit/cards_assignment/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentService interface {
	Create()
	Update()
	Delete(ctx context.Context, commentID primitive.ObjectID) error
}

type commentService struct {
	cardsRepository repository.CardsRepository
}

func NewCommentService(cardsRepository repository.CardsRepository) CommentService {
	return &commentService{
		cardsRepository: cardsRepository,
	}
}

func (s *commentService) Create() {}

func (s *commentService) Update() {}

func (s *commentService) Delete(ctx context.Context, commentId primitive.ObjectID) error {
	return s.cardsRepository.DeleteComment(ctx, commentId)
}
