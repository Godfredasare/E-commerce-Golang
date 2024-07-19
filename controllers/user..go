package controllers

import (
	"net/http"

	"github.com/Godfredasare/go-ecommerce/models"
	"github.com/Godfredasare/go-ecommerce/services"
	"github.com/Godfredasare/go-ecommerce/utils"
	"github.com/gin-gonic/gin"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error inserting user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successully"})

}
