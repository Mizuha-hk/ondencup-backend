package config

import (
	"os"
)

type Config struct {
	Server ServerConfig
	Database DatabaseConfig
    Hash HashConfig
    LiveKit LiveKitConfig
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

type LiveKitConfig struct {
    APIKey string
    APISecret string
}

type DatabaseConfig struct {
	DBUserName string
    DBPassword string
    DBHost string
    DBPort string
    DBName string
}

func GetConfig() Config {
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
        LiveKit: LiveKitConfig{
            APIKey: os.Getenv("LIVEKIT_KEY"),
            APISecret: os.Getenv("LIVEKIT_SECRET"),
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
