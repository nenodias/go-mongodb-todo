package tasks

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	COLLECTION = "tasks"
)

type Task struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name,omitempty"`
}
