package repository

import (
	"errors"

	"github.com/AlehaWP/YaPracticum.git/internal/shorter"
)

type Key string

var SerializeURLRepo func(map[string]string)

//UrlsData repository of urls. Realize Repository interface.
type URLRepo struct {
	data map[string]string
}

func (u *URLRepo) SaveURL(url []byte) string {
	r := shorter.MakeShortner(url)
	(*u).data[r] = string(url)
	SerializeURLRepo(u.Get())
	return r
}

func (u *URLRepo) GetURL(id string) (string, error) {
	if r, ok := (*u).data[id]; ok {
		return string(r), nil
	}
	return "", errors.New("not found")
}

func (u *URLRepo) Get() map[string]string {
	return u.data
}

func (u *URLRepo) ToSet() *map[string]string {
	return &u.data
}

// Init return obj with alocate data.
func Init() *URLRepo {
	return &URLRepo{
		data: make(map[string]string),
	}
}
