package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mateusffaria/pismo-challenge/configs"
	_ "github.com/mateusffaria/pismo-challenge/docs"
	accountsAPI "github.com/mateusffaria/pismo-challenge/internal/accounts/api"
	transactionsAPI "github.com/mateusffaria/pismo-challenge/internal/transactions/api"
	swagFiles "github.com/swaggo/files"
	ginSwag "github.com/swaggo/gin-swagger"
)

// @title Accounts API
// @version 1.0
// @description A transaction management software API in go using Gin-Framework
// @BasePath /api
func main() {
	if os.Getenv("APP_ENV") == "" {
		err := godotenv.Load("./configs/.local.env")
		if err != nil {
			fmt.Println("failed to load local .env")
		}
	}

	db := configs.NewDatabaseConnection(configs.DBConn{
		Port:     5432,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
	})

	dbc, err := db.DB()
	if err != nil {
		log.Fatal("failed to get db connection")
	}

	configs.RunMigrations(dbc)

	r := gin.Default()

	//Setup swagger api docs
	if os.Getenv("APP_ENV") != "production" {
		r.GET("/api/docs/*any", ginSwag.WrapHandler(swagFiles.Handler))
	}

	accountsAPI.SetupApi(r, db)
	transactionsAPI.SetupApi(r, db)

	r.Run(":" + "8080")
}
