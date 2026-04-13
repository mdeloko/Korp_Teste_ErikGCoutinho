package models

type Product struct{
	Code string			`json:"product_id"`
	Description string	`json:"description"`
	Amount int			`json:"amount"`
}