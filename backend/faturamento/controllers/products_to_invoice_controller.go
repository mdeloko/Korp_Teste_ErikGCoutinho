package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	usecases "github.com/mdeloko/Korp_Teste_ErikGCoutinho/useCases"
)

type ProductsToInvoiceController struct {
	productsToInvoiceUseCase usecases.ProductsToInvoiceUseCases
}

func NewProductsToInvoiceController (useCase usecases.ProductsToInvoiceUseCases) ProductsToInvoiceController {
	return ProductsToInvoiceController{
		productsToInvoiceUseCase: useCase,
	}
}

func (ptic *ProductsToInvoiceController) GetProductsToInvoice(ctx *gin.Context){
	pti, err := ptic.productsToInvoiceUseCase.GetProductsToInvoices()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,pti)
}

func (ptic *ProductsToInvoiceController) AddProductToInvoice(ctx *gin.Context) {
	var payload struct{
		Invoice_id int    `json:"invoice_id" binding:"required"`
    	Product_id string `json:"product_id" binding:"required"`
    	Amount     int    `json:"amount" binding:"required"`
	}
	err := ctx.BindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Malformed payload or missing information!",
		})
		return
	}
	
	err = ptic.productsToInvoiceUseCase.AddProductToInvoice(payload.Invoice_id,payload.Product_id,payload.Amount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Product linked to invoice with success!"})
}

func (ptic *ProductsToInvoiceController) RemoveProductFromInvoice(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "You must provide an actual link id!",
		})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id format"})
		return
	}

	err = ptic.productsToInvoiceUseCase.RemoveProductFromInvoice(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product removed from invoice successfully!"})
}
