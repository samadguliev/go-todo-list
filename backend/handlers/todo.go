package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samadguliev/go-todo-list/backend/database"
	"github.com/samadguliev/go-todo-list/backend/models"
)

func TodoList(c *fiber.Ctx) error {
	var todoList []models.Todo

	database.DB.Db.Select([]string{"ID", "task"}).Find(&todoList)

	return c.Status(200).JSON(todoList)
}

func AddItem(c *fiber.Ctx) error {
	todo := new(models.Todo)

	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&todo)

	return c.Status(200).JSON(todo)
}
