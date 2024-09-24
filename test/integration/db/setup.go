package db_test

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mateusffaria/pismo-challenge/configs"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	if os.Getenv("APP_ENV") == "" {
		err := godotenv.Load("../../../../configs/.local.env")
		if err != nil {
			fmt.Println("failed to load local .env")
		}
	}

	db := configs.NewDatabaseConnection(configs.DBConn{
		Port:     5433,
		Password: os.Getenv("DB_TEST_PASSWORD"),
		User:     os.Getenv("DB_TEST_USER"),
		Dbname:   os.Getenv("DB_TEST_NAME"),
		Host:     os.Getenv("DB_TEST_HOST"),
	})

	dbc, err := db.DB()
	if err != nil {
		log.Fatal("failed to get db connection")
	}

	configs.RunMigrations(dbc, "file://../../../../db/migrations")

	return db
}

func Sunset(db *gorm.DB) {
	db.Exec("DROP TABLE accounts CASCADE")
	db.Exec("DROP TABLE transactions CASCADE")
	db.Exec("DROP TABLE operation_types CASCADE")
	db.Exec("DROP TABLE schema_migrations CASCADE")
}
