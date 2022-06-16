package interfaces

import (
	"github.com/depri11/microservice-go-mongodb-docker/showtimes/src/helper"
	"github.com/depri11/microservice-go-mongodb-docker/showtimes/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShowtimeRepository interface {
	FindAll() ([]model.Showtime, error)
	FindByID(id string) (*model.Showtime, error)
	Insert(movie model.Showtime) (*mongo.InsertOneResult, error)
	Update(id string, movie model.Showtime) (*mongo.UpdateResult, error)
	Delete(id string) (*mongo.DeleteResult, error)
}

type ShowtimeService interface {
	FindAll() (*helper.Res, error)
	Create(movie model.Showtime) (*helper.Res, error)
	FindById(id string) (*helper.Res, error)
	Update(id string, movie model.Showtime) (*helper.Res, error)
	Delete(id string) (*helper.Res, error)
}
