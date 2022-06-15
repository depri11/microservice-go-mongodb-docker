package handler

import (
	"encoding/json"
	"net/http"

	"github.com/depri11/microservice-go-mongodb-docker/users/src/interfaces"
	"github.com/depri11/microservice-go-mongodb-docker/users/src/model"
)

type handler struct {
	service interfaces.UserService
}

func NewHandler(service interfaces.UserService) *handler {
	return &handler{service}
}

func (h *handler) FindAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := h.service.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data.Send(w)
}

func (h *handler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	res, err := h.service.Create(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Send(w)
}
