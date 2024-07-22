package controllers

import (
	"log"
	"net/http"

	"github.com/Godfredasare/go-ecommerce/services"
	"github.com/Godfredasare/go-ecommerce/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoginUser(ctx *gin.Context) {
	var req services.LoginModel

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("Error parsing product %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing"})
		return
	}

	user, err := services.ValidCredenial(&req)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email/password"})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	} 

	token, err := utils.CreateToken(user.ID.Hex(), user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})

}
