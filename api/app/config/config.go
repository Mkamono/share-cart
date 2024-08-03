package config

import "github.com/kelseyhightower/envconfig" // cspell:disable-line

type Config struct {
	Env    string `required:"true" envconfig:"GO_ENV" default:"local"`
	Port   string `required:"true" envconfig:"PORT" default:"8081"`
	DbUser string `required:"true" envconfig:"POSTGRES_USER"`
	DbPass string `required:"true" envconfig:"POSTGRES_PASSWORD"`
	DbName string `required:"true" envconfig:"POSTGRES_DB"`
	DbHost string `required:"true" envconfig:"POSTGRES_HOST"`
	DbPort string `required:"true" envconfig:"POSTGRES_PORT"`
}

func NewConfig() (Config, error) {
	c := Config{}
	err := envconfig.Process("", &c)
	return c, err
}
