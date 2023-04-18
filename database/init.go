package database

import (
	"os"

	"github.com/lionelritchie29/recipitor-be/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	dsn := os.Getenv("CONNECTION_STRING")
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&models.Item{})

	return db, err
}