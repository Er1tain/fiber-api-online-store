package client

import "github.com/gofiber/fiber/v2"

// Выход из аккаунта
func (controller ClientController) out(service fiber.Router) {
	service.Post("/out", func(c *fiber.Ctx) error {
		return c.SendString("Вы вышли из системы")
	})
}
