package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupDB() (*mongo.Database, error) {
	clientOption := options.Client().ApplyURI("mongodb+srv://dev:dev@cluster0.4uas9.mongodb.net/?retryWrites=true&w=majority")

	db, err := mongo.NewClient(clientOption)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = db.Connect(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB")

	return db.Database("microservice_users"), nil
}
