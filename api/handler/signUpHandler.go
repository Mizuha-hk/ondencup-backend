package handler

import (
	"errors"
	"net/http"
	"onden-backend/api/models"
	"onden-backend/db"
	"onden-backend/services"

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

	if _, err := models.GetUserByNameAndPassword(req.Name, req.Password); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound){
			return c.JSON(http.StatusInternalServerError, err);
		}		
	} else {
		return c.JSON(http.StatusConflict, map[string]string{"error": "A user with the given username and password already exists"})
	}

	user := new(models.User);
	user.ID = id.String();
	user.Name = req.Name;

	token, err := services.GenerateJWTToken(user.ID);
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":"Could not generate Token"});
	}

	hashedPassword, err := services.HashPassword(req.Password);
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err);
	}

	user.Password = hashedPassword;
	if err := db.DB.Create(user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err);
	}

	return c.JSON(http.StatusCreated, map[string]string{"token":token});
}