package config

import "os"

type Config struct {
	DatabaseURL string `env:"DATABASE_URL"`
	Port        string `env:"PORT"`
}

var c Config

func InitConfig() {
	c = Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
	}
}

func GetConfig() Config {
	return c
}
