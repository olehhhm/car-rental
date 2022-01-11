package config

import (
	"github.com/joho/godotenv"
)

var conf *Config

func Get() *Config {
	return conf
}

func Init() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}

	// Creating config value
	conf = New()
}
