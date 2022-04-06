package defoptions

import (
	"encoding/json"
	"errors"
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
	enableHTTPS  bool
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

func (d defOptions) HTTPS() bool {
	return d.enableHTTPS
}

type Config struct {
	ServAddr     string `env:"SERVER_ADDRESS"`
	BaseURL      string `env:"BASE_URL"`
	RepoFileName string `env:"FILE_STORAGE_PATH"`
	DBConnString string `env:"DATABASE_DSN"`
	EnableHTTPS  bool   `env:"ENABLE_HTTPS"`
}

//checkEnv for get options from env to default application options.
func (d *defOptions) checkEnv() {

	e := &Config{}
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
	if e.EnableHTTPS == true {
		d.enableHTTPS = true
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (d *defOptions) readConfig(file string) {

	config := &Config{}

	configFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(config)
	if err != nil {
		fmt.Println(err.Error())
	}
	if len(config.ServAddr) != 0 {
		d.servAddr = config.ServAddr
	}
	if len(config.BaseURL) != 0 {
		d.baseURL = config.BaseURL
	}
	if len(config.RepoFileName) != 0 {
		d.repoFileName = config.RepoFileName
	}
	if len(config.DBConnString) != 0 {
		d.dbConnString = config.DBConnString
	}
	if config.EnableHTTPS == true {
		d.enableHTTPS = true
	}
}

func (d *defOptions) saveConfiguration(file string) error {
	config := &Config{
		ServAddr:     d.servAddr,
		BaseURL:      d.baseURL,
		RepoFileName: d.repoFileName,
		DBConnString: d.dbConnString,
		EnableHTTPS:  d.enableHTTPS,
	}
	configFile, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return errors.New("не удалось найти файл " + file)
	}
	jsonParser := json.NewEncoder(configFile)
	jsonParser.Encode(&config)
	return nil
}

//setFlags for get options from console to default application options.
func (d *defOptions) setFlags() {

	flag.StringVar(&d.servAddr, "a", d.servAddr, "a server address string")
	flag.StringVar(&d.baseURL, "b", d.baseURL, "a response address string")
	flag.StringVar(&d.repoFileName, "f", d.repoFileName, "a file storage path string")
	flag.StringVar(&d.dbConnString, "d", d.dbConnString, "a db connection string")
	flag.BoolVar(&d.enableHTTPS, "s", d.enableHTTPS, "enable https connection")

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
		false,
	}

	f := appDir + `/config.json`
	// if ok, _ := exists(f); ok {
	// 	opt.readConfig(f)
	// }

	opt.checkEnv()
	opt.setFlags()
	opt.saveConfiguration(f)

	return opt
}
