package client

import (
	client_models "api/src/services/Client/Models"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Удаление аккаунта
func (controller ClientController) delete(service fiber.Router) {
	service.Delete("/delete", func(c *fiber.Ctx) error {
		log.Println("Выполнен запрос на удаление аккаунта!")

		request := new(client_models.AuthData)
		err := c.BodyParser(&request)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		//Удаление учётной записи из БД
		if client_models.DeleteClient(*request) {
			return c.Status(fiber.StatusOK).SendString("Учётная запись успешно удалена!")
		}

		return c.Status(fiber.ErrBadRequest.Code).SendString("Не получилось удалить учётную запись(")
	})
}
