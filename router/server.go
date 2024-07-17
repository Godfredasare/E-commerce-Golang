package router

import (
	"github.com/Godfredasare/go-ecommerce/controllers"
	"github.com/gin-gonic/gin"
)

func Router(server *gin.Engine) {
	server.GET("/api/product", controllers.GetAllProducts)
	server.GET("/api/product/:id", controllers.GetOneProduct)
	server.POST("/api/product", controllers.PostProduct)
	server.PUT("/api/product", controllers.UpdateProduct)
	server.GET("/api/product", controllers.DeleteProduct)

}
