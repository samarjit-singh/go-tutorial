package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Ninja struct {
	Name string
	Weapon string
}

func getNinja(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(ninja)
}

var ninja Ninja

func createNinja(ctx *fiber.Ctx) error {
	body := new (Ninja)
	err := ctx.BodyParser(body)
	
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	
	ninja = Ninja {
		Name : body.Name,
		Weapon : body.Weapon,
	}
	
	return ctx.Status(fiber.StatusOK).JSON(ninja)
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Use(logger.New())
	// app.Use(requestid.New())

	ninjaApp := app.Group("/ninja")
	ninjaApp.Get("", getNinja)
	ninjaApp.Post("", createNinja)

	app.Listen(":3000")
}