package db

import (
	"fmt"
	"kienmatu/go-todos/config"
	"kienmatu/go-todos/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostgresInstance(cfg *config.Configuration, migrate bool) *gorm.DB {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		"postgres",
		"duynghia123",
		"todos",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if migrate {
		db.AutoMigrate(&models.User{}, &models.Todo{})
	}
	return db
}
