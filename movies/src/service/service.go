package service

import (
	"errors"
	"time"

	"github.com/depri11/microservice-go-mongodb-docker/movies/src/helper"
	interfaces "github.com/depri11/microservice-go-mongodb-docker/movies/src/interfaces"
	model "github.com/depri11/microservice-go-mongodb-docker/movies/src/model"
)

type service struct {
	repository interfaces.MovieRepository
}

func NewService(repository interfaces.MovieRepository) *service {
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

func (s *service) Create(movie model.Movie) (*helper.Res, error) {
	movie.CreatedAt = time.Now()
	data, err := s.repository.Insert(movie)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil

}

func (s *service) Update(id string, movie model.Movie) (*helper.Res, error) {
	data, err := s.repository.Update(id, movie)
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
