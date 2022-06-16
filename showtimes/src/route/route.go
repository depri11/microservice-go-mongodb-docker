package route

import (
	"github.com/depri11/microservice-go-mongodb-docker/showtimes/src/handler"
	"github.com/depri11/microservice-go-mongodb-docker/showtimes/src/repository"
	"github.com/depri11/microservice-go-mongodb-docker/showtimes/src/service"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoute(r *mux.Router, db *mongo.Database) {
	route := r.PathPrefix("/showtimes").Subrouter()

	c := db.Collection("showtimes")

	repository := repository.NewRepository(c)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	route.HandleFunc("/", handler.FindAll).Methods("GET")
	route.HandleFunc("/{id}", handler.FindById).Methods("GET")
	route.HandleFunc("/", handler.Register).Methods("POST")
	route.HandleFunc("/{id}", handler.UpdateUser).Methods("PUT")
	route.HandleFunc("/{id}", handler.DeleteUser).Methods("DELETE")

}
