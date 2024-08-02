package main

import (
	"github.com/Godfredasare/go-ecommerce/config"
	"github.com/Godfredasare/go-ecommerce/database"
	"github.com/Godfredasare/go-ecommerce/router"
	"github.com/Godfredasare/go-ecommerce/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	database.InitDB()

	defer database.CloseDB()

	utils.InitializeValidatorUniversalTranslator()

	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	server.MaxMultipartMemory = 8 << 20  //8mib
	router.Router(server)

	server.Run(":8080")
}
