package repository

import (
	"errors"

	encription "github.com/AlehaWP/YaPracticum.git/internal/Encription"
	"github.com/AlehaWP/YaPracticum.git/internal/global"
	"github.com/AlehaWP/YaPracticum.git/internal/shorter"
)

type Key string

var SerializeURLRepo func(global.Repository)

//UrlsData repository of urls. Realize Repository interface.
type URLRepo struct {
	data map[string][]string
}

func (u *URLRepo) SaveURL(url []byte, userID string) string {
	r := shorter.MakeShortner(url)
	(*u).data[r] = []string{string(url)}
	SerializeURLRepo(u)
	return r
}

func (u *URLRepo) GetURL(id string) (string, error) {
	if r, ok := (*u).data[id]; ok {
		return string(r[0]), nil
	}
	return "", errors.New("not found")
}

func (u *URLRepo) Get() map[string][]string {
	return u.data
}

func (u *URLRepo) ToSet() *map[string][]string {
	return &u.data
}

// NewUrlRepo return obj with alocate data.
func NewURLRepo() *URLRepo {
	return &URLRepo{
		data: make(map[string][]string),
	}
}

type usersRepo struct {
	data      map[string]int
	currentID int
}

func (u *usersRepo) getID() int {
	u.currentID += 1
	return u.currentID
}

func FindUser(key string) (id int, finded bool) {
	if id, ok := ur.data[key]; ok {
		return id, true
	}
	return -1, false
}

func CreateUser() (string, error) {
	id := ur.getID()
	newKey, err := encription.EncriptInt(id)
	if err != nil {
		return "", err
	}
	ur.data[newKey] = id
	return newKey, nil
}

var ur *usersRepo

func init() {
	ur = &usersRepo{
		data:      make(map[string]int),
		currentID: 0,
	}
}
