package client

import "github.com/gofiber/fiber/v2"

// Эндпоинт регистрации
func (controller ClientController) reg(service fiber.Router) {
	service.Post("/reg", func(c *fiber.Ctx) error {
		return c.SendString("Зарегестрирован")
	})
}
