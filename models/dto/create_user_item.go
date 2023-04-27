package dto

type CreateUserItemDto struct {
	ItemId   int    `form:"itemId" json:"itemId" binding:"required"`
	Amount   string `form:"amount" json:"amount" binding:"required"`
	Quantity int    `form:"quantity" json:"quantity" binding:"required"`
}