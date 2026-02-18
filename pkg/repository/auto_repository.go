package repository

import (
	"context"
	"noleggio_auto/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type AutoRepository interface {
	Create(ctx context.Context, auto *model.Auto) error
	FindByID(ctx context.Context, id string) (*model.Auto, error)
	FindAll(ctx context.Context) ([]model.Auto, error)
	Update(ctx context.Context, id string, auto *model.Auto) error
	Delete(ctx context.Context, id string) error
}

type autoRepository struct {
	collection *mongo.Collection
}

func NewAutoRepository(db *mongo.Database) AutoRepository {
	return &autoRepository{
		collection: db.Collection("auto"),
	}
}

// CREATE
func (r *autoRepository) Create(ctx context.Context, auto *model.Auto) error {
	_, err := r.collection.InsertOne(ctx, auto)
	return err
}

// FIND ALL da implementare
func (r *autoRepository) FindAll(ctx context.Context) ([]model.Auto, error) {
    panic("non ancora implementato")
}

//FIND BY ID
func (r *autoRepository) FindByID(ctx context.Context, id string) (*model.Auto, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var auto model.Auto
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&auto)
	return &auto, err
}

//UPDATE
func (r *autoRepository) Update(ctx context.Context, id string, auto *model.Auto) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": auto}
	
	_, err = r.collection.UpdateOne(ctx, filter, update)
	return err
}

//DELETE
func (r *autoRepository) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}