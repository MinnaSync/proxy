package config

import (
	"sync"

	"github.com/caarlos0/env/v10"
)

type (
	Config struct {
		Port           string `env:"PORT" envDefault:"8080"`
		Environment    string `env:"ENVIRONMENT" envDefault:"production"`
		LogLevel       string `env:"LOG_LEVEL" envDefault:"info"`
		AllowedOrigins string `env:"ALLOWED_ORIGINS"`
	}
)

var (
	once sync.Once

	Conf Config
)

func load() {
	if err := env.Parse(&Conf); err != nil {
		panic(err)
	}
}

func init() {
	once.Do(load)
}
