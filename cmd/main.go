package main

import (
	"errors" // 追加
	"fmt"    // 追加
	"log"
	"net/http"
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

	services.HashInit(&config.Hash);

	services.AuthInit(&config.JWT);

	services.LiveKitInit(&config.LiveKit);

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
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

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

	router.SetupWebSocketRouter(e , configJWT);
	router.SetupRoomRouter(e, configJWT)
	e.GET("/health-check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "healthy");
	});

	e.Start(":" + config.Server.Port);
}
