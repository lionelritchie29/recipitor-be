package models

import (
	"gorm.io/gorm"
)

type ListItem struct {
	gorm.Model

	ItemId uint
	Item Item
	
	ListId uint
	List List
	
	Amount string
	Quantity int
}