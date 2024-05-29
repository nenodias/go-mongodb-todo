package db

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateByID(ctx mongo.SessionContext, collection string, id string, data, result any) error {
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	c := ctx.Client().Database(DBNAME).Collection(collection)
	opts := options.Update().SetUpsert(false)
	update, err := c.UpdateByID(ctx, primitiveID, bson.M{
		"$set": data,
	}, opts)
	if err != nil {
		return err
	}
	if update.ModifiedCount != 1 {
		return fmt.Errorf("%d documents updated with the given ID: %s", update.ModifiedCount, id)
	}
	err = FindByID(ctx, collection, id, result)
	if err != nil {
		return err
	}
	return nil
}
