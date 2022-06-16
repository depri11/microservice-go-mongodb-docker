package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserID     string             `json:"user_id,omitempty" bson:"user_id"`
	ShowTimeID string             `json:"showtime_id,omitempty" bson:"showtime_id"`
	Movies     []string           `json:"movies,omitempty" bson:"movies"`
	CreatedAt  time.Time          `json:"created_at,omitempty" bson:"CreatedAt"`
}
