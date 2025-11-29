package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Token string
}

type Config struct {
	Db   DbConfig
	Auth AuthConfig
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading dotenv file, use default config")
	}

	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Token: os.Getenv("TOKEN"),
		},
	}
}
