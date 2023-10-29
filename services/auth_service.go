package services

import (
	"fmt"
	"onden-backend/config"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtConfig *config.JWTConfig;

func AuthInit(config *config.JWTConfig) {
	jwtConfig = config;
}

func GenerateJWTToken(userId string) (string, error){
	secret_key := jwtConfig.SecretKey;
	expirationTime := time.Now().UTC().Add(24 * time.Hour);

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp": expirationTime.Unix(),
	});

	return token.SignedString([]byte(secret_key));
}

func DecodeJWTToken(tokenString string) (*jwt.Token, error){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"]);
        }
        return []byte(jwtConfig.SecretKey), nil
    })

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }

    return token, nil
}