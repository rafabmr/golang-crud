package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"rafaelbautista.ml/crud/models"
)

func CrearUsuario(c *fiber.Ctx, db *gorm.DB) error {
	newUser := models.User{Name: "Ejemplo Usuario", Age: 25}
	db.Create(&newUser)
	return c.SendString("usuario insertado: " + newUser.Name)
}

func HolaMundo(c *fiber.Ctx) error {
	return c.SendString("Hola mundo")
}

func Saludo(c *fiber.Ctx) error {
	return c.SendString("¡Hola desde la ruta de saludo!")
}

func Vista(c *fiber.Ctx) error {
	data := models.User{
		Name: "Rafael",
	}
	return c.Render("views/index.html", data)
}
