package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Find(collection string, documents any) error {
	client, ctx := GetConnection()
	defer client.Disconnect(ctx)

	c := client.Database(DBNAME).Collection(collection)
	cursor, err := c.Find(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		err := cursor.Close(context.Background())
		if err != nil {
			log.Println(err)
		}
	}()
	return cursor.All(ctx, documents)
}

func FindByID(collection string, id string, document any) error {
	client, ctx := GetConnection()
	defer client.Disconnect(ctx)

	c := client.Database(DBNAME).Collection(collection)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	return c.FindOne(context.Background(), filter).Decode(document)
}
