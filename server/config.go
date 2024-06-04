package main

import (
	"errors"
	"server/shared"

	"github.com/spf13/viper"
)

type Config struct {
	Postgres PostgresConfig
}

// Postgres config
type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
	Driver   string
}

// Load config file from given path
func LoadConfig(filePathNoExtension string) (*viper.Viper, error) {
	v := viper.New()

	v.AddConfigPath(filePathNoExtension)
	v.SetConfigName("config")
	v.SetConfigType("yml")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// Parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	shared.ErrorHandling(err)

	return &c, nil
}

// Get config
func GetConfig(filePathNoExtension string) (*Config, error) {
	cfgFile, err := LoadConfig(filePathNoExtension)
	shared.ErrorHandling(err)

	cfg, err := ParseConfig(cfgFile)
	shared.ErrorHandling(err)
	return cfg, nil
}
