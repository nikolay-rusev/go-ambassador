package main

import (
    "log"
    "ambassador/src/database"
    "github.com/gofiber/fiber/v2"
)

func main() {
    database.Connect()
    database.AutoMigrate()

    app := fiber.New()

    app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    log.Fatal(app.Listen(":8000"))
}
