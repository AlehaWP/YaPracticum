package repository

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

//CheckDBConnection trying connect to db.
func (s *ServerRepo) Close() {
	s.db.Close()
}

//CheckDBConnection trying connect to db.
func (s *ServerRepo) CheckDBConnection() error {
	err := s.db.PingContext(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (s *ServerRepo) createTables() error {
	db := s.db
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	q := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		user_uuid VARCHAR(36),
		user_enc_id VARCHAR(36),
		date_add timestamp
	)`
	if _, err := db.ExecContext(ctx, q); err != nil {
		return err
	}

	q = `CREATE TABLE IF NOT EXISTS urls (
		id SERIAL NOT NULL,
		shorten_url VARCHAR(32) UNIQUE,
		url VARCHAR(255),
		base_url VARCHAR(255),
		user_id INTEGER REFERENCES users (id),
		date_add TIMESTAMPTZ NOT NULL DEFAULT NOW()
	)`
	if _, err := db.ExecContext(ctx, q); err != nil {
		return err
	}
	return nil
}

func NewServerRepo(c string) (*ServerRepo, error) {
	db, err := sql.Open("postgres", c)
	if err != nil {
		return nil, err
	}
	sr := &ServerRepo{
		connStr: c,
		db:      db,
		ctx:     context.Background(),
	}
	if err := sr.CheckDBConnection(); err != nil {
		return nil, err
	}

	if err := sr.createTables(); err != nil {
		return nil, err
	}
	return sr, nil
}
