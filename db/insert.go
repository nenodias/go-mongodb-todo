package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Insert(ctx mongo.SessionContext, collection string, data any) (primitive.ObjectID, error) {
	c := ctx.Client().Database(DBNAME).Collection(collection)
	result, err := c.InsertOne(ctx, data)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}
