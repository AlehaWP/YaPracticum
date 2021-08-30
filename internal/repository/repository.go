package repository

import (
	"errors"

	"github.com/AlehaWP/YaPracticum.git/internal/shorter"
)

type Key string

//Repository Interface bd urls
type Repository interface {
	GetURL(string) (string, error)
	SaveURL([]byte) string
}

//UrlsData Repository of urls. Realize Repository interface
type UrlRepo map[string]string

func (u *UrlRepo) SaveURL(url []byte) string {
	r := shorter.MakeShortner(url)
	(*u)[r] = string(url)
	return r
}

func (u *UrlRepo) GetURL(id string) (string, error) {
	if r, ok := (*u)[id]; ok {
		return string(r), nil
	}
	return "", errors.New("Not found")
}
