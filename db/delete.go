package db

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DeleteById(ctx mongo.SessionContext, collection, id string) error {
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	c := ctx.Client().Database(DBNAME).Collection(collection)
	result, err := c.DeleteOne(ctx, bson.M{"_id": primitiveID}, options.Delete())
	if err != nil {
		return err
	}
	if result.DeletedCount != 1 {
		return fmt.Errorf("%d documents deleted with the given ID: %s", result.DeletedCount, id)
	}
	return nil
}
