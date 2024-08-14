package main

import (
	client "api/src/services/Client/Controller"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	api := fiber.New()

	client.Start(api)

	log.Fatal(api.Listen(":8000"))
}
