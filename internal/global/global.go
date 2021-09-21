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
	SaveURL([]byte, string, string) string
	FindUser(string) bool
	CreateUser() (string, error)
	GetUserURLs(string) []URLs
	CheckDBConnection(string) error
}
