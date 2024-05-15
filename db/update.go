package db

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateByID(collection string, id string, data, result any) error {
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
	opts := options.Update().SetUpsert(false)
	update, err := c.UpdateByID(context.Background(), primitiveID, bson.M{
		"$set": data,
	}, opts)
	if err != nil {
		return err
	}
	if update.MatchedCount == 0 {
		return errors.New("no document found with the given ID")
	}
	err = FindByID(collection, id, result)
	if err != nil {
		return err
	}
	return nil
}
