package routes

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func connectDB() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://admin:password@192.168.2.3:54320/backend")
	if err != nil {
		return nil, err
	}

	return conn, nil
}
