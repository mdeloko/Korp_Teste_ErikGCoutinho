package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecases "github.com/mdeloko/Korp_Teste_ErikGCoutinho/useCases"
)

type InvoiceController struct{
	invoiceUseCase usecases.InvoiceUseCase
}

func NewInvoiceController(useCase usecases.InvoiceUseCase) InvoiceController{
	return InvoiceController{
		invoiceUseCase: useCase,
	}
}

func (ic *InvoiceController) GetInvoices(ctx *gin.Context){
	invoices, err := ic.invoiceUseCase.GetInvoices()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,err)
	}
	ctx.JSON(http.StatusOK,invoices)
}