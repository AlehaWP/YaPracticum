package defoptions

import (
	"flag"
	"fmt"
	"os"

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

type EnvOptions struct {
	ServAddr     string `env:"SERVER_ADDRESS"`
	BaseURL      string `env:"BASE_URL"`
	RepoFileName string `env:"FILE_STORAGE_PATH"`
}

//CheckEnv for get options from env to default application options.
func (d *defOptions) checkEnv() {

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

//CheckFlags for get options from console to default application options.
func (d *defOptions) checkFlags() {
	appDir, _ := os.Getwd()
	a := flag.String("a", "localhost:8080", "a server address string")
	b := flag.String("b", "http://localhost:8080", "a response address string")
	f := flag.String("f", appDir+`\local.gob`, "a file storage path string")
	flag.Parse()
	d.servAddr = *a
	d.baseURL = *b
	d.repoFileName = *f
}

func NewDefOptions() Options {
	opt := new(defOptions)
	opt.checkFlags()
	opt.checkEnv()
	return opt
}
