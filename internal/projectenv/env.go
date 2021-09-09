package projectenv

import (
	"os"

	"fmt"

	"github.com/caarlos0/env/v6"
)

type EnvVars struct {
	ServAddr        string `env:"SERVER_ADDRESS"`
	BaseURL         string `env:"BASE_URL"`
	OptionsFileName string `env:"FILE_STORAGE_PATH"`
}

func (e *EnvVars) Get() {
	err := env.Parse(e)
	if err != nil {
		fmt.Println(err.Error())
	}
	if e.ServAddr == "" {
		e.ServAddr = "localhost:8080"
	}
	if e.BaseURL == "" {
		e.BaseURL = "http://localhost:8080"
	}
	if e.OptionsFileName == "" {
		appDir, _ := os.Getwd()
		e.OptionsFileName = appDir + `\local.gob`
	}
}

var Envs EnvVars

func init() {
	Envs.Get()
}
