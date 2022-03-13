package defoptions

import (
	"flag"
	"fmt"
	"os"

	"github.com/AlehaWP/YaPracticum.git/internal/models"
	"github.com/caarlos0/env/v6"
)

type defOptions struct {
	servAddr     string
	baseURL      string
	repoFileName string
	dbConnString string
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

func (d defOptions) DBConnString() string {
	return d.dbConnString
}

type EnvOptions struct {
	ServAddr     string `env:"SERVER_ADDRESS"`
	BaseURL      string `env:"BASE_URL"`
	RepoFileName string `env:"FILE_STORAGE_PATH"`
	DBConnString string `env:"DATABASE_DSN"`
}

//checkEnv for get options from env to default application options.
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
	if len(e.DBConnString) != 0 {
		d.dbConnString = e.DBConnString
	}
}

//setFlags for get options from console to default application options.
func (d *defOptions) setFlags() {

	flag.StringVar(&d.servAddr, "a", d.servAddr, "a server address string")
	flag.StringVar(&d.baseURL, "b", d.baseURL, "a response address string")
	flag.StringVar(&d.repoFileName, "f", d.repoFileName, "a file storage path string")
	flag.StringVar(&d.dbConnString, "d", d.dbConnString, "a db connection string")

	flag.Parse()

}

// NewDefOptions return obj like Options interfase.
func NewDefOptions() models.Options {
	appDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Не удалось найти каталог программы!")
	}
	opt := &defOptions{
		"localhost:8080",
		"http://localhost:8080",
		appDir + `/local.gob`,
		"user=kseikseich dbname=yap sslmode=disable",
	}

	opt.checkEnv()
	opt.setFlags()

	fmt.Println(opt)

	return opt
}
