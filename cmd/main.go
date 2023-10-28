package main

import (
	"errors" // 追加
	"fmt"    // 追加
	"log"
	"onden-backend/api/models"
	"onden-backend/api/router"
	"onden-backend/config"
	"onden-backend/db"
	"onden-backend/services"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := config.GetConfig()

	services.HashInit(&config.Hash)

	services.AuthInit(&config.JWT)

	var err error
	db.DB, err = db.Connect(config.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate the database: %v", err)
	}
	if err := db.DB.AutoMigrate(&models.Room{}); err != nil {
		log.Fatalf("Failed to migrate the database: %v", err)
	}

	e := echo.New()
	e.Use(middleware.Logger())

	router.SetupAuthRouter(e)

	configJWT := middleware.JWTConfig{
		SigningKey: []byte(config.JWT.SecretKey),
		ParseTokenFunc: func(tokenString string, c echo.Context) (interface{}, error) {
			keyFunc := func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(config.JWT.SecretKey), nil
			}

			token, err := jwt.Parse(tokenString, keyFunc)
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, errors.New("invalid token")
			}
			return token, nil
		},
	}

	// ミドルウェアの引数を渡してRouterを生成
	router.SetupRoomRouter(e, configJWT)

	e.Start(":" + config.Server.Port)
}
