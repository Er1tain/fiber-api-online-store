package client

import "github.com/gofiber/fiber/v2"

// Эндпоинт аутентификации
func (controller ClientController) auth(service fiber.Router) {
	service.Post("/auth", func(c *fiber.Ctx) error {
		return c.SendString("Вы успешно вошли!")
	})
}
