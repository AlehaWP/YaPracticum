package defoptions

import (
	"os"

	"fmt"

	"github.com/caarlos0/env/v6"
)

type Options interface {
	ServAddr() string
	RespBaseURL() string
	RepoFileName() string
}

type defOptions struct {
	servAddr     string
	baseURL      string
	repoFileName string
}

func (d defOptions) ServAddr() string {
	return d.servAddr
}

func (d defOptions) RespBaseURL() string {
	return d.baseURL
}

func (d defOptions) RepoFileName() string {
	return d.repoFileName
}

func (d *defOptions) tryGetFromEnv() {
	type EnvOptions struct {
		ServAddr     string `env:"SERVER_ADDRESS"`
		BaseURL      string `env:"BASE_URL"`
		RepoFileName string `env:"FILE_STORAGE_PATH"`
	}
	e := &EnvOptions{}
	err := env.Parse(e)
	if err != nil {
		fmt.Println(err.Error())
	}
	if len(e.ServAddr) != 0 {
		d.servAddr = e.ServAddr
	}
	if len(e.BaseURL) != 0 {
		d.baseURL = e.BaseURL
	}
	if len(e.RepoFileName) != 0 {
		d.repoFileName = e.RepoFileName
	}
}

func NewdefOptions() Options {
	appDir, _ := os.Getwd()
	opt := defOptions{
		servAddr:     "localhost:8080",
		baseURL:      "http:localhost:8080",
		repoFileName: appDir + `\local.gob`,
	}
	opt.tryGetFromEnv()
	fmt.Println(opt.ServAddr())
	return opt
}
