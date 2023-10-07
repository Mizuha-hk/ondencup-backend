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

func SignUp(c echo.Context) error {
	req := new(models.UserReqModel);
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message":"Invalid Request"});
	}

	if req.Name == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message":"User_name or Password is empty"});
	}

	id, err := uuid.NewRandom();
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":"Failed to generate uuid"});
	}

	if db.DB == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":"Database is not initialized"});
	}

	existingUser := &models.User{};
	if err := db.DB.Where("name = ? AND password = ?", req.Name, req.Password).First(existingUser).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound){
			return c.JSON(http.StatusInternalServerError, err);
		}
	} else {
		return c.JSON(http.StatusConflict, map[string]string{"error": "A user with the given username and password already exists"})
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