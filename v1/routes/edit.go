package routes

import "github.com/gofiber/fiber/v2"

func EditRoute(group fiber.Router) {
	group.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Edit Route")
	})

	group.Get("/username", func(c *fiber.Ctx) error {
		return c.SendString("Edit Username")
	})

	group.Get("/password", func(c *fiber.Ctx) error {
		return c.SendString("Edit Password")
	})

	group.Get("/email", func(c *fiber.Ctx) error {
		return c.SendString("Edit Email")
	})

	group.Get("/firstname", func(c *fiber.Ctx) error {
		return c.SendString("Edit First Name")
	})

	group.Get("/lastname", func(c *fiber.Ctx) error {
		return c.SendString("Edit Last Name")
	})
}
