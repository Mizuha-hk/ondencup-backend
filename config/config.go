package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	DBUserName string
    DBPassword string
    DBHost string
    DBPort string
    DBName string
}

func GetConfig() Config {
    err := godotenv.Load(".env");
    if err != nil {
        // panic("Error loading .env file");
    }

    return Config {
        Server: ServerConfig{
            Port: os.Getenv("SERVER_PORT"),
        },
        Database: DatabaseConfig{
            DBUserName: os.Getenv("DB_USER"),
            DBPassword: os.Getenv("DB_PASSWORD"),
            DBHost: os.Getenv("DB_HOST"),
            DBPort: os.Getenv("DB_PORT"),
            DBName: os.Getenv("DB_NAME"),
        },
    }
}
