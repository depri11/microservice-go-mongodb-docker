package service

import (
	"errors"
	"time"

	"github.com/depri11/microservice-go-mongodb-docker/bookings/src/helper"
	interfaces "github.com/depri11/microservice-go-mongodb-docker/bookings/src/interfaces"
	model "github.com/depri11/microservice-go-mongodb-docker/bookings/src/model"
)

type service struct {
	repository interfaces.BookingRepository
}

func NewService(repository interfaces.BookingRepository) *service {
	return &service{repository}
}

func (s *service) FindAll() (*helper.Res, error) {
	data, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}

func (s *service) FindById(id string) (*helper.Res, error) {
	data, err := s.repository.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDucoment" {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}

func (s *service) Create(booking model.Booking) (*helper.Res, error) {
	booking.CreatedAt = time.Now()
	data, err := s.repository.Insert(booking)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil

}

func (s *service) Update(id string, booking model.Booking) (*helper.Res, error) {
	data, err := s.repository.Update(id, booking)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}

func (s *service) Delete(id string) (*helper.Res, error) {
	data, err := s.repository.Delete(id)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}
