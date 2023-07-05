package routes

import (
	"context"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mpsparrow/go-habit/v1/components"
)

func IDRoute(group fiber.Router) {
	group.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("ID Route")
	})

	group.Get("/sessions", func(c *fiber.Ctx) error {
		return c.SendString("Profile Info Basic")
	})

	group.Get("/logout", func(c *fiber.Ctx) error {
		return c.SendString("Logout")
	})

	group.Get("/logout/all", func(c *fiber.Ctx) error {
		return c.SendString("Logout of all sessions")
	})

	group.Delete("/delete", func(c *fiber.Ctx) error {
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
		return components.DeleteUser(c, conn, id)
	})
}
