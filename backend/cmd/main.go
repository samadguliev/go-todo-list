package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samadguliev/go-todo-list/backend/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
