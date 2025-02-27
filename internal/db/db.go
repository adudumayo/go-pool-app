package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://pooluser:password@localhost/pooldb")
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	return conn, nil
}
