package controllers

import (
	"fmt"
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

	ctx.JSON(http.StatusCreated, gin.H{"message": "Product inserted successfully"})
}

func GetAllProducts(ctx *gin.Context) {

	products, err := services.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error Getting all products"})
		return
	}

	ctx.JSON(http.StatusFound, products)

}

func GetOneProduct(ctx *gin.Context) {

	id := ctx.Param("id")

	products, err := services.FindOne(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error Getting One products"})
		return
	}

	ctx.JSON(http.StatusFound, products)
}

func UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	var product models.Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		log.Printf("Error parsing product %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing"})
		return
	}

	result, err := services.Update(id, &product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error updated products"})
		return
	}

	fmt.Println(result)

	ctx.JSON(http.StatusCreated, gin.H{"message": "Product updated successfully"})

}

func DeleteProduct(ctx *gin.Context) {

}
