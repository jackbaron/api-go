package database

import (
	"fmt"

	"github.com/nhatth/api-service/internal/app/helpers"
	userEntity "github.com/nhatth/api-service/internal/app/services/user/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDatabase(cfg helpers.Config) *gorm.DB {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUsername,
		cfg.DBPasssword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	//? Migrate
	db.AutoMigrate(&userEntity.User{})

	return db
}

func GetDBConnection() *gorm.DB {
	return db
}
