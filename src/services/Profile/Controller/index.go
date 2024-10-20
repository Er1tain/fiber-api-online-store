package profile

import "github.com/gofiber/fiber/v2"

type ProfileController struct {
	app     *fiber.App
	service fiber.Router
}

//Запуск хранилища пользовательских данных
func Start(app *fiber.App) {
	store := ProfileController{app: app}

	//Маршрутизация
	api := app.Group("/api/client")
	profile := api.Group("/profile", func(c *fiber.Ctx) error {
		c.Set("Service", "profile")
		return c.Next()
	})
	store.service = profile
	store.upload()
}
