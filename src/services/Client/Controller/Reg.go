package client

import (
	client_models "api/src/services/Client/Models"
	"api/src/shared"
	"api/src/shared/responses"
	"api/src/shared/validation"
	"log"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func sendRegCode(email_addressee string, code *int) bool {
	from := "sanyapridava@mail.ru"
	password := "e6Ds5aTimK0ZgutjLkSM"

	toList := []string{email_addressee}

	host := "smtp.mail.ru"

	port := "587"

	*code = rand.Intn(9999-1000+1) + 1000
	code_value := strconv.Itoa(*code)

	body := []byte("Код подтверждения: " + code_value)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, toList, body)

	return err == nil
}

// Эндпоинт регистрации
func (controller ClientController) reg(service fiber.Router) {
	code := 0
	service.Post("/reg", func(c *fiber.Ctx) error {
		log.Println("Выполнен запрос на получение кода подтверждения!")

		request := new(client_models.Client)
		err := c.BodyParser(&request)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		if !validation.CheckEmail(request.Email) {
			return c.Status(fiber.StatusBadRequest).SendString("Некорректный email!")
		}

		if sendRegCode(request.Email, &code) {
			return c.Status(fiber.StatusOK).SendString("Код подтверждения отправлен на почту!")
		}

		return c.Status(fiber.StatusBadRequest).SendString("Не удалось отправить код(")
	})

	service.Post("/reg/:code", func(c *fiber.Ctx) error {
		log.Println("Выполнен запрос на регистрацию!")

		if code == 0 || c.Params("code") != strconv.Itoa(code) {
			return c.Status(fiber.StatusBadRequest).SendString("Неверный код подтверждения(")
		}

		//Обнуляем 4хзначный код во избежании повторного использования
		code = 0

		request := new(client_models.Client)
		err := c.BodyParser(&request)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
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
