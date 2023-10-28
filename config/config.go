package config

import (
	"os"

	//"github.com/joho/godotenv"
)

type Config struct {
	Server ServerConfig
	Database DatabaseConfig
    Hash HashConfig
    JWT JWTConfig
}

type ServerConfig struct {
	Port string
}

type HashConfig struct {
    Key string
}

type JWTConfig struct {
    SecretKey string
}

type DatabaseConfig struct {
	DBUserName string
    DBPassword string
    DBHost string
    DBPort string
    DBName string
}

func GetConfig() Config {
    // err := godotenv.Load(".env");
    // if err != nil {
    //     panic("Error loading .env file");
    // }

    return Config {
        Server: ServerConfig{
            Port: os.Getenv("SERVER_PORT"),
        },
        Hash: HashConfig{
            Key: os.Getenv("HASH_KEY"),
        },
        JWT: JWTConfig{
            SecretKey: os.Getenv("JWT_SECRET_KEY"),
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
