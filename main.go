package main

import (
	"github.com/gofiber/fiber/v2"
	"rafaelbautista.ml/crud/config"
	"rafaelbautista.ml/crud/routes"
)

func main() {
	db := config.ConnectDB()

	app := fiber.New()
	routes.Routes(app, db)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
