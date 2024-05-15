package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DBNAME = "todos"
)

func GetConnection() (client *mongo.Client, ctx context.Context) {
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017/?authSource=admin"))
	if err != nil {
		log.Fatal(err)
	}
	return
}
