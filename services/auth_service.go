package services

import (
	"onden-backend/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtConfig *config.JWTConfig;

func AuthInit(config *config.JWTConfig) {
	jwtConfig = config;
}

func GenerateJWTToken(userId string) (string, error){
	secret_key := jwtConfig.SecretKey;
	expirationTime := time.Now().Add(24 * time.Hour);

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"user_id": userId,
		"exp": expirationTime.Unix(),
	});

	return token.SignedString([]byte(secret_key));
}