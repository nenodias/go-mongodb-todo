package users

import (
	"github.com/gofiber/fiber/v3"
	"github.com/nenodias/go-mongodb-todo/db"
)

func getAll(c fiber.Ctx) error {
	users := []User{}
	err := db.Find(COLLECTION, nil, &users)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func getById(c fiber.Ctx) error {
	user := new(User)
	err := db.FindByID(COLLECTION, c.Params("id"), user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}
