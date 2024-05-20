package tags

import (
	"github.com/gofiber/fiber/v3"
	"github.com/nenodias/go-mongodb-todo/db"
)

func updateById(c fiber.Ctx) error {
	body := new(Tag)
	result := Tag{}
	if err := c.Bind().Body(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err := db.UpdateByID(COLLECTION, c.Params("id"), body, &result)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
}
