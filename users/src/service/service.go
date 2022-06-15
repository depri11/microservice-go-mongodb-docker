package service

import (
	"errors"

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

func (s *service) Create(user model.User) (*helper.Res, error) {
	data, err := s.repository.Insert(user)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil

}

func (s *service) Update(id string, user model.User) (*helper.Res, error) {
	data, err := s.repository.Update(id, user)
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
