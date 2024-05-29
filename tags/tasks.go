package tags

import (
	"sort"

	"github.com/nenodias/go-mongodb-todo/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func findOrCreate(ctx mongo.SessionContext, name string) (tag Tag, err error) {
	filter := bson.M{"name": name}
	err = db.FindOne(ctx, COLLECTION, filter, &tag)
	if err != nil && err != mongo.ErrNoDocuments {
		return
	}
	if tag.Name != "" {
		return
	}

	tag.Name = name
	id, err := db.Insert(ctx, COLLECTION, tag)
	if err != nil {
		return
	}
	tag.ID = id

	return tag, nil
}

func AddTask(ctx mongo.SessionContext, taskID string, names []string) error {
	for _, name := range names {
		tag, err := findOrCreate(ctx, name)
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
		err = db.UpdateByID(ctx, COLLECTION, tag.ID.Hex(), tag, result)
		if err != nil {
			return err
		}
	}
	return nil
}

func RemoveTask(ctx mongo.SessionContext, taskID string, names ...string) error {
	filter := bson.M{"tasks": taskID}
	if len(names) > 0 {
		filter["name"] = bson.M{"$in": names}
	}
	var documents []Tag
	err := db.Find(ctx, COLLECTION, filter, &documents)
	if err != nil {
		return err
	}
	var result Tag
	for _, document := range documents {
		i := sort.SearchStrings(document.Tasks, taskID)
		document.Tasks = append(document.Tasks[:i], document.Tasks[i+1:]...)
		err = db.UpdateByID(ctx, COLLECTION, document.ID.Hex(), document, &result)
		if err != nil {
			return err
		}
	}
	return nil
}
