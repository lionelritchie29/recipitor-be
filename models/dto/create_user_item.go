package dto

type CreateUserListDto struct {
	Name string `form:"name" json:"name" binding:"required"`
	Items []ListItemDto `form:"items" json:"items" binding:"required"`
}

type ListItemDto struct {
	Id   uint    `form:"id" json:"id" binding:"required"`
	Amount   string `form:"amount" json:"amount" binding:"required"`
	Quantity int    `form:"quantity" json:"quantity" binding:"required"`	
}