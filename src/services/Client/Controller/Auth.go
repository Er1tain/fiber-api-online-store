package client

import (
	client_models "api/src/services/Client/Models"
	"api/src/shared"
	"api/src/shared/responses"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Эндпоинт аутентификации
func (controller ClientController) auth(service fiber.Router) {
	service.Post("/auth", func(c *fiber.Ctx) error {
		log.Println("Выполнен запрос на аутентификацию!")

		request := new(client_models.AuthData)
		err := c.BodyParser(&request)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		//Поиск в БД записи с соответствующими почтой и паролем
		surname, name, res := client_models.FindClient(request.Email, request.Password)
		if res {
			//Создание jwt-токена
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"email": request.Email, "surname": surname, "name": name, "createdAt": time.Now(),
			})

			tokenString, err := token.SignedString(shared.HmacSampleSecret)

			if err == nil {
				return c.Status(fiber.StatusOK).JSON(responses.SendJWT{
					Token: tokenString,
				})
			}
		}

		return c.Status(fiber.ErrBadRequest.Code).SendString("Неверный логин и/или пароль!")
	})
}
