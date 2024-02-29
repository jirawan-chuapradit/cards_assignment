package repository

import (
	"context"

	"github.com/jirawan-chuapradit/cards_assignment/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepository interface {
	Create(ctx context.Context, user models.User) error
}

type usersRepository struct {
	CardsAssignmentDatabase mongo.Database
}

func NewUsersRepository(conn *mongo.Client) UsersRepository {
	return &usersRepository{
		CardsAssignmentDatabase: *conn.Database("cards_assignment"),
	}
}

// create user
func (r *usersRepository) Create(ctx context.Context, user models.User) error {
	_, err := r.CardsAssignmentDatabase.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
