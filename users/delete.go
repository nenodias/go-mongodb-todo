package users

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/nenodias/go-mongodb-todo/db"
	"go.mongodb.org/mongo-driver/mongo"
)

func deleteItem(c fiber.Ctx) error {
	err := db.DoConnection(context.Background(), func(ctx mongo.SessionContext) error {
		err := db.DeleteById(ctx, COLLECTION, c.Params("id"))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	return c.SendStatus(http.StatusNoContent)
}
