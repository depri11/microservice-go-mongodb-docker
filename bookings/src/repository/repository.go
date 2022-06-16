package repository

import (
	"context"
	"errors"

	"github.com/depri11/microservice-go-mongodb-docker/bookings/src/model"
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

func (r *repository) FindAll() ([]model.Booking, error) {
	ctx := context.TODO()

	userCursor, err := r.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var bookings []model.Booking

	err = userCursor.All(ctx, &bookings)
	if err != nil {
		return nil, err
	}

	return bookings, nil
}

func (r *repository) Insert(user model.Booking) (*mongo.InsertOneResult, error) {
	ctx := context.TODO()
	return r.C.InsertOne(ctx, user)
}

func (r *repository) FindByID(id string) (*model.Booking, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()

	var booking model.Booking
	err = r.C.FindOne(ctx, bson.M{"_id": p}).Decode(&booking)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocument")
		}

		return nil, err
	}

	return &booking, nil
}

func (r *repository) Update(id string, booking model.Booking) (*mongo.UpdateResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	return r.C.UpdateOne(ctx, bson.M{"_id": p}, bson.M{"$set": booking})
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
