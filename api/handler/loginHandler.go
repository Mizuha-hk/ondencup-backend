package handler

import (
	"net/http"
	"onden-backend/api/models"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	// シークレットキーを設定
	secretKey := "your_secret_key" // 実際のアプリケーションでは、安全な方法でキーを管理してください
	// トークンの有効期限を設定
	expirationTime := time.Now().Add(24 * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     expirationTime.Unix(),
	})

	return token.SignedString([]byte(secretKey))
}