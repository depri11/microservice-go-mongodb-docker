package repository

import (
	"context"

	"github.com/depri11/microservice-go-mongodb-docker/users/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	C *mongo.Collection
}

func NewRepository(c *mongo.Collection) *repository {
	return &repository{c}
}

func (r *repository) FindAll() ([]model.User, error) {
	ctx := context.TODO()

	userCursor, err := r.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var users []model.User

	err = userCursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *repository) Insert(user model.User) (*mongo.InsertOneResult, error) {
	ctx := context.TODO()
	return r.C.InsertOne(ctx, user)
}
