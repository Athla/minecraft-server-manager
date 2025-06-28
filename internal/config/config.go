package config

type Config struct {
	WhitelistPath string
	AuthConfig    *AuthConfig
	SqlDriver     string
	SqlConnString string
}

func NewConfig() (*Config, error) {

	return nil, nil
}
