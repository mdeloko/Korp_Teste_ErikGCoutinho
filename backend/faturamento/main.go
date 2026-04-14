package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mdeloko/Korp_Teste_ErikGCoutinho/controllers"
	"github.com/mdeloko/Korp_Teste_ErikGCoutinho/repositories"
	usecases "github.com/mdeloko/Korp_Teste_ErikGCoutinho/useCases"
)

func main(){
	server := gin.Default()

	dbConn, err := repositories.ConnectDB()
	if err != nil {
		panic(err)
	}

	InvoiceRepository := repositories.NewInvoiceRepository(dbConn)
	InvoiceUseCase := usecases.NewInvoiceUseCase(InvoiceRepository)
	InvoiceController := controllers.NewInvoiceController(InvoiceUseCase)

	ProductsToInvoiceRepository := repositories.NewProductsToInvoicesRepository(dbConn)
	ProductsToInvoiceUseCase := usecases.NewProductsToInvoiceUseCase(ProductsToInvoiceRepository)
	ProductsToInvoiceController := controllers.NewProductsToInvoiceController(ProductsToInvoiceUseCase)


	server.GET("/ping",func(ctx *gin.Context){
		ctx.JSON(200,gin.H{
			"message":"pong!",
			"time":time.Now().Local(),
		})
	})
	
	server.GET("/invoices",InvoiceController.GetInvoices)
	server.GET("/invoice/:id",InvoiceController.GetInvoice)
	server.POST("/invoice",InvoiceController.CreateInvoice)
	server.PATCH("/invoice/status/:id",InvoiceController.UpdateInvoiceStatus)
	server.DELETE("/invoice/:id",InvoiceController.DeleteInvoice)

	server.GET("/products-to-invoice",ProductsToInvoiceController.GetProductsToInvoice)
	server.POST("/products-to-invoice",ProductsToInvoiceController.AddProductToInvoice)
	server.DELETE("/products-to-invoice/:id",ProductsToInvoiceController.RemoveProductFromInvoice)

	server.Run(":5001")
	
}