package main

import (
	"log"
	"net/http"

	"github.com/depri11/microservice-go-mongodb-docker/movies/src/database"
	"github.com/depri11/microservice-go-mongodb-docker/movies/src/route"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	db, err := database.SetupDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	route.SetupRoute(r, db)

	if err := http.ListenAndServe("0.0.0.0:8080", r); err != nil {
		log.Fatal(err)
	}

}
