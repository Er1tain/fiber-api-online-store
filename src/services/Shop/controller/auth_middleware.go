package shop

import (
	profile_middleware "api/src/services/Profile/Controller/middleware"
	"api/src/shared"
	"api/src/shared/validation"

	"github.com/gofiber/fiber/v2"
)

type tokenForShopService struct {
	Token string `json:"token"`
}

func (controller ShopController) authMiddleware() {
	controller.service.Use('/', func(c *fiber.Ctx) error {
		request := new(tokenForShopService)
		err := c.BodyParser(&request)
		if err != nil || request.Token == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Для покупки товаров необходимо авторизоваться!")
		}

		//Токен сгенерирован данным сервисом?
		if !validation.ValidateJWT(request.Token) {
			return c.Status(fiber.StatusBadRequest).SendString("Токен сгенерирован сторонним сервисом!")
		}

		//Если токен просрочен
		if !shared.CheckTokenLifeTime(request.Token) {
			return c.SendString("Данный токен уже просрочен!")
		}

		email, err := profile_middleware.GetUserEmail(request.Token)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Не удалось выяснить email(")
		}

		surname, name, err := profile_middleware.GetUserSurnameName(request.Token)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Не удалось выяснить фамилию и имяпользователя(")
		}

		c.Locals("email", email)
		c.Locals("surname", surname)
		c.Locals("name", name)
		return c.Next()

	})
}
