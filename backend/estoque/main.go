package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/mdeloko/Korp_Teste_ErikGCoutinho/controllers"
	"github.com/mdeloko/Korp_Teste_ErikGCoutinho/repositories"
	usecases "github.com/mdeloko/Korp_Teste_ErikGCoutinho/useCases"
)

func main(){
	corsConfig := cors.Config{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowMethods: []string{"GET","POST","PATCH","DELETE"},
		AllowHeaders: []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"},
	}
	server := gin.Default()
	server.Use(cors.New(corsConfig))

	dbConn, err := repositories.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductRepository := repositories.NewProductRepository(dbConn)
	ProductUseCase := usecases.NewProductUseCase(ProductRepository)
	ProductController := controllers.NewProductController(ProductUseCase)

	server.GET("/ping",func(ctx *gin.Context){
		ctx.JSON(200,gin.H{
			"message":"pong!",
			"time":time.Now().Local(),
		})
	})
	server.GET("/products",ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.PATCH("/product/amount/:id", ProductController.DecrementOrIncrementProduct)
	server.PATCH("/product/rename/:id", ProductController.RenameProduct)

	server.Run(":5000")
	
}