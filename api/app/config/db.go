package config

type DB struct {
	DbUser string `required:"true" envconfig:"POSTGRES_USER"`
	DbPass string `required:"true" envconfig:"POSTGRES_PASSWORD"`
	DbName string `required:"true" envconfig:"POSTGRES_DB"`
	DbHost string `required:"true" envconfig:"POSTGRES_HOST"`
	DbPort string `required:"true" envconfig:"POSTGRES_PORT"`
}
