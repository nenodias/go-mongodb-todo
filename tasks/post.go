package tasks

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/nenodias/go-mongodb-todo/db"
	"github.com/nenodias/go-mongodb-todo/tags"
	"go.mongodb.org/mongo-driver/mongo"
)

func addItem(c fiber.Ctx) error {
	body := new(Task)
	if err := c.Bind().Body(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err := db.DoConnection(context.Background(), func(ctx mongo.SessionContext) error {
		id, err := db.Insert(ctx, COLLECTION, body)
		if err != nil {
			return err
		}
		body.ID = id
		err = tags.AddTask(ctx, body.ID.Hex(), body.Tags)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(body)
}
