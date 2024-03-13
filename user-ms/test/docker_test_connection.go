package test

import (
	"context"
	"fmt"
	"log"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func OpenConnection() (database *mongo.Database, close func()) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Docker test error to build connection pool: %s", err)
		return
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "7-jammy",
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	_ = resource.Expire(360)

	if err != nil {
		log.Fatalf("Erro do create mongo container: %s", err)
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(
		fmt.Sprintf("mongodb://localhost:%s", resource.GetPort("27017/tcp"))))

	if err != nil {
		log.Fatalf("Error trying to open connection: %s", err)
	}

	return client.Database("user_test"), func() {
		err := resource.Close()
		if err != nil {
			log.Println("Error trying to close connection")
		}
		err = pool.Purge(resource)
		if err != nil {
			log.Println("Cannot purge the resources")
		}
	}

}
