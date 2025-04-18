package routes

import (
	"day23Assignment/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/products", controllers.CreateProduct)
	r.GET("/products", controllers.GetAllProducts)
	r.GET("/products/:id", controllers.GetProductByID)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)

	r.GET("/products/:id/image", controllers.GetProductImage)
	r.PUT("/products/:id/image", controllers.UpdateProductImage)

	r.GET("/inventory/:id", controllers.GetStockByID)
	r.PUT("/inventory/:id", controllers.UpdateStock)

	r.GET("/order/:id", controllers.GetOrderByID)
	r.POST("/order/:id", controllers.CreateOrder)

	return r
}
