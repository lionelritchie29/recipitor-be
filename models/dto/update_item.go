package dto

type UpdateItemDto struct {
	Name        string `form:"name" json:"name" binding:"-"`
	Image       string `form:"image" json:"image" binding:"-"`
	Description string `form:"description" json:"description" binding:"-"`
}