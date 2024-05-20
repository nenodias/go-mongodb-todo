package tasks

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/nenodias/go-mongodb-todo/db"
	"github.com/nenodias/go-mongodb-todo/tags"
)

func deleteItem(c fiber.Ctx) error {
	err := tags.RemoveTask(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	err = db.DeleteById(COLLECTION, c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	return c.SendStatus(http.StatusNoContent)
}
