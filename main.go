package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/Taker-Academy/kedubak-LoicCODING/database"
)

func main() {
	app := fiber.New()
	client := database.databaseToConnect()

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("get / called")
        return c.SendString("Hello, World!")
    })
	fmt.Println(client.NumberSessionsInProgress())
	app.Listen(":8080")
}