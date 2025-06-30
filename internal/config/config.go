package config

import (
	"os"
	"strings"
	"time"
)

type Config struct {
	WhitelistPath string
	AuthConfig    *AuthConfig
	SqlDriver     string
	SqlConnString string
}

func NewConfig() (*Config, error) {
	c := &Config{}
	if err := c.Load(); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Config) Load() error {
	c.WhitelistPath = "whitelist.txt"
	data, err := os.ReadFile(c.WhitelistPath)
	if err != nil {
		// should panic since it's a must
		panic(err)
	}

	whitelist := strings.Split(string(data), "\n")
	c.AuthConfig = &AuthConfig{
		Whitelist: whitelist,
		TokenExp:  time.Minute * 15,
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
	c.SqlDriver = "sqlite3"
	c.SqlConnString = "mine-server-manager.db"
	return nil
}
