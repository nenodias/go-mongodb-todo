package tasks

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	COLLECTION = "tasks"
)

type Task struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"name,omitempty"`
	Description string             `json:"description" bson:"name,omitempty"`
	Tags        []string           `json:"tags" bson:"tags,omitempty"`
	Assing      string             `json:"assing" bson:"assing,omitempty"`
}
