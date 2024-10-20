package profile

import (
	profile "api/src/services/Profile"
	profile_middleware "api/src/services/Profile/Controller/middleware"
	"api/src/shared"
	"api/src/shared/validation"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func (controller ProfileController) download() {
	controller.service.Use("/download_avatar", func(c *fiber.Ctx) error {
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

	storage := profile.NewStorage()
	controller.service.Post("/download_avatar", func(c *fiber.Ctx) error {
		user_email := c.Locals("email").(string)

		//Аватарка хранится под именем хэшированной почты
		key_byte := []byte(user_email)
		md5Hash := md5.Sum(key_byte)
		key := hex.EncodeToString(md5Hash[:]) + ".jpg"

		fmt.Println(key)

		bytes_avatar, res := storage.Load(key)
		if res {
			c.Set("Content-Type", "image/jpg")
			return c.Send(bytes_avatar)
		}
		return c.SendStatus(fiber.StatusNotFound)
	})
}
