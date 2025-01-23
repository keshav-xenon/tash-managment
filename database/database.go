package database

import (
	"log"

	"task-management/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
