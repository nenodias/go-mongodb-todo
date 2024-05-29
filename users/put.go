package users

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/nenodias/go-mongodb-todo/db"
	"go.mongodb.org/mongo-driver/mongo"
)

func updateById(c fiber.Ctx) error {
	body := new(User)
	result := User{}
	if err := c.Bind().Body(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err := db.DoConnection(context.Background(), func(ctx mongo.SessionContext) error {
		return db.UpdateByID(ctx, COLLECTION, c.Params("id"), body, &result)
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
}
