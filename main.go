package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/nenodias/go-mongodb-todo/tags"
	"github.com/nenodias/go-mongodb-todo/tasks"
	"github.com/nenodias/go-mongodb-todo/users"
)

func main() {
	app := fiber.New()
	v1 := app.Group("/v1")
	users.SetRoutes(v1)
	tasks.SetRoutes(v1)
	tags.SetRoutes(v1)
	log.Fatal(app.Listen(":8000"))
}
