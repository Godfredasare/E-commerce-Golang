package router

import (
	"github.com/Godfredasare/go-ecommerce/controllers"
	"github.com/Godfredasare/go-ecommerce/middleware"
	"github.com/gin-gonic/gin"
)

func Router(server *gin.Engine) {
	
	AuthMiddleware := server.Group("/api")
	AuthMiddleware.Use(middleware.AuthMiddleware)

	AuthMiddleware.POST("/product", controllers.PostProduct)
	AuthMiddleware.PUT("/product/:id", controllers.UpdateProduct)
	AuthMiddleware.DELETE("/product/:id", controllers.DeleteProduct)

	
	server.GET("/api/product", controllers.GetAllProducts)
	server.GET("/api/product/:id", controllers.GetOneProduct)

	server.POST("/api/user", controllers.InsertUser)
	server.GET("/api/user", controllers.GetAllUsers)
	server.PUT("/api/user/:id", controllers.UpdateUser)

	server.POST("/api/auth", controllers.LoginUser)


}
