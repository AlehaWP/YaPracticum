package repository

import (
	"github.com/AlehaWP/YaPracticum.git/internal/shorter"
)

type Key string

//Repository Interface bd urls
type Repository interface {
	GetURL(string) (string, bool)
	SaveURL([]byte) string
}

//UrlsData Repository of urls. Realize Repository interface
type UrlsData map[string][]byte

func (u *UrlsData) SaveURL(url []byte) string {
	r := shorter.MD5(url)
	(*u)[r] = url
	return r
}

func (u *UrlsData) GetURL(id string) (string, bool) {
	if r, ok := (*u)[id]; ok {
		return string(r), true
	}
	return "", false
}
