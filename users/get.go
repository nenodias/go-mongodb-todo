package users

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/nenodias/go-mongodb-todo/db"
	"go.mongodb.org/mongo-driver/mongo"
)

func getAll(c fiber.Ctx) error {
	users := []User{}
	err := db.DoConnection(context.Background(), func(ctx mongo.SessionContext) error {
		return db.Find(ctx, COLLECTION, nil, &users)
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func getById(c fiber.Ctx) error {
	user := new(User)
	err := db.DoConnection(context.Background(), func(ctx mongo.SessionContext) error {
		return db.FindByID(ctx, COLLECTION, c.Params("id"), user)
	})
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}
