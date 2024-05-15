package db

import "go.mongodb.org/mongo-driver/bson/primitive"

func Insert(collection string, data any) (primitive.ObjectID, error) {
	// Get the MongoDB connection
	client, ctx := GetConnection()
	defer client.Disconnect(ctx)

	c := client.Database(DBNAME).Collection(collection)
	result, err := c.InsertOne(ctx, data)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}
