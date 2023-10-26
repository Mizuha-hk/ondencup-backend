package models

import (
	"onden-backend/db"
	"onden-backend/services"
)

type User struct {
	ID string `gorm:"primaryKey" json:"id"`
	Name string `json:"user_name"`
	Password string `json:"password"`
}

func GetUserByNameAndPassword(name, password string) (*User, error) {
	var user User;

	hashedPassword, err := services.HashPassword(password);
	if(err != nil) {
		return nil, err;
	}

	result := db.DB.Where("name = ? AND password = ?", name, hashedPassword).First(&user);

	if result.Error != nil {
		return nil, result.Error;
	}
	return &user, nil;
}