package service

import (
	"errors"
	"time"

	"github.com/depri11/microservice-go-mongodb-docker/showtimes/src/helper"
	"github.com/depri11/microservice-go-mongodb-docker/showtimes/src/interfaces"
	model "github.com/depri11/microservice-go-mongodb-docker/showtimes/src/model"
)

type service struct {
	repository interfaces.ShowtimeRepository
}

func NewService(repository interfaces.ShowtimeRepository) *service {
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

func (s *service) Create(showtime model.Showtime) (*helper.Res, error) {
	// var movieid = []string{"62aad5601136b55656bd6692", "62aad955bc7b8c454950d6c7"}

	showtime.CreatedAt = time.Now()
	// showtime.Movies = movieid

	data, err := s.repository.Insert(showtime)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil

}

func (s *service) Update(id string, showtime model.Showtime) (*helper.Res, error) {
	data, err := s.repository.Update(id, showtime)
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
