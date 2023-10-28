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

	db.DB.Limit(20).Offset(offset).Where("is_finished = ?", false).Find(&rooms);

	return c.JSON(http.StatusOK, rooms);
}

func GetRoomById(c echo.Context) error {
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

	id, err := strconv.Atoi(c.Param("id"));
	if err != nil {
		return c.JSON(http.StatusBadRequest, err);
	}
	var room models.Room;

	if err := db.DB.Where("id = ?", id).First(&room).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err);
	}

	return c.JSON(http.StatusOK, room);
}

func MakeFinished(c echo.Context) error {
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

	id, err := strconv.Atoi(c.Param("id"));
	if err != nil {
		return c.JSON(http.StatusBadRequest, err);
	}

	var room models.Room;
	if err := db.DB.Where("id = ?", id).First(&room).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err);
	}

	room.IsFinished = true;
	if err := db.DB.Save(&room).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err);
	}

	return c.JSON(http.StatusOK, room);
}
