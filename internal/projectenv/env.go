package projectenv

import (
	"github.com/caarlos0/env/v6"
)

type EnvVars struct {
	ServAddr string `env:"SERVER_ADDRESS"`
	BaseURL  string `env:"BASE_URL"`
}

func (e *EnvVars) Get() {
	err := env.Parse(e)
	if err != nil || e.ServAddr == "" {
		e.ServAddr = "localhost:8080"
		e.BaseURL = "http://localhost:8080/"
	}
}

var Envs EnvVars

func Init() {
	Envs.Get()
}
