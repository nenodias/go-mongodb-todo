package tags

import (
	"sort"

	"github.com/nenodias/go-mongodb-todo/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func findOrCreate(name string) (tag Tag, err error) {
	filter := bson.M{"name": name}
	err = db.FindOne(COLLECTION, filter, &tag)
	if err != nil && err != mongo.ErrNoDocuments {
		return
	}
	if tag.Name != "" {
		return
	}

	tag.Name = name
	id, err := db.Insert(COLLECTION, tag)
	if err != nil {
		return
	}
	tag.ID = id

	return tag, nil
}

func addTask(taskID string, names []string) error {
	for _, name := range names {
		tag, err := findOrCreate(name)
		if err != nil {
			return err
		}
		i := sort.SearchStrings(tag.Tasks, taskID)
		if i < len(tag.Tasks) && tag.Tasks[i] == taskID {
			continue
		}
		tag.Tasks = append(tag.Tasks, taskID)
		sort.Strings(tag.Tasks)
		result := new(Tag)
		err = db.UpdateByID(COLLECTION, tag.ID.Hex(), tag, result)
		if err != nil {
			return err
		}
	}
	return nil
}
