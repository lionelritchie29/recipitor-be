package models

import "gorm.io/gorm"

type List struct {
	gorm.Model
	Name string
	UserId int
}