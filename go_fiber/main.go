package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Ninja struct {
	Name   string `json:"name"`
	Weapon string `json:"weapon"`
}

func getNinja(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(ninja)
}

var ninja Ninja

func createNinja(ctx *fiber.Ctx) error {
	body := new(Ninja)
	err := ctx.BodyParser(body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	ninja = Ninja{
		Name:   body.Name,
		Weapon: body.Weapon,
	}
	return ctx.Status(fiber.StatusOK).JSON(ninja)
}

func main() {
	fmt.Println("bok")
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello world")
	})
	ninjaApp := app.Group(("/ninja"))
	ninjaApp.Get("", getNinja)
	ninjaApp.Post("", createNinja)
	app.Listen(":8000")
}
