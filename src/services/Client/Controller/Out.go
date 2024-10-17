package client

import (
	"api/src/shared"
	"log"

	"github.com/gofiber/fiber/v2"
)

type OutRequest struct {
	Token string `json:"token"`
}

// Выход из аккаунта
func (controller ClientController) out(service fiber.Router) {
	service.Post("/out", func(c *fiber.Ctx) error {
		request := new(OutRequest)
		err := c.BodyParser(&request)
		if err != nil || request.Token == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		//Если токен просрочен
		if !shared.CheckTokenLifeTime(request.Token) {
			return c.SendString("Данный токен уже просрочен!")
		}

		//Добавление токена авторизации в чс
		shared.Black_list_tokens = append(shared.Black_list_tokens, request.Token)

		log.Println("Произошёл выход из системы для пользователя с токеном: " + request.Token)
		log.Println(shared.Black_list_tokens)
		return c.SendString("Вы вышли из системы")
	})
}
