package client

import (
	"github.com/gofiber/fiber/v2"
)

type ClientController struct {
	app *fiber.App
}

// Инициализация клиентского контроллера
func Start(app *fiber.App) {
	controller := ClientController{app: app}

	//Группируем обработчики контроллера
	api := app.Group("/api")
	service := api.Group("/client", func(c *fiber.Ctx) error {
		c.Set("Service", "client")
		return c.Next()
	})

	//Запуск обработчиков(эндпоинтов) контроллера
	controller.reg(service)
	controller.auth(service)
	controller.delete(service)
	controller.out(service)
}

// Эндпоинт регистрации
func (controller ClientController) reg(service fiber.Router) {
	service.Post("/reg", func(c *fiber.Ctx) error {
		return c.SendString("Зарегестрирован")
	})
}

// Эндпоинт аутентификации
func (controller ClientController) auth(service fiber.Router) {
	service.Post("/auth", func(c *fiber.Ctx) error {
		return c.SendString("Вы успешно вошли!")
	})
}

// Удаление аккаунта
func (controller ClientController) delete(service fiber.Router) {
	service.Delete("/delete", func(c *fiber.Ctx) error {
		return c.SendString("Ваш аккаунт удалён(")
	})
}

// Выход из аккаунта
func (controller ClientController) out(service fiber.Router) {
	service.Post("/out", func(c *fiber.Ctx) error {
		return c.SendString("Вы вышли из системы")
	})
}
