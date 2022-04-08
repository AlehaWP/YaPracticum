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
	config       string
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
	ServAddr     string `env:"SERVER_ADDRESS" json:"server_address"`
	BaseURL      string `env:"BASE_URL" json:"base_url"`
	RepoFileName string `env:"FILE_STORAGE_PATH" json:"file_storage_dsn"`
	DBConnString string `env:"DATABASE_DSN" json:"database_dsn"`
	Config       string `env:"CONFIG" json:"-"`
	EnableHTTPS  bool   `env:"ENABLE_HTTPS" json:"enable_https"`
}

func (d *defOptions) fillFromConf(c *Config) {
	if len(c.ServAddr) != 0 && len(d.servAddr) == 0 {
		d.servAddr = c.ServAddr
	}
	if len(c.BaseURL) != 0 && len(d.baseURL) == 0 {
		d.baseURL = c.BaseURL
	}
	if len(c.RepoFileName) != 0 && len(d.repoFileName) == 0 {
		d.repoFileName = c.RepoFileName
	}
	if len(c.DBConnString) != 0 && len(d.dbConnString) == 0 {
		d.dbConnString = c.DBConnString
	}
	if c.EnableHTTPS == true {
		d.enableHTTPS = true
	}
	if len(c.Config) != 0 && len(d.config) == 0 {
		d.config = c.Config
	}
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
	d.fillFromConf(config)
}

func (d *defOptions) setDefault(appDir string) {

	config := &Config{
		"localhost:80",
		"http://localhost:8080",
		appDir + `/local.gob`,
		"user=kseikseich dbname=yap sslmode=disable",
		"config.json",
		false,
	}
	d.fillFromConf(config)
}

//checkEnv for get options from env to default application options.
func (d *defOptions) checkEnv() {

	e := &Config{}
	err := env.Parse(e)
	if err != nil {
		fmt.Println(err.Error())
	}
	d.fillFromConf(e)

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
	flag.StringVar(&d.config, "c", d.config, "a config file name")

	flag.Parse()

}

// NewDefOptions return obj like Options interfasc.
func NewDefOptions() models.Options {
	appDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Не удалось найти каталог программы!")
	}

	opt := &defOptions{}

	opt.setFlags()
	opt.checkEnv()
	opt.setDefault(appDir)

	if len(opt.config) == 0 {
		opt.config = "config.json"
	}

	f := appDir + string(os.PathSeparator) + opt.config
	if ok, _ := exists(f); ok {
		opt.readConfig(f)
	}
	opt.saveConfiguration(f)

	return opt
}
