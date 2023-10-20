package controllers

import "github.com/gofiber/fiber/v2"

func Texto(c *fiber.Ctx) error {
	return c.SendString("Este es un texto")
}

func Json(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"estatus": "JSON",
		"clave":   "valor",
		"numero":  3,
	})
}

func Html(c *fiber.Ctx) error {
	return c.Render("views/index.html", nil)
}
