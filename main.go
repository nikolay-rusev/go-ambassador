package main

import (
    "log"
    "ambassador/src/database"
    "ambassador/src/routes"
    "github.com/gofiber/fiber/v2"
)

func main() {
    database.Connect()
    database.AutoMigrate()

    app := fiber.New()

    routes.Setup(app)

    log.Fatal(app.Listen(":8000"))
}
