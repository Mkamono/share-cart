package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Env  string `required:"true" envconfig:"GO_ENV" default:"local"`
	Port string `required:"true" envconfig:"PORT" default:"8081"`
}

func NewConfig() (Config, error) {
	c := Config{}
	err := envconfig.Process("", &c)
	return c, err
}
