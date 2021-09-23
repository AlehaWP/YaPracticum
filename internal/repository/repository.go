package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/AlehaWP/YaPracticum.git/internal/global"
	"github.com/AlehaWP/YaPracticum.git/internal/shorter"
)

var serializeURLRepo func(global.Repository)

//UrlsData repository of urls. Realize Repository interface.
type ServerRepo struct {
	connStr string
	db      *sql.DB
	ctx     context.Context
}

type UsersRepo struct {
	Data      map[string]int
	CurrentID int
}

func (s *ServerRepo) SaveURL(url []byte, baseURL, userID string) (string, error) {
	db := s.db
	ctx, cancelfunc := context.WithTimeout(s.ctx, 5*time.Second)
	defer cancelfunc()
	r := shorter.MakeShortner(url)
	q := `INSERT INTO urls (
		shorten_url,
		url,
		base_url,
		user_id
	) VALUES ($1,$2,$3, (SELECT COALESCE(id, 0) FROM users where user_enc_id=$4))
	ON CONFLICT (shorten_url) DO NOTHING`
	if _, err := db.ExecContext(ctx, q, r, string(url), baseURL, userID); err != nil {
		return "", err
	}
	return baseURL + r, nil
}

func (s *ServerRepo) GetURL(id string) (string, error) {
	db := s.db
	ctx, cancelfunc := context.WithTimeout(s.ctx, 5*time.Second)
	defer cancelfunc()
	q := `SELECT url FROM urls WHERE shorten_url=$1`
	var url string
	row := db.QueryRowContext(ctx, q, id)

	if err := row.Scan(&url); err != nil {
		return "", err
	}
	return url, nil
}

func (s *ServerRepo) GetUserURLs(userID string) []global.URLs {
	// ud := s.URLsData
	// m := make([]global.URLs, 0)
	// for key, value := range ud {
	// 	if value[2] == userID {
	// 		m = append(m, global.URLs{
	// 			ShortURL:    value[1] + key,
	// 			OriginalURL: value[0],
	// 		})
	// 	}
	// }

	// return m
	return nil
}

func (s *ServerRepo) FindUser(key string) (finded bool) {
	// ur := s.Users
	// if _, ok := ur.Data[key]; ok {
	// 	return true
	// }
	return false
}

func (s *ServerRepo) CreateUser() (string, error) {
	// ur := &s.Users
	// id := ur.getNewID()
	// newKey, err := encription.EncriptInt(id)
	// if err != nil {
	// 	return "", err
	// }
	// ur.Data[newKey] = id
	// // serializeURLRepo(s)
	// return newKey, nil
	return "", nil
}
