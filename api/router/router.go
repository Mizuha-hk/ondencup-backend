package router

import (
	"onden-backend/api/handler"

	"github.com/labstack/echo/v4"
)

var AuthRouter *echo.Group;
var CommonRouter *echo.Group;

func SetupAuthRouter(e *echo.Echo){
	AuthRouter = e.Group("/auth");

	AuthRouter.POST("/login", handler.Login);
	AuthRouter.POST("/sign-up", handler.SignUp);
}

func SetupRoomRouter(e *echo.Echo){
	CommonRouter = e.Group("/api");
	
	CommonRouter.GET("/room", handler.GetAllRooms);
}