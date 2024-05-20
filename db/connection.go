package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DBNAME               = "todos"
	DEFAULT_DATABASE_URI = "mongodb://root:example@localhost:27017/?authSource=admin"
)

func GetConnection() (client *mongo.Client, ctx context.Context) {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = DEFAULT_DATABASE_URI
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	return
}
