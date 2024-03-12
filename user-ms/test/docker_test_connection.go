package test

import (
	"log"

	"github.com/ory/dockertest/v3"
	"go.mongodb.org/mongo-driver/mongo"
)

func OpenConnection() (database *mongo.Database) {
	_, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("error")
	}
}
