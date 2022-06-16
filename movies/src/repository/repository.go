package repository

import (
	"context"
	"errors"

	"github.com/depri11/microservice-go-mongodb-docker/movies/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	C *mongo.Collection
}

func NewRepository(c *mongo.Collection) *repository {
	return &repository{c}
}

func (r *repository) FindAll() ([]model.Movie, error) {
	ctx := context.TODO()

	userCursor, err := r.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var movies []model.Movie

	err = userCursor.All(ctx, &movies)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *repository) Insert(user model.Movie) (*mongo.InsertOneResult, error) {
	ctx := context.TODO()
	return r.C.InsertOne(ctx, user)
}

func (r *repository) FindByID(id string) (*model.Movie, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()

	var movie model.Movie
	err = r.C.FindOne(ctx, bson.M{"_id": p}).Decode(&movie)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocument")
		}

		return nil, err
	}

	return &movie, nil
}

func (r *repository) Update(id string, movie model.Movie) (*mongo.UpdateResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	return r.C.UpdateOne(ctx, bson.M{"_id": p}, bson.M{"$set": movie})
}

func (r *repository) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	del, err := r.C.DeleteOne(context.TODO(), bson.M{"_id": p})
	if err != nil {
		return nil, err
	}

	return del, nil
}
