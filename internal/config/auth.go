package config

import (
	"os"
	"strings"
	"time"
)

type AuthConfig struct {
	JWTSecret  string
	Whitelist  []string
	BCryptCost int
	TokenExp   time.Duration
}

func LoadAuthConfig(whitelistPath string) (*AuthConfig, error) {
	var cfg AuthConfig

	cfg.Whitelist = loadWhitelist(whitelistPath)
	cfg.JWTSecret = os.Getenv("JWT_SECRET")
	cfg.BCryptCost = 12
	cfg.TokenExp = time.Hour * 10

	return &cfg, nil
}

func loadWhitelist(whitelistPath string) []string {
	data, err := os.ReadFile(whitelistPath)
	if err != nil {
		// should panic since it's a must
		panic(err)
	}

	whitelist := strings.Split(string(data), "\n")

	return whitelist
}
