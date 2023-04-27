package models

import (
	"gorm.io/gorm"
)

type List struct {
	gorm.Model

	ItemId int
	Item Item
	
	UserId int
	User User
	
	Amount string
	Quantity int
}