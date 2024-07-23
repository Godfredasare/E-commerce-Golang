package middleware

import (
	"net/http"

	"github.com/Godfredasare/go-ecommerce/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"messages": "Authorization token required"})
		ctx.Abort()
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"messages": err.Error()})
		ctx.Abort()
		return
	}

	ctx.Set("userId", userId)
	ctx.Next()
}
