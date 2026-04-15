package models

type Product struct{
	Code string			`json:"product_id" binding:"required"`
	Description string	`json:"description" binding:"required"`
	Amount int			`json:"amount" binding:"required"`
}