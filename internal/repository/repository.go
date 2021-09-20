package repository

import (
	"errors"

	encription "github.com/AlehaWP/YaPracticum.git/internal/Encription"
	"github.com/AlehaWP/YaPracticum.git/internal/global"
	"github.com/AlehaWP/YaPracticum.git/internal/shorter"

	"github.com/AlehaWP/YaPracticum.git/internal/serialize"
)

var serializeURLRepo func(global.Repository)

//UrlsData repository of urls. Realize Repository interface.
type ServerRepo struct {
	URLsData map[string][]string
	Users    UsersRepo
}

type UsersRepo struct {
	Data      map[string]int
	CurrentID int
}

func (s *ServerRepo) SaveURL(url []byte, baseURL, userID string) string {
	r := shorter.MakeShortner(url)
	(*s).URLsData[r] = []string{string(url), baseURL, userID}
	serializeURLRepo(s)
	return baseURL + r
}

func (s *ServerRepo) GetURL(id string) (string, error) {
	if r, ok := (*s).URLsData[id]; ok {
		return string(r[0]), nil
	}
	return "", errors.New("not found")
}

func (s *ServerRepo) GetUserURLs(userID string) []global.URLs {
	ud := s.URLsData
	m := make([]global.URLs, 0, 0)
	for key, value := range ud {
		if value[2] == userID {
			m = append(m, global.URLs{
				ShortURL:    value[1] + key,
				OriginalURL: value[0],
			})
		}
	}

	return m
}

func (u *UsersRepo) getNewID() int {
	u.CurrentID += 1
	return u.CurrentID
}

func (s *ServerRepo) FindUser(key string) (finded bool) {
	ur := s.Users
	if _, ok := ur.Data[key]; ok {
		return true
	}
	return false
}

func (s *ServerRepo) CreateUser() (string, error) {
	ur := &s.Users
	id := ur.getNewID()
	newKey, err := encription.EncriptInt(id)
	if err != nil {
		return "", err
	}
	ur.Data[newKey] = id
	serializeURLRepo(s)
	return newKey, nil
}

// NewUrlRepo return obj with alocate data.
func NewRepo(repoFileName string) *ServerRepo {
	servRepo := &ServerRepo{
		URLsData: make(map[string][]string),
		Users: UsersRepo{
			Data:      make(map[string]int),
			CurrentID: 0,
		},
	}
	serialize.NewSerialize(repoFileName)
	serialize.ReadRepoFromFile(servRepo)
	// fmt.Println(servRepo)
	serializeURLRepo = serialize.SaveRepoToFile

	return servRepo
}
