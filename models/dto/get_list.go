package dto

type GetListDto struct {
	Name string
	Items []GetListItemDto
}

type GetListItemDto struct {
	Id uint
	Name string
	Image string
	Amount string
	Quantity int
}