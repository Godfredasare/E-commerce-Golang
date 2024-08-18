package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Godfredasare/go-ecommerce/models"
	"github.com/Godfredasare/go-ecommerce/services"
	"github.com/Godfredasare/go-ecommerce/utils"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PostProduct(ctx *gin.Context) {
	var product models.Product

	product.Name = ctx.PostForm("name")
	product.Description = ctx.PostForm("description")
	product.Price, _ = strconv.ParseFloat((ctx.PostForm("price")), 64)
	product.Currency = ctx.PostForm("currency")
	product.Stock, _ = strconv.ParseInt((ctx.PostForm("stock")), 10, 64)
	product.Category = ctx.PostForm("category")

	errMessage := utils.Validation(&product)
	if len(errMessage) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Validation failed", "errors": errMessage})
		return
	}

	// Get user ID from context
	userID := ctx.GetString("userId")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Unauthorized user"})
		return
	}

	// Convert userID to primitive.ObjectID
	primitiveUserID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Printf("Error converting userID: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
		return
	}
	product.UserId = primitiveUserID

	// Handle file upload
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files := form.File["images"]

	// Save file locally

	var imagesID []string

	for _, file := range files {

		// Upload the file to specific dst.
		err = ctx.SaveUploadedFile(file, "assets/uploads/"+file.Filename)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
			return
		}

		imageID, _, err := utils.UploadToCloudinary(file)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		imagesID = append(imagesID, imageID)

	}

	product.ImagesID = imagesID

	// Save product to database
	err = services.CreateProduct(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error inserting product"})
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

	fmt.Println(products)
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

	var req models.Product

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("Error parsing product %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Errhhor parsing"})
		return
	}

	errMessage := utils.Validation(&req)
	if len(errMessage) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Validation failed", "errors": errMessage})
		return
	}

	//check if user_id in db matches with the user updating from the middleware
	product, err := services.FindOne(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Product do not exist"})
		return
	}

	userID := ctx.GetString("userId")
	primitiveUserID, _ := primitive.ObjectIDFromHex(userID)

	if product.UserId != primitiveUserID {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Unauthorize user"})
		return
	}

	//update product
	result, err := services.Update(id, &req)
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

	product, err := services.FindOne(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Product do not exist"})
		return
	}

	userID := ctx.GetString("userId")
	primitiveUserID, _ := primitive.ObjectIDFromHex(userID)

	if product.UserId != primitiveUserID {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Unauthorize user"})
		return
	}

	//delete images from cloudinary
	for _, imageId := range product.ImagesID {
		err := utils.DeleteFromCloudinary(imageId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
	}

	//Detele product
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

func SearchProduct(ctx *gin.Context) {
	searchQuery := ctx.Query("search")

	products, err := services.SearchProduct(searchQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusFound, products)
}

func GetProductByUserId(ctx *gin.Context) {
	id := ctx.Param("id")

	products, err := services.FindProductsByUser(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusFound, products)

}
