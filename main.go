package main

import "github.com/gofiber/fiber/v2"

type Ninja struct {
	Name string
	Weapon string
}

func getNinja(ctx *fiber.Ctx) error {
	wallace := Ninja {
		Name: "Samarjit",
		Weapon: "M416",
	}

	return ctx.Status(fiber.StatusOK).JSON(wallace)
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/ninja", getNinja)

	app.Listen(":3000")
}