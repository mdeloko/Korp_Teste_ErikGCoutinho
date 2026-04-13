package models

type Invoice struct{
	ID int 			`json:"invoice_id"`
	Status string	`json:"status"`
}