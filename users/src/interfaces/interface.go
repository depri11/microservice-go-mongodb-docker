package interfaces

import (
	"github.com/depri11/microservice-go-mongodb-docker/users/src/helper"
	"github.com/depri11/microservice-go-mongodb-docker/users/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	FindAll() ([]model.User, error)
	Insert(user model.User) (*mongo.InsertOneResult, error)
}

type UserService interface {
	FindAll() (*helper.Res, error)
	Create(user model.User) (*helper.Res, error)
}
