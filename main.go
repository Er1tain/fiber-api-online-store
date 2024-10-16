package main

import (
	client "api/src/services/Client/Controller"
	"api/src/shared"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title Fiber-API-Online-Store
// @version 1.0
// @description Golang GoFiber swagger auto generate step by step by swaggo
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
func main() {
	api := fiber.New()
	api.Get("/docs/*", swagger.HandlerDefault)
	client.Start(api)

	go func() {
		for {
			log.Println(shared.Black_list_tokens)
			time.Sleep(time.Second * 60)
		}
	}()

	log.Fatal(api.Listen(":8000"))
}
