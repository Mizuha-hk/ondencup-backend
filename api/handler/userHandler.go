package handler

import (
	"onden-backend/api/models"
	"onden-backend/db"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetUserName(c echo.Context) error {
	user := c.Get("user").(*jwt.Token);
	claims := user.Claims.(jwt.MapClaims);
	userId := claims["user_id"];
	var userObj models.User;
	db.DB.Where("id = ?", userId).First(&userObj)

	return c.JSON(200, userObj.Name);
}