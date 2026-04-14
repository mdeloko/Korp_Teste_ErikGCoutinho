package controllers

import (
	"net/http"
	"strconv"

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
		return
	}
	ctx.JSON(http.StatusOK,invoices)
}

func (ic *InvoiceController) GetInvoice(ctx *gin.Context){
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"Null/Blank id provided!",
		})
		return
	}
	intId,err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"Id can contain symbols or letters!",
		})
		return
	}
	invoice, err := ic.invoiceUseCase.GetInvoice(intId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,err)
		return
	}
	ctx.JSON(http.StatusOK,invoice)
}

func (ic *InvoiceController) CreateInvoice(ctx *gin.Context) {
	id, err := ic.invoiceUseCase.CreateInvoice()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,err)
		return
	}
	ctx.JSON(http.StatusCreated,gin.H{
		"message":"Invoice Created!",
		"invoice_id":id,
	})
}

func (ic *InvoiceController) UpdateInvoiceStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"Null/Blank id provided!",
		})
		return
	}
	intId,err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"Id can contain symbols or letters!",
		})
		return
	}

	var payload struct{
		Status string `json:"newStatus" binding:"required"`
	}
	err = ctx.BindJSON(&payload)
	if err != nil || payload.Status != "opened" && payload.Status != "closed" {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "You must provide a valid 'newStatus' field.",
        })
        return
    }

	invoice, err := ic.invoiceUseCase.UpdateInvoiceStatus(intId,payload.Status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,err)
		return
	}
	ctx.JSON(http.StatusCreated,gin.H{
		"message":"Invoice Updated!",
		"invoice_id":invoice.ID,
		"newStatus":invoice.Status,
	})
}

func (ic *InvoiceController) DeleteInvoice(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"Null/Blank id provided!",
		})
		return
	}
	intId,err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"Id can contain symbols or letters!",
		})
		return
	}

	wasDeleted, err := ic.invoiceUseCase.DeleteInvoice(intId)
	if wasDeleted {
		ctx.JSON(http.StatusOK,gin.H{
		"message":"Invoice Deleted!",
		"invoice_id":intId,
	})
	}else{
		ctx.JSON(http.StatusNotFound,gin.H{
		"message":"Invoice not Deleted!",
		"invoice_id":intId,
	})
	}
}