package service

import (
	"context"
	"time"

	"github.com/jirawan-chuapradit/cards_assignment/config"
	"github.com/jirawan-chuapradit/cards_assignment/models"
	"github.com/jirawan-chuapradit/cards_assignment/models/request"
	"github.com/jirawan-chuapradit/cards_assignment/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentService interface {
	Create(ctx context.Context, commentReq request.CreateCommentBody) error
	Update(ctx context.Context, commentReq request.UpdateCommentBody) error
	Delete(ctx context.Context, commentID primitive.ObjectID, deleteBy string) error
}

type commentService struct {
	cardsRepository repository.CardsRepository
}

func NewCommentService(cardsRepository repository.CardsRepository) CommentService {
	return &commentService{
		cardsRepository: cardsRepository,
	}
}

// update comment
func (s *commentService) Create(ctx context.Context, commentReq request.CreateCommentBody) error {
	now := time.Now().In(config.Location)
	comment := models.Comment{
		Img:         "", // find img from session
		Description: *commentReq.Description,
		CreatedBy:   "mock usr", // find user name from session
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}
	return s.cardsRepository.CreateComment(ctx, commentReq.CardID, comment)
}

// update comment
func (s *commentService) Update(ctx context.Context, commentReq request.UpdateCommentBody) error {
	return s.cardsRepository.UpdateComment(ctx, commentReq)
}

// delete comment
func (s *commentService) Delete(ctx context.Context, commentId primitive.ObjectID, deleteBy string) error {
	return s.cardsRepository.DeleteComment(ctx, commentId, deleteBy)
}
