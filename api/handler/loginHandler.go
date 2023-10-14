package handler

import (
	"log"
	"net/http"
	"onden-backend/api/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
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

	token, err := generateJWTToken(user.ID);
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":"Could not generate Token"});
	}

	return c.JSON(http.StatusOK, map[string]string{"token":token});
}

func generateJWTToken(userID string) (string, error) {

	err := godotenv.Load();
	if err != nil {
		log.Fatal("Error Loading .env");
		return "", err;
	}
	secretKey := os.Getenv("SECRET_KEY");

	expirationTime := time.Now().Add(24 * time.Hour);

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     expirationTime.Unix(),
	})

	return token.SignedString([]byte(secretKey));
}