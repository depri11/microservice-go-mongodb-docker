package interfaces

import (
	"github.com/depri11/microservice-go-mongodb-docker/users/src/helper"
	"github.com/depri11/microservice-go-mongodb-docker/users/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindByID(id string) (*model.User, error)
	Insert(user model.User) (*mongo.InsertOneResult, error)
	Update(id string, user model.User) (*mongo.UpdateResult, error)
	Delete(id string) (*mongo.DeleteResult, error)
}

type UserService interface {
	FindAll() (*helper.Res, error)
	Create(user model.User) (*helper.Res, error)
	FindById(id string) (*helper.Res, error)
	Update(id string, user model.User) (*helper.Res, error)
	Delete(id string) (*helper.Res, error)
}
