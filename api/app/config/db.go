package config

import "fmt"

type DB struct {
	DbUser string `required:"true" envconfig:"POSTGRES_USER"`
	DbPass string `required:"true" envconfig:"POSTGRES_PASSWORD"`
	DbName string `required:"true" envconfig:"POSTGRES_DB"`
	DbHost string `required:"true" envconfig:"POSTGRES_HOST"`
	DbPort string `required:"true" envconfig:"POSTGRES_PORT"`
}

func (c *Config) GetDBConnStr() string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		c.DB.DbUser,
		c.DB.DbPass,
		c.DB.DbName,
		c.DB.DbHost,
		c.DB.DbPort,
	)
}
