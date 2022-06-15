package service

import (
	"github.com/depri11/microservice-go-mongodb-docker/users/src/helper"
	interfaces "github.com/depri11/microservice-go-mongodb-docker/users/src/interfaces"
	"github.com/depri11/microservice-go-mongodb-docker/users/src/model"
)

type service struct {
	repository interfaces.UserRepository
}

func NewService(repository interfaces.UserRepository) *service {
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

func (s *service) Create(user model.User) (*helper.Res, error) {
	data, err := s.repository.Insert(user)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil

}
