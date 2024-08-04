package config

type AUTH0 struct {
	Domain   string `required:"true" envconfig:"AUTH0_DOMAIN"`
	Audience string `required:"true" envconfig:"AUTH0_AUDIENCE"`
}
