package controllers

import (
	"log"
	"net/http"

	"github.com/Godfredasare/go-ecommerce/models"
	"github.com/Godfredasare/go-ecommerce/services"
	"github.com/gin-gonic/gin"
)

func PostProduct(ctx *gin.Context) {
	var product models.Product

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		log.Printf("Error parsing product %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing"})
		return
	}

	err = services.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing"})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{"message": "Product inserted successfully"})
}

func GetAllProducts(ctx *gin.Context) {

}

func GetOneProduct(ctx *gin.Context) {

}

func UpdateProduct(ctx *gin.Context) {

}

func DeleteProduct(ctx *gin.Context) {

}
