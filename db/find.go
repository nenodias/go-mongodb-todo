package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Find(ctx mongo.SessionContext, collection string, filter bson.M, documents any) error {
	c := ctx.Client().Database(DBNAME).Collection(collection)
	if filter == nil {
		filter = bson.M{}
	}
	cursor, err := c.Find(ctx, filter)
	if err != nil {
		return err
	}
	defer func() {
		err := cursor.Close(ctx)
		if err != nil {
			log.Println(err)
		}
	}()
	return cursor.All(ctx, documents)
}

func FindByID(ctx mongo.SessionContext, collection string, id string, document any) error {

	c := ctx.Client().Database(DBNAME).Collection(collection)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	return c.FindOne(ctx, filter).Decode(document)
}

func FindOne(ctx mongo.SessionContext, collection string, filter bson.M, document any) error {
	c := ctx.Client().Database(DBNAME).Collection(collection)
	return c.FindOne(context.Background(), filter).Decode(document)
}
