package routes

import (
	"context"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mpsparrow/go-habit/v1/components"
)

func HabitsRoute(group fiber.Router) {
	group.Post("/create", func(c *fiber.Ctx) error {
		/*
			{
			  "name": "Exercise",
			  "description": "Daily exercise routine",
			  "startDate": "2023-07-06T09:00:00Z",
			  "intervals": 7
			}
		*/

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

		habit := components.Habit{}
		if err := c.BodyParser(&habit); err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}
		return components.CreateHabit(c, conn, id, habit)
	})

	group.Get("/list", func(c *fiber.Ctx) error {
		userID := c.Params("userid(int)")
		id, err := strconv.Atoi(userID)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID",
			})
		}

		conn, err := connectDB()
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close(context.Background())

		return components.GetHabits(c, conn, id)
	})

	group.Delete("/:habitid(int)/delete", func(c *fiber.Ctx) error {
		conn, err := connectDB()
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close(context.Background())

		userID := c.Params("userid(int)")
		userid, err := strconv.Atoi(userID)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID",
			})
		}

		habitID := c.Params("habitid(int)")
		habitid, err := strconv.Atoi(habitID)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID",
			})
		}
		return components.DeleteHabit(c, conn, userid, habitid)
	})

	group.Get("/:habitid(int)", func(c *fiber.Ctx) error {
		conn, err := connectDB()
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close(context.Background())

		userID := c.Params("userid(int)")
		userid, err := strconv.Atoi(userID)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID",
			})
		}

		habitID := c.Params("habitid(int)")
		habitid, err := strconv.Atoi(habitID)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid habit ID",
			})
		}

		return components.GetHabit(c, conn, userid, habitid)
	})

}
