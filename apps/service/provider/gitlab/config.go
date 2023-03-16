package gitlab

import (
	"net/http"

	"github.com/caarlos0/env/v6"
	"github.com/infraboard/mcenter/common/validate"
)

func NewConfigFromHTTP(r *http.Request) *Config {
	conf := NewDefaultConfig()

	qs := r.URL.Query()
	addr := qs.Get("address")
	if addr != "" {
		conf.Address = addr
	}
	conf.PrivateToken = r.Header.Get("GITLAB_PRIVATE_TOKEN")

	return conf
}

func NewDefaultConfig() *Config {
	return &Config{
		Address: "https://gitlab.com",
	}
}

func LoadConfigFromEnv() (*Config, error) {
	conf := NewDefaultConfig()
	err := env.Parse(conf)
	if err != nil {
		return nil, err
	}
	if err := conf.Validate(); err != nil {
		return nil, err
	}
	return conf, nil
}

type Config struct {
	Address      string `json:"address" env:"GITLAB_ADDRESS" validate:"required"`
	PrivateToken string `json:"private_token" env:"GITLAB_PRIVATE_TOKEN" validate:"required"`
}

func (c *Config) Validate() error {
	return validate.Validate(c)
}

func (c *Config) SetAddress(address string) {
	if address != "" {
		c.Address = address
	}
}

func (c *Config) SetPrivateToken(tk string) {
	if tk != "" {
		c.PrivateToken = tk
	}
}
