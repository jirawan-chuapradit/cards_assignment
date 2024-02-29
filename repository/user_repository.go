package repository

import (
	"context"

	"github.com/jirawan-chuapradit/cards_assignment/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepository interface {
	Create(ctx context.Context, user models.User) error
	FindByEmail(ctx context.Context, email string) (models.User, error)
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
// TODO: set email is unique in db
func (r *usersRepository) Create(ctx context.Context, user models.User) error {
	_, err := r.CardsAssignmentDatabase.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

// get user
func (r *usersRepository) FindByEmail(ctx context.Context, email string) (models.User, error) {
	user := models.User{}
	result := r.CardsAssignmentDatabase.Collection("users").FindOne(ctx, bson.M{"email": email})
	if err := result.Decode(&user); err != nil {
		return user, err
	}
	return user, nil
}
