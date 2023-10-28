package handler

import (
	"net/http"
	"onden-backend/api/models"
	"onden-backend/db"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetRooms(c echo.Context) error {

	user := c.Get("user")
	if user == nil {
		return c.JSON(http.StatusUnauthorized, "token is missing")
	}
		
	token := user.(*jwt.Token)
		
	claims := token.Claims.(jwt.MapClaims);
	userId := claims["user_id"].(string);
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, "invalid user")
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"));
	if err != nil {
		offset = 0;
	}

	var rooms []models.Room;

	db.DB.Limit(20).Offset(offset).Where("is_finished = ?", "FALSE").Find(&rooms);

	return c.JSON(http.StatusOK, rooms);
}