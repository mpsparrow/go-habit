package components

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

type User struct {
	UserName  string
	Pass      string
	FirstName string
	LastName  string
	Email     string
}

type UserFull struct {
	UserID    int
	UserName  string
	Pass      string
	FirstName string
	LastName  string
	Email     string
}

func CreateUser(c *fiber.Ctx, conn *pgx.Conn, user User) error {
	query := `
		INSERT INTO Users (UserName, Pass, FirstName, LastName, Email)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING UserID
	`

	var userID int
	err := conn.QueryRow(context.Background(), query, user.UserName, user.Pass, user.FirstName, user.LastName, user.Email).Scan(&userID)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User registered successfully",
		"userid":  userID,
	})
}

func DeleteUser(c *fiber.Ctx, conn *pgx.Conn, userid int) error {
	deleteQuery := `
		DELETE FROM Users WHERE UserID = $1;
	`

	result, err := conn.Exec(context.Background(), deleteQuery, userid)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to delete user",
		})
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   "User not found",
			"message": "The user has already been deleted or does not exist",
			"userid":  userid,
		})
	}

	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
		"userid":  userid,
	})
}

func ListUsers(c *fiber.Ctx, conn *pgx.Conn) error {
	query := `
		SELECT UserID, UserName, Pass, FirstName, LastName, Email
		FROM Users
	`

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve users",
		})
	}
	defer rows.Close()

	users := make([]UserFull, 0)

	for rows.Next() {
		var user UserFull
		err := rows.Scan(
			&user.UserID,
			&user.UserName,
			&user.Pass,
			&user.FirstName,
			&user.LastName,
			&user.Email,
		)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve user data",
			})
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve users",
		})
	}

	return c.JSON(users)
}
