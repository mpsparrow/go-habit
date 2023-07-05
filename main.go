package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/mpsparrow/go-habit/v1/routes"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(favicon.New())
	app.Use(recover.New())
	app.Use(limiter.New(limiter.Config{
		Max:               20,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	app.Get("/metrics", monitor.New(monitor.Config{Title: "Backend Metrics"}))

	v1 := app.Group("/api/v1")

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API v1")
	})

	// User Routes
	users := v1.Group("/users")
	routes.UserRoute(users)

	// Specific User
	id := users.Group("/:userid(int)")
	routes.IDRoute(id)

	// Specific User Data
	data := id.Group("/data")
	routes.DataRoute(data)

	// Habit Data
	habits := data.Group("/habits")
	routes.HabitsRoute(habits)

	// Specific User Profile
	profile := id.Group("/profile")
	routes.ProfileRoute(profile)

	// Edit Specific User Settings
	edit := profile.Group("/edit")
	routes.EditRoute(edit)

	log.Fatal(app.Listen(":5000"))
}
