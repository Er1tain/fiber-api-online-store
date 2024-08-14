package client

import "github.com/gofiber/fiber/v2"

// Удаление аккаунта
func (controller ClientController) delete(service fiber.Router) {
	service.Delete("/delete", func(c *fiber.Ctx) error {
		return c.SendString("Ваш аккаунт удалён(")
	})
}
