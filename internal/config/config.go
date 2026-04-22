package config

import "github.com/ilyakaznacheev/cleanenv"

// Config general app Config
type Config struct {
	GroupAPIKey string `yaml:"groupAPIKey"`
	ChatID      string `yaml:"chatID"`
	Logfile     string `yaml:"logfile"`
}

// NewConfig create new Config from file
func NewConfig(filename string) (*Config, error) {

	var c Config

	err := cleanenv.ReadConfig(filename, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
