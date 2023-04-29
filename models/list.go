package models

import "gorm.io/gorm"

type List struct {
	gorm.Model
	Description string
	UserId int
}