package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name				string
	Image				string
	Description	string
}