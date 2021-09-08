package projectenv

import (
	"os"

	"github.com/caarlos0/env/v6"
)

type EnvVars struct {
	ServAddr        string `env:"SERVER_ADDRESS"`
	BaseURL         string `env:"BASE_URL"`
	OptionsFileName string `env:"FILE_STORAGE_PATH"`
}

func (e *EnvVars) Get() {
	err := env.Parse(e)
	if err != nil || e.ServAddr == "" {
		e.ServAddr = "localhost:8080"
		e.BaseURL = "http://localhost:8080"
		appDir, _ := os.Getwd()
		e.OptionsFileName = appDir + `\local.db`
	}
}

var Envs EnvVars

func init() {
	Envs.Get()
}
