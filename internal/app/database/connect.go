package database

import (
	"fmt"
	"log"

	"github.com/nhatth/api-service/pkg/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg utils.Config) *gorm.DB {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUsername,
		cfg.DBPasssword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	log.Println(dns)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	defer func() {
		dbIntansce, _ := db.DB()
		_ = dbIntansce.Close()
	}()

	return db
}
