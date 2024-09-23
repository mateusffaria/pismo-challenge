package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mateusffaria/pismo-challenge/docs"
	accountsAPI "github.com/mateusffaria/pismo-challenge/internal/accounts/api"
	swagFiles "github.com/swaggo/files"
	ginSwag "github.com/swaggo/gin-swagger"
)

// @title Accounts API
// @version 1.0
// @description A transaction management software API in go using Gin-Framework
// @BasePath /api
func main() {
	r := gin.Default()

	//Setup swagger api docs
	r.GET("/api/docs/*any", ginSwag.WrapHandler(swagFiles.Handler))

	accountsAPI.SetupApi(r)

	r.Run(":" + "8080")
}
