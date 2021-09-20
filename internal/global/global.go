package global

type CtxString string

type Urls struct {
	url    string
	userID int
}

//Options interface for program options.
type Options interface {
	ServAddr() string
	RespBaseURL() string
	RepoFileName() string
}

//Repository interface repo urls.
type Repository interface {
	GetURL(string) (string, error)
	SaveURL([]byte, string) string
	FindUser(string) bool
	CreateUser() (string, error)
}
