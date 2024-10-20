package profile

import (
	profile "api/src/services/Profile"
	profile_middleware "api/src/services/Profile/Controller/middleware"
	"api/src/shared"
	"api/src/shared/validation"
	"io/ioutil"
	"log"

	"github.com/gofiber/fiber/v2"
)

type RequestDataStoreUpload struct {
	Token string `json:"token"`
}

func (controller ProfileController) upload() {
	controller.service.Use("/upload_avatar", func(c *fiber.Ctx) error {
		log.Println("Проверка перед обновлением аватарки....")

		Token := c.FormValue("token")
		if Token == "" {
			log.Println("Ошибка при чтении тела запроса!_________")
			return c.SendStatus(fiber.StatusBadRequest)
		}

		//Токен сгенерирован данным сервисом?
		if !validation.ValidateJWT(Token) {
			return c.Status(fiber.StatusBadRequest).SendString("Токен сгенерирован сторонним сервисом!")
		}

		//Если токен просрочен
		if !shared.CheckTokenLifeTime(Token) {
			return c.SendString("Данный токен уже просрочен!")
		}

		email, err := profile_middleware.GetUserEmail(Token)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Не удалось выяснить email(")
		}

		c.Locals("email", email)
		return c.Next()
	})

	store := profile.NewStorage()
	controller.service.Put("/upload_avatar", func(c *fiber.Ctx) error {
		log.Println("Обновление аватарки пользователя!___________")

		user_email := c.Locals("email").(string)
		avatar, err := c.FormFile("image")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Ошибка получения изображения")
		}

		// Чтение содержимого файла
		f, err := avatar.Open()
		if err != nil {
			return err
		}
		defer f.Close()

		// Чтение данных в бинарный формат
		data, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}

		//Сохранение фото с именем почты пользователя в хранилище
		if store.Save(user_email, data) {
			return c.Status(fiber.StatusAccepted).SendString("Установлен новый автар для вас!")
		}

		return c.Status(fiber.StatusBadRequest).SendString("Не удалось обновить аватар(")
	})

}
