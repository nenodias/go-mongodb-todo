package tasks

import (
	"github.com/gofiber/fiber/v3"
	"github.com/nenodias/go-mongodb-todo/db"
	"github.com/nenodias/go-mongodb-todo/tags"
)

func updateById(c fiber.Ctx) error {
	body := new(Task)
	if err := c.Bind().Body(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var prevTask Task
	err := db.FindByID(COLLECTION, c.Params("id"), &prevTask)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	result := Task{}
	err = db.UpdateByID(COLLECTION, c.Params("id"), body, &result)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	err = updateTagsTask(c.Params("id"), prevTask.Tags, result.Tags)
	if err != nil {
		err := tags.RemoveTask(c.Params("id"), result.Tags...)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		err = db.DeleteById(COLLECTION, c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		_, err = db.Insert(COLLECTION, &result)
		if err != nil {
			db.DeleteById(COLLECTION, c.Params("id"))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		err = tags.AddTask(result.ID.Hex(), result.Tags)
		if err != nil {
			db.DeleteById(COLLECTION, c.Params("id"))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	return c.JSON(result)
}

func updateTagsTask(id string, oldTags, newTags []string) error {
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
		err := tags.AddTask(id, diffTags)
		if err != nil {
			return err
		}
	}
	if len(mapOldTags) > 0 {
		deleteTags := make([]string, 0, len(mapOldTags))
		for k := range mapOldTags {
			deleteTags = append(deleteTags, k)
		}
		err := tags.RemoveTask(id, deleteTags...)
		if err != nil {
			return err
		}

	}
	return nil
}
