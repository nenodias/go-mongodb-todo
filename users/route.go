package users

import "github.com/gofiber/fiber/v3"

func SetRoutes(r fiber.Router) {
	users := r.Group("/users")
	users.Post("/", addUser)
	users.Get("/", getAll)
	users.Get("/:id", getById)
	users.Put("/:id", updateUserById)
	users.Patch("/:id", updateUserById)
}
