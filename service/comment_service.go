package service

import (
	"context"
	"time"

	"github.com/jirawan-chuapradit/cards_assignment/config"
	"github.com/jirawan-chuapradit/cards_assignment/models/request"
	"github.com/jirawan-chuapradit/cards_assignment/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentService interface {
	Create()
	Update(ctx context.Context, commentReq request.UpdateCommentBody) error
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

// update comment
func (s *commentService) Update(ctx context.Context, commentReq request.UpdateCommentBody) error {
	now := time.Now().In(config.Location)
	commentReq.UpdatedAt = &now
	return s.cardsRepository.UpdateComment(ctx, commentReq)
}

// delete comment
func (s *commentService) Delete(ctx context.Context, commentId primitive.ObjectID) error {
	return s.cardsRepository.DeleteComment(ctx, commentId)
}
