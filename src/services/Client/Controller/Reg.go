package client

import (
	client_models "api/src/services/Client/Models"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Эндпоинт регистрации
func (controller ClientController) reg(service fiber.Router) {
	service.Post("/reg", func(c *fiber.Ctx) error {
		log.Println("Выполнен запрос на регистрацию!")

		response := new(client_models.Client)
		err := c.BodyParser(&response)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		//Создание учётной записи в БД
		if client_models.CreateClient(*response) {
			return c.Status(fiber.StatusOK).JSON(response)
		}

		return c.Status(fiber.ErrBadRequest.Code).SendString("Данный email уже используется")

	})
}
