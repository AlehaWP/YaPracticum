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
	servAddr     string `env:"SERVER_ADDRESS"`
	baseURL      string `env:"BASE_URL"`
	repoFileName string `env:"FILE_STORAGE_PATH"`
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
	err := env.Parse(d)
	if err != nil {
		fmt.Println(err.Error())
	}
	if d.servAddr == "" {
		d.servAddr = "localhost:8080"
	}
	if d.baseURL == "" {
		d.baseURL = "http://localhost:8080"
	}
	if d.repoFileName == "" {
		appDir, _ := os.Getwd()
		d.repoFileName = appDir + `\local.gob`
	}
}

func NewdefOptions() Options {
	var o defOptions
	var oEnv defOptions
	oEnv.tryGetFromEnv()
	o = oEnv
	return o
}
