package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/suhaibkhan/apitestgo/internal/domain"
)

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=password dbname=apitestdb port=5432 sslmode=disable TimeZone=Asia/Dubai"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetMaxOpenConns(10)

	db.AutoMigrate(&domain.Todo{})

	return db
}
