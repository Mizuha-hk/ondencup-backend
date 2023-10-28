package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"onden-backend/config"
)

var hashConfig *config.HashConfig;

func HashInit(config *config.HashConfig) {
	hashConfig = config;
}

func HashPassword(password string) (string, error) {
	h := hmac.New(sha256.New, []byte(hashConfig.Key));
	h.Write([]byte(password));
	return hex.EncodeToString(h.Sum(nil)), nil;
}