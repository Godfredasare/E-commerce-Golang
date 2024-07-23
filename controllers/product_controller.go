package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Godfredasare/go-ecommerce/models"
	"github.com/Godfredasare/go-ecommerce/services"
	"github.com/Godfredasare/go-ecommerce/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PostProduct(ctx *gin.Context) {
	var product models.Product

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		log.Printf("Error parsing product %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing"})
		return
	}

	userID := ctx.GetString("userId")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Unauthorize user"})
		return
	}
	fmt.Println(userID)

	primitiveUserID, _ := primitive.ObjectIDFromHex(userID)
	product.UserId = primitiveUserID

	errMessage := utils.Validation(&product)
	if len(errMessage) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Validation failed", "errors": errMessage})
		return
	}

	err = services.CreateProduct(&product)
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

	errMessage := utils.Validation(&product)
	if len(errMessage) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Validation failed", "errors": errMessage})
		return
	}

	result, err := services.Update(id, &product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error updated products"})
		return
	}

	if result <= 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Product/ID do not exist"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})

}

func DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := services.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error Deteting products"})
		return
	}

	if result <= 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Product/ID do not exist"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})

}
