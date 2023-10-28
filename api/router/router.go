package router

import (
	"onden-backend/api/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var AuthRouter *echo.Group;
var CommonRouter *echo.Group;

func SetupAuthRouter(e *echo.Echo){
	AuthRouter = e.Group("/auth");

	AuthRouter.POST("/login", handler.Login);
	AuthRouter.POST("/sign-up", handler.SignUp);
}

// 引数を増やしてミドルウェアを適応する。
func SetupRoomRouter(e *echo.Echo, configJWT middleware.JWTConfig) *echo.Group {
	CommonRouter := e.Group("/api")
	CommonRouter.Use(middleware.JWTWithConfig(configJWT)) // JWTミドルウェアの設定を適用
	CommonRouter.GET("/room", handler.GetAllRooms)
	return CommonRouter
}
