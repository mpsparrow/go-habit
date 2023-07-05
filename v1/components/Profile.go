package components

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

type Basic struct {
	UserID   int
	UserName string
	Email    string
}

type Full struct {
	UserID    int
	UserName  string
	Pass      string
	FirstName string
	LastName  string
	Email     string
}

func ProfileBasic(c *fiber.Ctx, conn *pgx.Conn, userid int) error {
	query := `
		SELECT UserID, UserName, Email
		FROM Users
		WHERE UserID = $1
	`

	var userData Basic

	err := conn.QueryRow(context.Background(), query, userid).Scan(
		&userData.UserID,
		&userData.UserName,
		&userData.Email,
	)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve user data",
		})
	}

	return c.JSON(userData)
}

func ProfileFull(c *fiber.Ctx, conn *pgx.Conn, userid int) error {
	query := `
		SELECT UserID, UserName, Pass, FirstName, LastName, Email
		FROM Users
		WHERE UserID = $1
	`

	var userData Full

	err := conn.QueryRow(context.Background(), query, userid).Scan(
		&userData.UserID,
		&userData.UserName,
		&userData.Pass,
		&userData.FirstName,
		&userData.LastName,
		&userData.Email,
	)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve user data",
		})
	}

	return c.JSON(userData)
}
