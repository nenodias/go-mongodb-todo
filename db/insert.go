package db

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Insert(collection string, data any) (primitive.ObjectID, error) {
	// Get the MongoDB connection
	client, ctx := GetConnection()
	defer func() {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Println(err)
		}
	}()

	c := client.Database(DBNAME).Collection(collection)
	result, err := c.InsertOne(ctx, data)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}
