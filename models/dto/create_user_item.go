package dto

type CreateUserListDto struct {
	Description string `form:"description" json:"description" binding:"required"`
	Items []ListItemDto `form:"items" json:"items" binding:"required"`
}

type ListItemDto struct {
	Id   uint    `form:"id" json:"id" binding:"required"`
	Amount   string `form:"amount" json:"amount" binding:"required"`
	Quantity int    `form:"quantity" json:"quantity" binding:"required"`	
}