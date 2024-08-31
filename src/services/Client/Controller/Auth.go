package client

import "github.com/gofiber/fiber/v2"

type AuthData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Эндпоинт аутентификации
func (controller ClientController) auth(service fiber.Router) {
	service.Post("/auth", func(c *fiber.Ctx) error {

		return c.SendString("Вы успешно вошли!")
	})
}
