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
