package models

type ProductToInvoice struct{
	Id int				`json:"prod_to_inv_id"`
	Invoice_id int		`json:"invoice_id"`
	Product_id string	`json:"product_id"`
	Amount int			`json:"amount"`
}