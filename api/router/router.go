package router

import (
	"onden-backend/api/handler"

	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo){
	api := e.Group("/api");

	api.POST("/login", handler.Login);
	api.POST("/signUp", handler.SignUp);

	api.POST("/user", handler.CreateUser);
	api.GET("/user/:id", handler.GetUserById);
	api.GET("/user", handler.GetAllUsers);
	api.PUT("/user/:id", handler.UpdateUser);
	api.DELETE("/user/:id", handler.DeleteUser);
}