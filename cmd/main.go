package main

import (
	"log"
	"onden-backend/api/models"
	"onden-backend/api/router"
	"onden-backend/config"
	"onden-backend/db"
	"onden-backend/services"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := config.GetConfig();
	
	services.HashInit(&config.Hash);
	
	services.AuthInit(&config.JWT);
	
	var err error;
	db.DB, err = db.Connect(config.Database);
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err);
	}

	if err := db.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate the database: %v", err);
	}
	if err := db.DB.AutoMigrate(&models.Room{}); err != nil {
		log.Fatalf("Failed to migrate the database: %v", err);
	}

	e := echo.New();
	
	router.SetupAuthRouter(e);
	router.SetupRoomRouter(e);
	
	e.Use(middleware.Logger());
	
	router.CommonRouter.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWT.SecretKey),
		TokenLookup: "cookie:token", 
	}));

	e.Start(":" + config.Server.Port);
}