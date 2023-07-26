package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samadguliev/go-todo-list/backend/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.TodoList)

	app.Post("/add-item", handlers.AddItem)
}
