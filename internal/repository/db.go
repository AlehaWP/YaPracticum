package repository

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

func (s *ServerRepo) CheckDBConnection(connStr string) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = db.PingContext(context.Background())
	if err != nil {
		return err
	}
	defer db.Close()

	return nil
}
