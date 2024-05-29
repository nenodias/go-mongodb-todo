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
	DEFAULT_DATABASE_URI = "mongodb://127.0.0.1:27017/?replicaSet=rs0"
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

func DoConnection(ctx context.Context, fn func(ctx mongo.SessionContext) error) error {
	client, ctx := GetConnection()
	defer func() {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Println(err)
		}
	}()
	return client.UseSession(ctx, func(sctx mongo.SessionContext) error {
		err := sctx.StartTransaction()
		if err != nil {
			return err
		}
		err = fn(sctx)
		if err != nil {
			sctx.AbortTransaction(sctx)
			return err
		}
		return sctx.CommitTransaction(sctx)
	})
}
