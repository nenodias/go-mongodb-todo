package users

import (
	"github.com/gofiber/fiber/v3"
	"github.com/nenodias/go-mongodb-todo/db"
)

func updateUserById(c fiber.Ctx) error {
	body := new(User)
	result := User{}
	if err := c.Bind().Body(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err := db.UpdateByID("users", c.Params("id"), body, &result)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
}
