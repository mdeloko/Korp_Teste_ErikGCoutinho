package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mdeloko/Korp_Teste_ErikGCoutinho/controllers"
	"github.com/mdeloko/Korp_Teste_ErikGCoutinho/repositories"
	usecases "github.com/mdeloko/Korp_Teste_ErikGCoutinho/useCases"
)

func main() {
	corsConfig := cors.Config{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowMethods: []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders: []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"},
	}
	server := gin.Default()
	server.Use(cors.New(corsConfig))

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

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong!",
			"time":    time.Now().Format(time.RFC1123),
		})
	})

	server.GET("/invoices", InvoiceController.GetInvoices)
	server.GET("/invoice/:id", InvoiceController.GetInvoice)
	server.POST("/invoice", InvoiceController.CreateInvoice)
	server.PATCH("/invoice/status/:id", InvoiceController.UpdateInvoiceStatus)
	server.DELETE("/invoice/:id", InvoiceController.DeleteInvoice)

	server.GET("/products-to-invoice", ProductsToInvoiceController.GetProductsToInvoice)
	server.POST("/products-to-invoice", ProductsToInvoiceController.AddProductToInvoice)
	server.DELETE("/products-to-invoice/:invoiceId/:productId", ProductsToInvoiceController.RemoveProductFromInvoice)
	server.GET("/products-to-invoice/invoice/:id", ProductsToInvoiceController.GetProductsToInvoiceByInvoiceId)

	server.Run(":5001")

}
