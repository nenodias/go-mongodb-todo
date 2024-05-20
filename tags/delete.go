package tags

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/nenodias/go-mongodb-todo/db"
)

func deleteItem(c fiber.Ctx) error {
	err := db.DeleteById(COLLECTION, c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	return c.SendStatus(http.StatusNoContent)
}
