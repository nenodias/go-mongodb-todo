package tasks

import "github.com/gofiber/fiber/v3"

func SetRoutes(r fiber.Router) {
	users := r.Group("/tasks")
	users.Post("/", addItem)
	users.Get("/", getAll)
	users.Get("/:id", getById)
	users.Put("/:id", updateById)
	users.Patch("/:id", updateById)
	users.Delete("/:id", deleteItem)
}
