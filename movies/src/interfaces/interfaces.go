package interfaces

import (
	"github.com/depri11/microservice-go-mongodb-docker/movies/src/helper"
	"github.com/depri11/microservice-go-mongodb-docker/movies/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type MovieRepository interface {
	FindAll() ([]model.Movie, error)
	FindByID(id string) (*model.Movie, error)
	Insert(movie model.Movie) (*mongo.InsertOneResult, error)
	Update(id string, movie model.Movie) (*mongo.UpdateResult, error)
	Delete(id string) (*mongo.DeleteResult, error)
}

type MovieService interface {
	FindAll() (*helper.Res, error)
	Create(movie model.Movie) (*helper.Res, error)
	FindById(id string) (*helper.Res, error)
	Update(id string, movie model.Movie) (*helper.Res, error)
	Delete(id string) (*helper.Res, error)
}
