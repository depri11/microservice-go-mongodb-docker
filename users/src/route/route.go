package route

import (
	"github.com/depri11/microservice-go-mongodb-docker/users/src/handler"
	"github.com/depri11/microservice-go-mongodb-docker/users/src/repository"
	"github.com/depri11/microservice-go-mongodb-docker/users/src/service"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoute(r *mux.Router, db *mongo.Database) {
	route := r.PathPrefix("/users").Subrouter()

	c := db.Collection("users")

	repository := repository.NewRepository(c)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	route.HandleFunc("/", handler.FindAll).Methods("GET")
	route.HandleFunc("/", handler.Register).Methods("POST")

}
