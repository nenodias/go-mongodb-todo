package db

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DeleteById(collection, id string) error {
	client, ctx := GetConnection()
	defer func() {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Println(err)
		}
	}()
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	c := client.Database(DBNAME).Collection(collection)
	result, err := c.DeleteOne(ctx, bson.M{"_id": primitiveID}, options.Delete())
	if err != nil {
		return err
	}
	if result.DeletedCount != 1 {
		return fmt.Errorf("%d documents deleted with the given ID: %s", result.DeletedCount, id)
	}
	return nil
}
