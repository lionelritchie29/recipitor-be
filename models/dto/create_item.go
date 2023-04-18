package dto

type CreateItemDto struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Image       string `form:"image" json:"image" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
}