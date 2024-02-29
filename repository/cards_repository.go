package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/jirawan-chuapradit/cards_assignment/config"
	"github.com/jirawan-chuapradit/cards_assignment/models"
	"github.com/jirawan-chuapradit/cards_assignment/models/request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CardsRepository interface {
	FindById(ctx context.Context, cardsId primitive.ObjectID) (cardsDetails models.CardsDetail, err error)
	FindAll(ctx context.Context) ([]models.CardsDetail, error)
	Create(ctx context.Context, card models.CardsDetail) (models.CardsDetail, error)
	Update(ctx context.Context, cardReq request.UpdateCardRequestBody) error
	Store(ctx context.Context, cardsId primitive.ObjectID) error

	DeleteComment(ctx context.Context, commentId primitive.ObjectID, deleteBy string) error
	UpdateComment(ctx context.Context, commentReq request.UpdateCommentBody) error
	CreateComment(ctx context.Context, cardId primitive.ObjectID, comment models.Comment) error
}

type cardsRepository struct {
	CardsAssignmentDatabase mongo.Database
}

func NewCardsRepository(conn *mongo.Client) CardsRepository {
	return &cardsRepository{
		CardsAssignmentDatabase: *conn.Database("cards_assignment"),
	}
}

// find card by cards Id
func (r *cardsRepository) FindById(ctx context.Context, cardsId primitive.ObjectID) (cardsDetails models.CardsDetail, err error) {
	filter := map[string]interface{}{
		"_id": cardsId,
		"is_archive": bson.M{
			"$ne": true,
		},
	}
	result := r.CardsAssignmentDatabase.Collection("cards").FindOne(ctx, filter)
	if err = result.Decode(&cardsDetails); err != nil {
		return
	}
	log.Println(cardsDetails)
	return
}

// find cards
func (r *cardsRepository) FindAll(ctx context.Context) ([]models.CardsDetail, error) {
	var cards []models.CardsDetail
	filter := map[string]interface{}{
		"is_archive": bson.M{
			"$ne": true,
		},
	}
	cursor, err := r.CardsAssignmentDatabase.Collection("cards").Find(ctx, filter)
	if err != nil {
		return cards, err
	}

	if err := cursor.All(ctx, &cards); err != nil {
		return cards, err
	}
	return cards, nil
}

// create card
func (r *cardsRepository) Create(ctx context.Context, card models.CardsDetail) (models.CardsDetail, error) {
	result, err := r.CardsAssignmentDatabase.Collection("cards").InsertOne(ctx, card)
	if err != nil {
		return card, err
	}

	card.ID = result.InsertedID.(*primitive.ObjectID)
	return card, nil
}

// update card
func (r *cardsRepository) Update(ctx context.Context, cardReq request.UpdateCardRequestBody) error {
	filter := map[string]interface{}{
		"_id": cardReq.ID,
	}
	update := map[string]interface{}{
		"$set": bson.M{
			"title":       cardReq.Title,
			"description": cardReq.Description,
			"status":      cardReq.Status,
			"updated_at":  cardReq.UpdatedAt,
		},
	}

	return r.CardsAssignmentDatabase.Collection("cards").FindOneAndUpdate(ctx, filter, update).Err()
}

// store card in an archive
func (r *cardsRepository) Store(ctx context.Context, cardsId primitive.ObjectID) error {
	filter := map[string]interface{}{
		"_id": cardsId,
	}
	update := map[string]interface{}{
		"$set": map[string]interface{}{
			"is_archive": true,
			"updated_at": time.Now().In(config.Location),
		},
	}

	return r.CardsAssignmentDatabase.Collection("cards").FindOneAndUpdate(ctx, filter, update).Err()
}

// delete comment
func (r *cardsRepository) DeleteComment(ctx context.Context, commentId primitive.ObjectID, deleteBy string) error {
	filter := map[string]interface{}{
		"comments._id":        commentId,
		"comments.created_by": deleteBy,
	}

	return r.CardsAssignmentDatabase.Collection("cards").FindOneAndDelete(ctx, filter).Err()
}

// update comment
func (r *cardsRepository) UpdateComment(ctx context.Context, commentReq request.UpdateCommentBody) error {
	filter := map[string]interface{}{
		"comments._id":        commentReq.ID,
		"comments.created_by": commentReq.UpdatedBy,
	}

	update := map[string]interface{}{
		"$set": bson.M{
			"comments.$.description": commentReq.Description,
			"comments.$.updated_at":  commentReq.UpdatedAt,
		},
	}
	result, err := r.CardsAssignmentDatabase.Collection("cards").UpdateOne(ctx, filter, update)
	if result.MatchedCount == 0 {
		return errors.New("cannot modify comment")
	}
	return err
}

// create comment
func (r *cardsRepository) CreateComment(ctx context.Context, cardId primitive.ObjectID, comment models.Comment) error {
	filter := bson.M{
		"_id": cardId,
	}
	commentId := primitive.NewObjectID()
	comment.ID = &commentId
	update := bson.M{
		"$push": bson.M{
			"comments": comment,
		},
	}
	_, err := r.CardsAssignmentDatabase.Collection("cards").UpdateOne(ctx, filter, update)

	return err
}
