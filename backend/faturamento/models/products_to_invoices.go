package models

type ProductToInvoice struct{
	Id int				`json:"prod_to_inv_id"`
	Invoice_id int		`json:"invoice_id" binding:"required"`
	Product_id string	`json:"product_id" binding:"required"`
	Amount int			`json:"amount" binding:"required"`
}