package models

import "context"

type URLs struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

//Repository interface repo urls.
type Repository interface {
	GetURL(context.Context, string) (string, error)
	SaveURL(context.Context, string, string, string) (string, error)
	SaveURLs(context.Context, map[string]string, string, string) (map[string]string, error)
	FindUser(context.Context, string) bool
	CreateUser() (string, error)
	GetUserURLs(string) ([]URLs, error)
	CheckDBConnection() error
	SetURLsToDel([]string, string) error
}
