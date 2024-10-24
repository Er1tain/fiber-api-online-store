package shop

import "github.com/gofiber/fiber/v2"

type ShopController struct {
	app     *fiber.App
	service fiber.Router
}

//Запуск магазина одежды
func Start(app *fiber.App) {
	controller := ShopController{app: app}

	api := app.Group("/api")
	shop := api.Group("/shop", func(c *fiber.Ctx) error {
		c.Set("Service", "shop")
		return c.Next()
	})

	controller.service = shop

	//middleware
	controller.authMiddleware()

	controller.getListClothes()
	controller.addToBasket()
	controller.pay()
}
