package handler

import (
	"onden-backend/api/models"
	"onden-backend/db"
	"onden-backend/services"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetLiveKitToken(c echo.Context) error {
	user := c.Get("user").(*jwt.Token);
	room := c.QueryParam("room");
	claims := user.Claims.(jwt.MapClaims);
	userId := claims["user_id"];
	var userObj *models.User;
	db.DB.Where("id = ?", userId).First(&userObj);

	result, err := services.GetLiveKitToken(userObj.Name, room);
	if err != nil {
		return c.JSON(500, err.Error());
	}
	
	return c.JSON(200, map[string]string{"token": result});
}