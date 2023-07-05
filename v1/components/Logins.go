package components

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

type Login struct {
	UserID int
	IP     string
}

func insertLogin(conn *pgx.Conn, login Login) error {
	query := `
		INSERT INTO Logins (UserID, IP)
		VALUES ($1, $2)
	`

	_, err := conn.Exec(context.Background(), query, login.UserID, login.IP)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func CreateLogin(c *fiber.Ctx, conn *pgx.Conn) error {
	login := Login{
		UserID: 1,
		IP:     c.IP(),
	}

	err := insertLogin(conn, login)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
