package handler

import (
	"encoding/json"
	"net/http"

	"github.com/depri11/microservice-go-mongodb-docker/bookings/src/interfaces"
	"github.com/depri11/microservice-go-mongodb-docker/bookings/src/model"
	"github.com/gorilla/mux"
)

type handler struct {
	service interfaces.BookingService
}

func NewHandler(service interfaces.BookingService) *handler {
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

func (h *handler) FindById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	data, err := h.service.FindById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data.Send(w)
}

func (h *handler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var booking model.Booking
	json.NewDecoder(r.Body).Decode(&booking)

	res, err := h.service.Create(booking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Send(w)
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	var booking model.Booking
	json.NewDecoder(r.Body).Decode(&booking)

	res, err := h.service.Update(id, booking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Send(w)
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	res, err := h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Send(w)
}
