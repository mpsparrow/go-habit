package components

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

type Habit struct {
	HabitID     int
	UserID      int
	Name        string
	Description string
	StartDate   time.Time
	Intervals   int
}

func CreateHabit(c *fiber.Ctx, conn *pgx.Conn, userid int, habit Habit) error {
	query := `
		INSERT INTO Habits (UserID, Name, Description, StartDate, Intervals)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING HabitID
	`

	var habitID int
	err := conn.QueryRow(context.Background(), query, userid, habit.Name, habit.Description, habit.StartDate, habit.Intervals).Scan(&habitID)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create habit",
		})
	}

	return c.JSON(fiber.Map{
		"message":  "Habit created successfully",
		"habit_id": habitID,
	})
}

func DeleteHabit(c *fiber.Ctx, conn *pgx.Conn, userid int, habitid int) error {
	query := `
		DELETE FROM Habits
		WHERE UserID = $1 AND HabitID = $2
	`

	result, err := conn.Exec(context.Background(), query, userid, habitid)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete habit",
		})
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Habit not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Habit deleted successfully",
		"habitID": habitid,
	})
}

func GetHabit(c *fiber.Ctx, conn *pgx.Conn, userid int, habitid int) error {
	query := `
		SELECT HabitID, UserID, Name, Description, StartDate, Intervals
		FROM Habits
		WHERE UserID = $1 AND HabitID = $2
	`

	row := conn.QueryRow(context.Background(), query, userid, habitid)

	habit := Habit{}
	err := row.Scan(&habit.HabitID, &habit.UserID, &habit.Name, &habit.Description, &habit.StartDate, &habit.Intervals)
	if err != nil {
		if err == pgx.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Habit not found",
			})
		}
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch habit",
		})
	}

	return c.JSON(habit)
}

func GetHabits(c *fiber.Ctx, conn *pgx.Conn, userid int) error {
	query := `
		SELECT HabitID, UserID, Name, Description, StartDate, Intervals
		FROM Habits
		WHERE UserID = $1
	`

	rows, err := conn.Query(context.Background(), query, userid)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}
	defer rows.Close()

	habits := []Habit{}
	for rows.Next() {
		habit := Habit{}
		err := rows.Scan(&habit.HabitID, &habit.UserID, &habit.Name, &habit.Description, &habit.StartDate, &habit.Intervals)
		if err != nil {
			log.Println(err)
			continue
		}
		habits = append(habits, habit)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(habits)
}
