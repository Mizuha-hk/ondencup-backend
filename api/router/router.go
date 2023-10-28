package router

import (
	"onden-backend/api/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupAuthRouter(e *echo.Echo){
	authRouter := e.Group("/auth");

	authRouter.POST("/login", handler.Login);
	authRouter.POST("/sign-up", handler.SignUp);
}

func SetupRoomRouter(e *echo.Echo, configJWT middleware.JWTConfig) *echo.Group {
	commonRouter := e.Group("/api")
	commonRouter.Use(middleware.JWTWithConfig(configJWT)) // JWTミドルウェアの設定を適用
	commonRouter.GET("/room/offset/:offset", handler.GetRooms);
	commonRouter.GET("/room/id/:id", handler.GetRoomById);
	commonRouter.GET("/room/finish/:id", handler.MakeFinished);
	return commonRouter
}
