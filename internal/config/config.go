package config

import (
	"errors"
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
	JWTSecret := os.Getenv("JWT_SECRET")
	if JWTSecret == "" {
		return errors.New("JWT_SECRET IS REQUIRED.")
	}
	if len(JWTSecret) < 32 {
		return errors.New("JWT_SECRET MUST BE AT LEAST 32 CHARACTERS.")
	}

	c.AuthConfig = &AuthConfig{
		Whitelist: whitelist,
		TokenExp:  time.Minute * 15,
		JWTSecret: JWTSecret,
	}
	c.SqlDriver = "sqlite3"
	c.SqlConnString = "mine-server-manager.db"
	return nil
}
