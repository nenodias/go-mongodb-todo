package tasks

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/nenodias/go-mongodb-todo/db"
	"github.com/nenodias/go-mongodb-todo/tags"
)

func addItem(c fiber.Ctx) error {
	body := new(Task)
	if err := c.Bind().Body(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	id, err := db.Insert(COLLECTION, body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	body.ID = id

	err = tags.AddTask(body.ID.Hex(), body.Tags)
	if err != nil {
		deleteErr := db.DeleteById(COLLECTION, body.ID.Hex())
		if deleteErr != nil {
			log.Println(err)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(body)
}
