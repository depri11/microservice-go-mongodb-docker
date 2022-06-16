package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Showtime struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Date      string             `json:"date,omitempty" bson:"date"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"CreatedAt"`
	Movies    []string           `json:"movies,omitempty" bson:"movies"`
}
