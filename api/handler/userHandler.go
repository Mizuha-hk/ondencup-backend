package handler

import (
	"errors"
	"net/http"
	"onden-backend/api/models"
	"onden-backend/db"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateUser(c echo.Context) error {
	if db.DB == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":"Database is not Initialized"});
	}

	req := new(models.UserReqModel);
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err);
	}

	if req.Name == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message":"Invalid user name or password"});
	}
	
	existingUser := &models.User{};
	if err := db.DB.Where("name = ? AND password = ?", req.Name, req.Password).First(existingUser).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusInternalServerError, err);
		}
	}else{
		return c.JSON(http.StatusConflict, map[string]string{"message":"A user with the given username and password already exists"});
	}

	id, err := uuid.NewRandom();
	if err != nil {
		c.JSON(http.StatusInternalServerError, err);
	}

	user := new(models.User);
	user.ID = id.String();
	user.Name = req.Name;
	user.Password = req.Password;
	if err := db.DB.Create(user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err);
	}

	return c.JSON(http.StatusCreated, user);
}

func GetUserById(c echo.Context) error {
	if db.DB == nil {
		c.JSON(http.StatusInternalServerError,"database is not initialized");
	}

	id := c.Param("id");
	var user models.User;

	if err := db.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, "User not found");
	}

	return c.JSON(http.StatusOK, user);
}

func GetAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, "This endpoint is unhandled but working");
}

func UpdateUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "This endpoint is unhandled but working");
}

func DeleteUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "This endpoint is unhandled but working");
}