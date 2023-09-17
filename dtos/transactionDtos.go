package dtos

type Transaction struct {
	Id       int     `json:"id" binding:"required"`
	Amount   float64 `json:"amount" binding:"required"`
	Type     string  `json:"type" binding:"required"`
	ParentId int     `json:"parent_id"`
}
