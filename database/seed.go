package database

import (
	"github.com/lionelritchie29/recipitor-be/models"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	db.Where("name = ?", "Apple").FirstOrCreate(&models.Item{
		Name: "Apple",
		Description: "Apple fruit",
		Image: "https://i.ibb.co/bg7wK41/apple-158989157.jpg",
	})

	db.Where("name = ?", "Chicken Breast").FirstOrCreate(&models.Item{
		Name: "Chicken Breast",
		Description: "Breast of chicken",
		Image: "https://i.ibb.co/bRnWYMm/Chicken-Breast-Boneless-3-4-Pieces-Hero-Shot-1.jpg",
	})

	db.Where("name = ?", "Banana").FirstOrCreate(&models.Item{
		Name: "Banana",
		Description: "Banana fruit",
		Image: "https://i.ibb.co/mz8n7Lw/Banana-Single-1.jpg",
	})
}