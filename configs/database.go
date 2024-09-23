package configs

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConn struct {
	SSL      bool
	Port     int
	Host     string
	User     string
	Password string
	Dbname   string
}

func NewDatabaseConnection(dbc DBConn) *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", dbc.Host, dbc.Port, dbc.User, dbc.Password, dbc.Dbname)
	if dbc.SSL {
		sqlInfo += " sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("database connection failed")
	}

	return db
}
