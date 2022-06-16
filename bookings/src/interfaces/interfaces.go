package interfaces

import (
	"github.com/depri11/microservice-go-mongodb-docker/bookings/src/helper"
	"github.com/depri11/microservice-go-mongodb-docker/bookings/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookingRepository interface {
	FindAll() ([]model.Booking, error)
	FindByID(id string) (*model.Booking, error)
	Insert(booking model.Booking) (*mongo.InsertOneResult, error)
	Update(id string, booking model.Booking) (*mongo.UpdateResult, error)
	Delete(id string) (*mongo.DeleteResult, error)
}

type BookingService interface {
	FindAll() (*helper.Res, error)
	Create(booking model.Booking) (*helper.Res, error)
	FindById(id string) (*helper.Res, error)
	Update(id string, booking model.Booking) (*helper.Res, error)
	Delete(id string) (*helper.Res, error)
}
