package main

import "github.com/gofiber/fiber/v2"

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

	app.Get("/ninja", getNinja)
	app.Post("/ninja", createNinja)

	app.Listen(":3000")
}