package db

import (
	"fmt"
	"onden-backend/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB;

func Connect(config config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DBUserName,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{});
}
