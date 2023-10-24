package handler

import (
	"net/http"
	"onden-backend/api/models"
	"onden-backend/services"

	"github.com/labstack/echo/v4"
)

func Login (c echo.Context) error {
	req := new(models.UserReqModel);
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message":"Invalid Request"});
	}

	user, err := models.GetUserByNameAndPassword(req.Name, req.Password);
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message":"Invalid Credentials"});
	}

	token, err := services.GenerateJWTToken(user.ID);
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":"Could not generate Token"});
	}

	return c.JSON(http.StatusOK, map[string]string{"token":token});
}