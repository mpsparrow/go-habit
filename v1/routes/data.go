package routes

import "github.com/gofiber/fiber/v2"

func DataRoute(group fiber.Router) {
	group.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Data Route")
	})

	group.Get("/:userid(int)", func(c *fiber.Ctx) error {
		return c.SendString("Data Route")
	})

	group.Get("/:userid(int)", func(c *fiber.Ctx) error {
		return c.SendString("Data Route")
	})

	group.Get("/:userid(int)", func(c *fiber.Ctx) error {
		return c.SendString("Data Route")
	})
}
