package controllers

import (
	"net/http"

	"github.com/Godfredasare/go-ecommerce/models"
	"github.com/Godfredasare/go-ecommerce/services"
	"github.com/Godfredasare/go-ecommerce/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertUser(ctx *gin.Context) {
	var user models.Users

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing user"})
		return
	}

	errMessage := utils.Validation(user)
	if len(errMessage) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Validation error", "errors": errMessage})
		return
	}

	err = services.CreateUser(&user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error inserting user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successully"})

}

func GetAllUsers(ctx *gin.Context) {
	users, err := services.FindAllUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error getting users"})
		return
	}

	ctx.JSON(http.StatusFound, users)
}

func UpdateUser(ctx *gin.Context) {
	var user models.Users
	id := ctx.Param("id")

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing user"})
		return
	}

	errMessage := utils.Validation(user)
	if len(errMessage) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Validation error", "errors": errMessage})
		return
	}

	result, err := services.UpdateUser(id, &user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error inserting user"})
		return
	}

	if result <= 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "user/id not found"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User updated successully"})

}
