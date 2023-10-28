package main

import (
	"log"
	"net/http"
	"onden-backend/api/models"
	"onden-backend/api/router"
	"onden-backend/config"
	"onden-backend/db"

	"github.com/labstack/echo/v4"
)

func main() {
	config := config.GetConfig();
	
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

	router.SetupRouter(e);

	e.GET("/health-check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "healthy");
	});

	e.Start(":" + config.Server.Port);
}
