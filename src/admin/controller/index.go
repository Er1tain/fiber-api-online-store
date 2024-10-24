package admin

import (
	"github.com/gofiber/fiber/v2"
)

type Admin struct {
	app     *fiber.App
	service fiber.Router
}

// Запуск админки
func Start(app *fiber.App) {
	controller := Admin{app: app}

	admin := app.Group("/admin", func(c *fiber.Ctx) error {
		c.Set("Service", "admin")
		return c.Next()
	})

	controller.service = admin

	controller.openAdminPanel()

}
