package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samadguliev/go-todo-list/backend/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/todos", handlers.TodoList)
	app.Post("/todos", handlers.AddItem)
	app.Put("/todos/:ID", handlers.UpdateItem)
	app.Delete("/todos/:ID", handlers.DeleteItem)
}
