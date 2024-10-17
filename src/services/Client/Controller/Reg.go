package client

import (
	client_models "api/src/services/Client/Models"
	"api/src/shared"
	"api/src/shared/responses"
	"api/src/shared/validation"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Эндпоинт регистрации
func (controller ClientController) reg(service fiber.Router) {
	service.Post("/reg", func(c *fiber.Ctx) error {
		log.Println("Выполнен запрос на регистрацию!")

		request := new(client_models.Client)
		err := c.BodyParser(&request)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		if !validation.CheckEmail(request.Email) {
			return c.Status(fiber.StatusBadRequest).SendString("Некорректный email!")
		}

		//Создание учётной записи в БД
		if client_models.CreateClient(*request) {
			//Создание jwt-токена
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"email": request.Email, "surname": request.Surname, "name": request.Name, "createdAt": time.Now(),
			})

			tokenString, err := token.SignedString(shared.HmacSampleSecret)

			if err == nil {
				return c.Status(fiber.StatusOK).JSON(responses.SendJWT{
					Token: tokenString,
				})
			}
		}

		return c.Status(fiber.ErrBadRequest.Code).SendString("Данный email уже используется")

	})
}
