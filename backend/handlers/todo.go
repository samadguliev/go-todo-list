package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samadguliev/go-todo-list/backend/database"
	"github.com/samadguliev/go-todo-list/backend/models"
)

func TodoList(c *fiber.Ctx) error {
	var todoList []models.TodoAPI

	database.DB.Db.Model(&models.Todo{}).Select([]string{"ID", "task", "is_done"}).Find(&todoList)

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

func UpdateItem(c *fiber.Ctx) error {
	todo := new(models.Todo)
	id := c.Params("ID")

	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Where("id = ?", id).Updates(&todo)
	return c.SendStatus(200)
}

func DeleteItem(c *fiber.Ctx) error {
	id := c.Params("ID")
	var todo models.Todo

	result := database.DB.Db.Delete(&todo, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
