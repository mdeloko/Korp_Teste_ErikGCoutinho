package controllers

import (
	"net/http"

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
		ctx.JSON(http.StatusInternalServerError,err)
	}
	ctx.JSON(http.StatusOK,pti)
}