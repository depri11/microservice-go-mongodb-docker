package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title"`
	Director  string             `json:"director,omitempty" bson:"director"`
	Rating    float32            `json:"rating,omitempty" bson:"rating"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"CreatedAt"`
}
