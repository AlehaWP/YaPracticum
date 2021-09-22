package repository

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type QueryToDB func() error

//CheckDBConnection trying connect to db.
func (s *ServerRepo) CheckDBConnection() error {
	checkFunc := func(db *sql.DB) error {
		err := db.PingContext(context.Background())
		if err != nil {
			return err
		}
		return nil
	}
	err := s.newConnect(checkFunc)
	return err

}

func (s *ServerRepo) newConnect(qf ...func(*sql.DB) error) error {

	db, err := sql.Open("postgres", s.connStr)
	if err != nil {
		return err
	}

	for _, q := range qf {
		if err := q(db); err != nil {
			return err
		}

	}

	defer db.Close()

	return nil
}

func createTables(db *sql.DB) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	if _, err := db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS URLS (
						shorten_url char(32),
						url char(255),
						base_url char(255),
						user_id char(32)
	)`); err != nil {
		return err
	}

	if _, err := db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS USERS (
			id int,
			user_id char(32)
	)`); err != nil {
		return err
	}
	return nil
}

func NewServerRepo(c string) (*ServerRepo, error) {
	servRepo := &ServerRepo{
		URLsData: make(map[string][]string),
		Users: UsersRepo{
			Data:      make(map[string]int),
			CurrentID: 0,
		},
		connStr: c,
	}
	if err := servRepo.CheckDBConnection(); err != nil {
		return nil, err
	}
	if err := servRepo.newConnect(createTables); err != nil {
		return nil, err
	}
	return servRepo, nil
	// return &ServerRepo{
	// 	connStr: c,
	// }
}
