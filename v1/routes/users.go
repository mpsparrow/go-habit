package routes

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mpsparrow/go-habit/v1/components"
)

func UserRoute(group fiber.Router) {
	group.Post("/create", func(c *fiber.Ctx) error {
		/*
			{
			  "UserName": "john.doe",
			  "Pass": "some-password-hash",
			  "FirstName": "John",
			  "LastName": "Doe",
			  "Email": "john.doe@example.com"
			}
		*/

		conn, err := connectDB()
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close(context.Background())

		user := components.User{}
		if err := c.BodyParser(&user); err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}
		return components.CreateUser(c, conn, user)
	})

	group.Get("/list", func(c *fiber.Ctx) error {
		conn, err := connectDB()
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close(context.Background())

		return components.ListUsers(c, conn)
	})
}
