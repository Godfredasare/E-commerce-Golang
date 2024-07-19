package router

import (
	"github.com/Godfredasare/go-ecommerce/controllers"
	"github.com/gin-gonic/gin"
)

func Router(server *gin.Engine) {

	server.GET("/api/product", controllers.GetAllProducts)
	server.GET("/api/product/:id", controllers.GetOneProduct)
	server.POST("/api/product", controllers.PostProduct)
	server.PUT("/api/product/:id", controllers.UpdateProduct)
	server.DELETE("/api/product/:id", controllers.DeleteProduct)

	server.POST("/api/user", controllers.InsertUser)

}
