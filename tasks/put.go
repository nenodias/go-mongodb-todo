package tasks

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/nenodias/go-mongodb-todo/db"
	"github.com/nenodias/go-mongodb-todo/tags"
	"go.mongodb.org/mongo-driver/mongo"
)

func updateById(c fiber.Ctx) error {
	body := new(Task)
	if err := c.Bind().Body(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var prevTask Task
	result := Task{}
	err := db.DoConnection(context.Background(), func(ctx mongo.SessionContext) error {
		err := db.FindByID(ctx, COLLECTION, c.Params("id"), &prevTask)
		if err != nil {
			return err
		}

		err = db.UpdateByID(ctx, COLLECTION, c.Params("id"), body, &result)
		if err != nil {
			return err
		}

		return updateTagsTask(ctx, c.Params("id"), prevTask.Tags, result.Tags)
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(result)
}

func updateTagsTask(ctx mongo.SessionContext, id string, oldTags, newTags []string) error {
	mapOldTags := make(map[string]int, len(oldTags))
	for k, v := range oldTags {
		mapOldTags[string(v)] = k
	}
	var diffTags []string
	for _, v := range newTags {
		if _, ok := mapOldTags[string(v)]; !ok {
			diffTags = append(diffTags, string(v))
		} else {
			delete(mapOldTags, string(v))
		}
	}
	if len(diffTags) > 0 {
		err := tags.AddTask(ctx, id, diffTags)
		if err != nil {
			return err
		}
	}
	if len(mapOldTags) > 0 {
		deleteTags := make([]string, 0, len(mapOldTags))
		for k := range mapOldTags {
			deleteTags = append(deleteTags, k)
		}
		err := tags.RemoveTask(ctx, id, deleteTags...)
		if err != nil {
			return err
		}

	}
	return nil
}
