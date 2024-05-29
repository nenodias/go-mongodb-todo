package tags

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/nenodias/go-mongodb-todo/db"
	"go.mongodb.org/mongo-driver/mongo"
)

func getAll(c fiber.Ctx) error {
	tasks := []Tag{}
	err := db.DoConnection(context.Background(), func(ctx mongo.SessionContext) error {
		return db.Find(ctx, COLLECTION, nil, &tasks)
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(tasks)
}

func getById(c fiber.Ctx) error {
	user := new(Tag)
	err := db.DoConnection(context.Background(), func(ctx mongo.SessionContext) error {
		return db.FindByID(ctx, COLLECTION, c.Params("id"), user)
	})
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}
