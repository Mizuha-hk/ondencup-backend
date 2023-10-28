package handler

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetAllRooms(c echo.Context) error {
	user := c.Get("user").(*jwt.Token);
	claims := user.Claims.(jwt.MapClaims);
	userId := claims["user_id"].(string);

	return c.JSON(http.StatusOK, map[string]string{"message":userId});
}