package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"rafaelbautista.ml/crud/controllers"
)

func Routes(app *fiber.App, db *gorm.DB) {

	app.Get("/", func(c *fiber.Ctx) error {
		return controllers.HolaMundo(c)
	})

	app.Get("/saludo", func(c *fiber.Ctx) error {
		return controllers.Saludo(c)
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		return controllers.Vista(c)
	})

	app.Get("/user", func(c *fiber.Ctx) error {
		return controllers.CrearUsuario(c, db)
	})

	app.Get("/test/texto", func(c *fiber.Ctx) error {
		return controllers.Texto(c)
	})

	app.Get("/test/json", func(c *fiber.Ctx) error {
		return controllers.Json(c)
	})

	app.Get("/test/html", func(c *fiber.Ctx) error {
		return controllers.Html(c)
	})

}
