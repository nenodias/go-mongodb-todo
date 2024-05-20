package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Find(collection string, filter bson.M, documents any) error {
	client, ctx := GetConnection()
	defer func() {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Println(err)
		}
	}()

	c := client.Database(DBNAME).Collection(collection)
	if filter == nil {
		filter = bson.M{}
	}
	cursor, err := c.Find(ctx, filter)
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
	defer func() {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Println(err)
		}
	}()

	c := client.Database(DBNAME).Collection(collection)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	return c.FindOne(context.Background(), filter).Decode(document)
}

func FindOne(collection string, filter bson.M, document any) error {
	client, ctx := GetConnection()
	defer func() {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Println(err)
		}
	}()

	c := client.Database(DBNAME).Collection(collection)
	return c.FindOne(context.Background(), filter).Decode(document)
}
