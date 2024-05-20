package tags

import (
	"github.com/gofiber/fiber/v3"
	"github.com/nenodias/go-mongodb-todo/db"
)

func getAll(c fiber.Ctx) error {
	tasks := []Tag{}
	err := db.Find(COLLECTION, nil, &tasks)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(tasks)
}

func getById(c fiber.Ctx) error {
	user := new(Tag)
	err := db.FindByID(COLLECTION, c.Params("id"), user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}
