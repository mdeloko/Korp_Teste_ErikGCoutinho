package models

type Invoice struct{
	ID int 			`json:"invoice_id" binding:"required"`
	Status string	`json:"status" binding:"required"`
}