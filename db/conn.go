package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Setup() (*mongo.Client, error) {
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")

	cardsAssignmentDB := client.Database("cards_assignment")
	if err := cardsAssignmentDB.CreateCollection(ctx, "users"); err != nil {
		return nil, err
	}
	if err := cardsAssignmentDB.CreateCollection(ctx, "cards"); err != nil {
		return nil, err
	}

	if err := cardsAssignmentDB.CreateCollection(ctx, "cards_history"); err != nil {
		return nil, err
	}

	// create index
	indexs, err := cardsAssignmentDB.Collection("users").Indexes().ListSpecifications(ctx)
	if err != nil {
		return nil, err
	}
	if err := createIndex(indexs, "ix_email", bson.D{{Key: "email", Value: 1}}, true, cardsAssignmentDB.Collection("users")); err != nil {
		return nil, err
	}
	if err := createIndex(indexs, "ix_username", bson.D{{Key: "username", Value: 1}}, true, cardsAssignmentDB.Collection("users")); err != nil {
		return nil, err
	}

	if err := createIndex(indexs, "ix_card_id", bson.D{{Key: "card_id", Value: 1}}, false, cardsAssignmentDB.Collection("cards_history")); err != nil {
		return nil, err
	}

	if err := createIndex(indexs, "ix_id_is_archive", bson.D{{Key: "_id", Value: 1}, {Key: "is_archive", Value: 1}}, false, cardsAssignmentDB.Collection("cards")); err != nil {
		return nil, err
	}

	if err := createIndex(indexs, "ix_is_archive", bson.D{{Key: "is_archive", Value: 1}}, false, cardsAssignmentDB.Collection("cards")); err != nil {
		return nil, err
	}

	if err := createIndex(indexs, "ix_id_created_by", bson.D{{Key: "_id", Value: 1}, {Key: "created_by", Value: 1}}, false, cardsAssignmentDB.Collection("cards")); err != nil {
		return nil, err
	}

	return client, nil
}

func createIndex(indexs []*mongo.IndexSpecification, name string, keys primitive.D, isUnique bool, collection *mongo.Collection) error {
	ctx := context.Background()
	for _, i := range indexs {
		if i.Name == name {
			return nil
		}
	}

	idxOpt := options.Index()
	idxOpt.SetName(name)
	idxOpt.SetUnique(isUnique)
	idx := mongo.IndexModel{
		Keys:    keys,
		Options: idxOpt,
	}
	if _, err := collection.Indexes().CreateOne(ctx, idx); err != nil {
		return err
	}

	return nil
}
