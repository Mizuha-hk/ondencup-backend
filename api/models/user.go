package models

import (
	"errors"
	"onden-backend/db"
)

type User struct {
	ID string `gorm:"primaryKey" json:"id"`
	Name string `json:"user_name"`
	Password string `json:"password"`
}

func GetUserByNameAndPassword(name, password string) (*User, error) {
	var user User;

	result := db.DB.Where("name = ? AND password = ?", name, password).First(&user);

	if result.Error != nil {
		return nil, errors.New("invalid credentials");
	}
	return &user, nil;
}