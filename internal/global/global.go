package global

type CtxString string

type URLs struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

//Options interface for program options.
type Options interface {
	ServAddr() string
	RespBaseURL() string
	RepoFileName() string
	DBConnString() string
}

//Repository interface repo urls.
type Repository interface {
	GetURL(string) (string, error)
	SaveURL(string, string, string) (string, error)
	FindUser(string) bool
	CreateUser() (string, error)
	GetUserURLs(string) ([]URLs, error)
	CheckDBConnection() error
}
