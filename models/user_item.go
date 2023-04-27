package models

import (
	"gorm.io/gorm"
)

type UserItem struct {
	gorm.Model

	ItemId int
	Item Item
	
	UserId int
	User User
	
	Amount string
	Quantity int
}