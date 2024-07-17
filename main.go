package main

import (
	"github.com/Godfredasare/go-ecommerce/config"
	"github.com/Godfredasare/go-ecommerce/database"
	"github.com/Godfredasare/go-ecommerce/router"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	database.InitDB()

	defer database.CloseDB()

	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	router.Router(server)

	server.Run(":8080")
}

