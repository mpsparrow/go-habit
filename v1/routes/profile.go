package routes

import (
	"context"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mpsparrow/go-habit/v1/components"
)

func ProfileRoute(group fiber.Router) {
	group.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Profile Route")
	})

	group.Get("/basic", func(c *fiber.Ctx) error {
		conn, err := connectDB()
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close(context.Background())

		userID := c.Params("userid(int)")

		id, err := strconv.Atoi(userID)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID",
			})
		}
		return components.ProfileBasic(c, conn, id)
	})

	group.Get("/full", func(c *fiber.Ctx) error {
		conn, err := connectDB()
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close(context.Background())

		userID := c.Params("userid(int)")

		id, err := strconv.Atoi(userID)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID",
			})
		}
		return components.ProfileFull(c, conn, id)
	})
}
