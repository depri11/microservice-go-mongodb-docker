package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	FirstName string             `json:"firstname,omitempty" bson:"firstname"`
	LastName  string             `json:"lastname,omitempty" bson:"lastname"`
}
