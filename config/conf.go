package config

import (
	"os"
)

type Config struct {
	DbHost      string
	DbPort      string
	DbUser      string
	DbPassword  string
	DbName      string
	ForexApiKey string
	ForexAPI    string
}

func LoadConfig() Config {
	return Config{
		DbHost:      os.Getenv("DB_HOST"),
		DbPort:      os.Getenv("DB_PORT"),
		DbUser:      os.Getenv("DB_USER"),
		DbPassword:  os.Getenv("DB_PASSWORD"),
		DbName:      os.Getenv("DB_NAME"),
		ForexApiKey: os.Getenv("API_KEY"),
		ForexAPI:    os.Getenv("FOREX_API"),
	}
}
