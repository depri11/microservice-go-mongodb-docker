package repository

import (
	"context"
	"errors"

	"github.com/depri11/microservice-go-mongodb-docker/users/src/model"
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

func (r *repository) FindByID(id string) (*model.User, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()

	var user model.User
	err = r.C.FindOne(ctx, bson.M{"_id": p}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocument")
		}

		return nil, err
	}

	return &user, nil
}

func (r *repository) Update(id string, user model.User) (*mongo.UpdateResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	return r.C.UpdateOne(ctx, bson.M{"_id": p}, bson.M{"$set": user})
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
